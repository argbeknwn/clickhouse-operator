// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chi

import (
	"context"
	"fmt"
	"time"

	chop "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	chopmetrics "github.com/altinity/clickhouse-operator/pkg/apis/metrics"
	chopclientset "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned"
	chopclientsetscheme "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned/scheme"
	chopinformers "github.com/altinity/clickhouse-operator/pkg/client/informers/externalversions/clickhouse.altinity.com/v1"
	choplisters "github.com/altinity/clickhouse-operator/pkg/client/listers/clickhouse.altinity.com/v1"
	chopparser "github.com/altinity/clickhouse-operator/pkg/parser"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"

	appsinformers "k8s.io/client-go/informers/apps/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	kube "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcore "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"github.com/golang/glog"
)

// Controller defines CRO controller
type Controller struct {
	kubeClient              kube.Interface
	chopClient              chopclientset.Interface
	chiLister               choplisters.ClickHouseInstallationLister
	chiListerSynced         cache.InformerSynced
	statefulSetLister       appslisters.StatefulSetLister
	statefulSetListerSynced cache.InformerSynced
	configMapLister         corelisters.ConfigMapLister
	configMapListerSynced   cache.InformerSynced
	serviceLister           corelisters.ServiceLister
	serviceListerSynced     cache.InformerSynced
	queue                   workqueue.RateLimitingInterface
	recorder                record.EventRecorder
	metricsExporter         *chopmetrics.Exporter
}

const (
	componentName   = "clickhouse-operator"
	runWorkerPeriod = time.Second
)

const (
	successSynced         = "Synced"
	errResourceExists     = "ErrResourceExists"
	messageResourceSynced = "ClickHouseInstallation synced successfully"
	messageResourceExists = "Resource %q already exists and is not managed by ClickHouseInstallation"
	messageUnableToDecode = "unable to decode object (invalid type)"
	messageUnableToSync   = "unable to sync caches for %s controller"
)

// CreateController creates instance of Controller
func CreateController(
	chopClient chopclientset.Interface,
	kubeClient kube.Interface,
	chiInformer chopinformers.ClickHouseInstallationInformer,
	statefulSetInformer appsinformers.StatefulSetInformer,
	configMapInformer coreinformers.ConfigMapInformer,
	serviceInformer coreinformers.ServiceInformer,
	chopMetricsExporter *chopmetrics.Exporter,
) *Controller {

	// Initializations
	chopclientsetscheme.AddToScheme(scheme.Scheme)
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(
		&typedcore.EventSinkImpl{
			Interface: kubeClient.CoreV1().Events(""),
		},
	)
	recorder := eventBroadcaster.NewRecorder(
		scheme.Scheme,
		core.EventSource{
			Component: componentName,
		},
	)

	// Create Controller instance
	controller := &Controller{
		// kubeClient used to Create() k8s resources as c.kubeClient.AppsV1().StatefulSets(namespace).Create(name)
		kubeClient: kubeClient,
		// chopClient used to Update() CRD k8s resource as c.chopClient.ClickhouseV1().ClickHouseInstallations(chi.Namespace).Update(chiCopy)
		chopClient: chopClient,

		// chiLister used as chiLister.ClickHouseInstallations(namespace).Get(name)
		chiLister: chiInformer.Lister(),
		// chiListerSynced used in waitForCacheSync()
		chiListerSynced: chiInformer.Informer().HasSynced,

		// statefulSetLister used as statefulSetLister.StatefulSets(namespace).Get(name)
		statefulSetLister: statefulSetInformer.Lister(),
		// statefulSetListerSynced used in waitForCacheSync()
		statefulSetListerSynced: statefulSetInformer.Informer().HasSynced,

		// configMapLister used as configMapLister.ConfigMaps(namespace).Get(name)
		configMapLister: configMapInformer.Lister(),
		// configMapListerSynced used in waitForCacheSync()
		configMapListerSynced: configMapInformer.Informer().HasSynced,

		// serviceLister used as serviceLister.Services(namespace).Get(name)
		serviceLister: serviceInformer.Lister(),
		// serviceListerSynced used in waitForCacheSync()
		serviceListerSynced: serviceInformer.Informer().HasSynced,

		// queue used to organize events queue processed by operator
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "chi"),

		// not used explicitly
		recorder: recorder,

		// export metrics to Prometheus
		metricsExporter: chopMetricsExporter,
	}

	chiInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			chi := obj.(*chop.ClickHouseInstallation)
			glog.V(1).Infof("chiInformer.AddFunc - %s/%s added", chi.Namespace, chi.Name)
			controller.enqueueObject(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			glog.V(1).Info("chiInformer.UpdateFunc")

			newChi := newObj.(*chop.ClickHouseInstallation)
			oldChi := oldObj.(*chop.ClickHouseInstallation)

			// Update is called on after each Update() call on k8s resource
			if oldChi.IsNew() && !newChi.IsNew() {
				glog.V(1).Infof("chiInformer.UpdateFunc - no update required - switch from new to known chi.ResourceVersion: %s->%s", oldChi.ResourceVersion, newChi.ResourceVersion)
				return
			}

			// Update is called periodically, don't know why
			if newChi.ResourceVersion == oldChi.ResourceVersion {
				glog.V(1).Info("chiInformer.UpdateFunc - ResourceVersion is the same - periodical housekeeping")
			} else {
				// Looks like real update has happened
				glog.V(1).Infof("chiInformer.UpdateFunc - UPDATE REQUIRED chi.ResourceVersion: %s->%s", oldChi.ResourceVersion, newChi.ResourceVersion)
				glog.V(1).Infof("\n===old===:\n%s\n===new===:\n%s\n=========\n", chopparser.Yaml(oldChi), chopparser.Yaml(newChi))
			}

			controller.enqueueObject(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			chi := obj.(*chop.ClickHouseInstallation)
			glog.V(1).Infof("chiInformer.DeleteFunc - CHI %s/%s deleted", chi.Namespace, chi.Name)
		},
	})

	statefulSetInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			statefulSet := obj.(*apps.StatefulSet)
			glog.V(1).Infof("statefulSetInformer.AddFunc - %s/%s added", statefulSet.Namespace, statefulSet.Name)
			controller.handleObject(obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			glog.V(1).Info("statefulSetInformer.UpdateFunc")
			newStatefulSet := newObj.(*apps.StatefulSet)
			oldStatefulSet := oldObj.(*apps.StatefulSet)
			if newStatefulSet.ResourceVersion == oldStatefulSet.ResourceVersion {
				glog.V(1).Infof("statefulSetInformer.UpdateFunc - no update required, no ResourceVersion change %s", newStatefulSet.ResourceVersion)
				return
			}

			if newStatefulSet.Status.ReadyReplicas == newStatefulSet.Status.Replicas {
				glog.V(1).Infof("statefulSetInformer.UpdateFunc - %s/%s is Ready", newStatefulSet.Namespace, newStatefulSet.Name)
			}

			glog.V(1).Info("statefulSetInformer.UpdateFunc - UPDATE REQUIRED")
			controller.handleObject(newObj)
		},
		DeleteFunc: func(obj interface{}) {
			statefulSet := obj.(*apps.StatefulSet)
			glog.V(1).Infof("statefulSetInformer.DeleteFunc - %s/%s deleted", statefulSet.Namespace, statefulSet.Name)
			controller.handleObject(obj)
		},
	})

	return controller
}

// Run syncs caches, starts workers
func (c *Controller) Run(ctx context.Context, threadiness int) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	glog.V(1).Info("Starting ClickHouseInstallation controller")
	if !waitForCacheSync(
		"ClickHouseInstallation",
		ctx.Done(),
		c.chiListerSynced,
		c.statefulSetListerSynced,
		c.configMapListerSynced,
		c.serviceListerSynced,
	) {
		// Unable to sync
		return
	}

	glog.V(1).Info("ClickHouseInstallation controller: starting workers")
	for i := 0; i < threadiness; i++ {
		glog.V(1).Infof("ClickHouseInstallation controller: starting worker %d with poll period %d", i+1, runWorkerPeriod)
		go wait.Until(c.runWorker, runWorkerPeriod, ctx.Done())
	}
	defer glog.V(1).Info("ClickHouseInstallation controller: shutting down workers")

	glog.V(1).Info("ClickHouseInstallation controller: workers started")
	<-ctx.Done()
}

// runWorker is a convenience wrap over processNextWorkItem()
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem processes objects from the workqueue
func (c *Controller) processNextWorkItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		glog.V(1).Info("processNextWorkItem(): shutdown request")
		return false
	}

	err := func(item interface{}) error {
		defer c.queue.Done(item)

		stringItem, ok := item.(string)
		if !ok {
			// Item is impossible to process, no more retries
			c.queue.Forget(item)
			utilruntime.HandleError(fmt.Errorf("unexpected item in the queue - %#v", item))
			return nil
		}

		// Main reconcile loop function sync an item
		if err := c.syncItem(stringItem); err != nil {
			// Item will be retried later
			return fmt.Errorf("unable to sync an object '%s': %s", stringItem, err.Error())
		}

		// Item is processed, no more retries
		c.queue.Forget(item)

		return nil
	}(item)
	if err != nil {
		utilruntime.HandleError(err)
	}

	// Continue iteration
	return true
}

// syncItem is the main reconcile loop function - reconcile CHI object identified by `key`
func (c *Controller) syncItem(key string) error {
	// Here we assume that `key` identifies CHI object

	glog.V(1).Infof("syncItem(%s) start", key)

	// Extract namespace and name from key - action
	// opposite to <key, err := cache.MetaNamespaceKeyFunc(obj)> in enqueueObject function
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		// Unable to split key into parts
		utilruntime.HandleError(fmt.Errorf("incorrect resource key: %s", key))
		return nil
	}
	glog.V(1).Infof("syncItem(%s/%s) namespace/name extracted", namespace, name)

	// We have `namespace` and `name` which are expected to point to a CHI instance

	// Find CHI specified by namespace/name
	chi, err := c.chiLister.ClickHouseInstallations(namespace).Get(name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("ClickHouseInstallation object '%s' no longer exists in the work queue", key))
			// Not found is not an error on higher level - return "all is ok"
			return nil
		}
		return err
	}

	// And sync this CHI
	return c.syncChi(chi)
}

// syncChi sync CHI to desired state
func (c *Controller) syncChi(chi *chop.ClickHouseInstallation) error {
	// Check CHI object already in sync
	if chi.IsNew() {
		// CHI is a new one - need to create all its objects
		return c.syncNewChi(chi)
	} else {
		// Check consistency of existent resources controlled by the CHI object
		return c.syncKnownChi(chi)
	}
}

// syncNewChi sync new CHI - creates all its resources
func (c *Controller) syncNewChi(chi *chop.ClickHouseInstallation) error {
	// CHI is a new one - need to create all its objects
	// Operator receives CHI struct partially filled by data from .yaml file provided by user
	// We need to create all resources that are needed to run user's .yaml specification
	glog.V(1).Infof("syncNewChi(%s/%s)", chi.Namespace, chi.Name)

	chi, err := c.createCHIResources(chi)
	if err != nil {
		glog.V(2).Infof("ClickHouseInstallation (%q): unable to create controlled resources: %q", chi.Name, err)
		return err
	}

	// Update CHI status in k8s
	if err := c.updateCHIResource(chi); err != nil {
		glog.V(2).Infof("ClickHouseInstallation (%q): unable to update status of CHI resource: %q", chi.Name, err)
		return err
	}

	glog.V(2).Infof("ClickHouseInstallation (%q): controlled resources are synced (created)", chi.Name)
	return nil
}

// syncKnownChi sync CHI which was already created earlier
func (c *Controller) syncKnownChi(chi *chop.ClickHouseInstallation) error {
	glog.V(1).Infof("syncKnownChi(%s/%s)", chi.Namespace, chi.Name)

	c.listStatefulSetResources(chi)
	chi, _ = c.createCHIResources(chi)
	c.updateCHIResource(chi)

	podFQDNs := chopparser.ListPodFQDNs(chi)

	// Check hostnames of the Pods from current CHI object included into chopmetrics.Exporter state
	c.metricsExporter.EnsureControlledValues(chi.Name, podFQDNs)

	return nil
}

// updateCHIResource updates ClickHouseInstallation resource
func (c *Controller) updateCHIResource(chi *chop.ClickHouseInstallation) error {
	_, err := c.chopClient.ClickhouseV1().ClickHouseInstallations(chi.Namespace).Update(chi)

	return err
}

// enqueueObject adds ClickHouseInstallation object to the workqueue
func (c *Controller) enqueueObject(obj interface{}) {
	// Make a string key for an API object
	// The key uses the format <namespace>/<name>
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		utilruntime.HandleError(err)
		return
	}

	c.queue.AddRateLimited(key)
}

// handleObject enqueues CHI which is owner of `obj` into reconcile loop
func (c *Controller) handleObject(obj interface{}) {
	object, ok := obj.(meta.Object)
	if !ok {
		ts, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf(messageUnableToDecode))
			return
		}
		object, ok = ts.Obj.(meta.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf(messageUnableToDecode))
			return
		}
	}

	// object is an instance of meta.Object

	// Checking that we control current StatefulSet Object
	ownerRef := meta.GetControllerOf(object)
	if ownerRef == nil {
		// No owner
		return
	}

	// Ensure owner is of a proper kind
	if ownerRef.Kind != chop.ClickHouseInstallationCRDResourceKind {
		return
	}

	glog.V(2).Infof("Processing object: %s", object.GetName())

	// Get owner - it is expected to be CHI
	chi, err := c.chiLister.ClickHouseInstallations(object.GetNamespace()).Get(ownerRef.Name)

	if err != nil {
		glog.V(2).Infof("ignoring orphaned object '%s' of ClickHouseInstallation '%s'", object.GetSelfLink(), ownerRef.Name)
		return
	}

	// Add CHI object into reconcile loop
	c.enqueueObject(chi)
}

// waitForCacheSync is a logger-wrapper over cache.WaitForCacheSync() and it waits for caches to populate
func waitForCacheSync(name string, stopCh <-chan struct{}, cacheSyncs ...cache.InformerSynced) bool {
	glog.V(1).Infof("Syncing caches for %s controller", name)
	if !cache.WaitForCacheSync(stopCh, cacheSyncs...) {
		utilruntime.HandleError(fmt.Errorf(messageUnableToSync, name))
		return false
	}
	glog.V(1).Infof("Caches are synced for %s controller", name)
	return true
}

// clusterWideSelector returns labels.Selector object
func clusterWideSelector(name string) labels.Selector {
	return labels.SelectorFromSet(labels.Set{
		chopparser.ChopGeneratedLabel: name,
	})
	/*
		glog.V(2).Infof("ClickHouseInstallation (%q) listing controlled resources", chi.Name)
		ssList, err := c.statefulSetLister.StatefulSets(chi.Namespace).List(clusterWideSelector(chi.Name))
		if err != nil {
			return err
		}
		// Listing controlled resources
		for i := range ssList {
			glog.V(2).Infof("ClickHouseInstallation (%q) controlls StatefulSet: %q", chi.Name, ssList[i].Name)
		}
	*/
}
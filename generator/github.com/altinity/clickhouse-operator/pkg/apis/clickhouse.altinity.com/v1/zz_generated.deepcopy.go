//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiCleanup) DeepCopyInto(out *ChiCleanup) {
	*out = *in
	if in.UnknownObjects != nil {
		in, out := &in.UnknownObjects, &out.UnknownObjects
		*out = new(ChiObjectsCleanup)
		**out = **in
	}
	if in.ReconcileFailedObjects != nil {
		in, out := &in.ReconcileFailedObjects, &out.ReconcileFailedObjects
		*out = new(ChiObjectsCleanup)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiCleanup.
func (in *ChiCleanup) DeepCopy() *ChiCleanup {
	if in == nil {
		return nil
	}
	out := new(ChiCleanup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiCluster) DeepCopyInto(out *ChiCluster) {
	*out = *in
	if in.Zookeeper != nil {
		in, out := &in.Zookeeper, &out.Zookeeper
		*out = new(ChiZookeeperConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplateNames)
		**out = **in
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(ChiClusterLayout)
		(*in).DeepCopyInto(*out)
	}
	out.Address = in.Address
	if in.CHI != nil {
		in, out := &in.CHI, &out.CHI
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiCluster.
func (in *ChiCluster) DeepCopy() *ChiCluster {
	if in == nil {
		return nil
	}
	out := new(ChiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterAddress) DeepCopyInto(out *ChiClusterAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterAddress.
func (in *ChiClusterAddress) DeepCopy() *ChiClusterAddress {
	if in == nil {
		return nil
	}
	out := new(ChiClusterAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiClusterLayout) DeepCopyInto(out *ChiClusterLayout) {
	*out = *in
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]ChiShard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ChiReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostsField != nil {
		in, out := &in.HostsField, &out.HostsField
		*out = new(HostsField)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiClusterLayout.
func (in *ChiClusterLayout) DeepCopy() *ChiClusterLayout {
	if in == nil {
		return nil
	}
	out := new(ChiClusterLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDefaults) DeepCopyInto(out *ChiDefaults) {
	*out = *in
	if in.DistributedDDL != nil {
		in, out := &in.DistributedDDL, &out.DistributedDDL
		*out = new(ChiDistributedDDL)
		**out = **in
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplateNames)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDefaults.
func (in *ChiDefaults) DeepCopy() *ChiDefaults {
	if in == nil {
		return nil
	}
	out := new(ChiDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiDistributedDDL) DeepCopyInto(out *ChiDistributedDDL) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiDistributedDDL.
func (in *ChiDistributedDDL) DeepCopy() *ChiDistributedDDL {
	if in == nil {
		return nil
	}
	out := new(ChiDistributedDDL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHost) DeepCopyInto(out *ChiHost) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplateNames)
		**out = **in
	}
	out.Address = in.Address
	out.Config = in.Config
	out.ReconcileAttributes = in.ReconcileAttributes
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = new(appsv1.StatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.CurStatefulSet != nil {
		in, out := &in.CurStatefulSet, &out.CurStatefulSet
		*out = new(appsv1.StatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.DesiredStatefulSet != nil {
		in, out := &in.DesiredStatefulSet, &out.DesiredStatefulSet
		*out = new(appsv1.StatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.CHI != nil {
		in, out := &in.CHI, &out.CHI
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHost.
func (in *ChiHost) DeepCopy() *ChiHost {
	if in == nil {
		return nil
	}
	out := new(ChiHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostAddress) DeepCopyInto(out *ChiHostAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostAddress.
func (in *ChiHostAddress) DeepCopy() *ChiHostAddress {
	if in == nil {
		return nil
	}
	out := new(ChiHostAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostConfig) DeepCopyInto(out *ChiHostConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostConfig.
func (in *ChiHostConfig) DeepCopy() *ChiHostConfig {
	if in == nil {
		return nil
	}
	out := new(ChiHostConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostReconcileAttributes) DeepCopyInto(out *ChiHostReconcileAttributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostReconcileAttributes.
func (in *ChiHostReconcileAttributes) DeepCopy() *ChiHostReconcileAttributes {
	if in == nil {
		return nil
	}
	out := new(ChiHostReconcileAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiHostTemplate) DeepCopyInto(out *ChiHostTemplate) {
	*out = *in
	if in.PortDistribution != nil {
		in, out := &in.PortDistribution, &out.PortDistribution
		*out = make([]ChiPortDistribution, len(*in))
		copy(*out, *in)
	}
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiHostTemplate.
func (in *ChiHostTemplate) DeepCopy() *ChiHostTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiHostTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiObjectsCleanup) DeepCopyInto(out *ChiObjectsCleanup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiObjectsCleanup.
func (in *ChiObjectsCleanup) DeepCopy() *ChiObjectsCleanup {
	if in == nil {
		return nil
	}
	out := new(ChiObjectsCleanup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodDistribution) DeepCopyInto(out *ChiPodDistribution) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodDistribution.
func (in *ChiPodDistribution) DeepCopy() *ChiPodDistribution {
	if in == nil {
		return nil
	}
	out := new(ChiPodDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplate) DeepCopyInto(out *ChiPodTemplate) {
	*out = *in
	in.Zone.DeepCopyInto(&out.Zone)
	if in.PodDistribution != nil {
		in, out := &in.PodDistribution, &out.PodDistribution
		*out = make([]ChiPodDistribution, len(*in))
		copy(*out, *in)
	}
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplate.
func (in *ChiPodTemplate) DeepCopy() *ChiPodTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPodTemplateZone) DeepCopyInto(out *ChiPodTemplateZone) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPodTemplateZone.
func (in *ChiPodTemplateZone) DeepCopy() *ChiPodTemplateZone {
	if in == nil {
		return nil
	}
	out := new(ChiPodTemplateZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiPortDistribution) DeepCopyInto(out *ChiPortDistribution) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiPortDistribution.
func (in *ChiPortDistribution) DeepCopy() *ChiPortDistribution {
	if in == nil {
		return nil
	}
	out := new(ChiPortDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReconciling) DeepCopyInto(out *ChiReconciling) {
	*out = *in
	if in.Cleanup != nil {
		in, out := &in.Cleanup, &out.Cleanup
		*out = new(ChiCleanup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReconciling.
func (in *ChiReconciling) DeepCopy() *ChiReconciling {
	if in == nil {
		return nil
	}
	out := new(ChiReconciling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReplica) DeepCopyInto(out *ChiReplica) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplateNames)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*ChiHost, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ChiHost)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.Address = in.Address
	if in.CHI != nil {
		in, out := &in.CHI, &out.CHI
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReplica.
func (in *ChiReplica) DeepCopy() *ChiReplica {
	if in == nil {
		return nil
	}
	out := new(ChiReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiReplicaAddress) DeepCopyInto(out *ChiReplicaAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiReplicaAddress.
func (in *ChiReplicaAddress) DeepCopy() *ChiReplicaAddress {
	if in == nil {
		return nil
	}
	out := new(ChiReplicaAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiServiceTemplate) DeepCopyInto(out *ChiServiceTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiServiceTemplate.
func (in *ChiServiceTemplate) DeepCopy() *ChiServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShard) DeepCopyInto(out *ChiShard) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplateNames)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*ChiHost, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ChiHost)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.Address = in.Address
	if in.CHI != nil {
		in, out := &in.CHI, &out.CHI
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShard.
func (in *ChiShard) DeepCopy() *ChiShard {
	if in == nil {
		return nil
	}
	out := new(ChiShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiShardAddress) DeepCopyInto(out *ChiShardAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiShardAddress.
func (in *ChiShardAddress) DeepCopy() *ChiShardAddress {
	if in == nil {
		return nil
	}
	out := new(ChiShardAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiSpec) DeepCopyInto(out *ChiSpec) {
	*out = *in
	if in.TaskID != nil {
		in, out := &in.TaskID, &out.TaskID
		*out = new(string)
		**out = **in
	}
	if in.Templating != nil {
		in, out := &in.Templating, &out.Templating
		*out = new(ChiTemplating)
		**out = **in
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(ChiReconciling)
		(*in).DeepCopyInto(*out)
	}
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = new(ChiDefaults)
		(*in).DeepCopyInto(*out)
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(Configuration)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(ChiTemplates)
		(*in).DeepCopyInto(*out)
	}
	if in.UseTemplates != nil {
		in, out := &in.UseTemplates, &out.UseTemplates
		*out = make([]ChiUseTemplate, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiSpec.
func (in *ChiSpec) DeepCopy() *ChiSpec {
	if in == nil {
		return nil
	}
	out := new(ChiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiStatus) DeepCopyInto(out *ChiStatus) {
	*out = *in
	if in.TaskIDsStarted != nil {
		in, out := &in.TaskIDsStarted, &out.TaskIDsStarted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TaskIDsCompleted != nil {
		in, out := &in.TaskIDsCompleted, &out.TaskIDsCompleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FQDNs != nil {
		in, out := &in.FQDNs, &out.FQDNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NormalizedCHI != nil {
		in, out := &in.NormalizedCHI, &out.NormalizedCHI
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiStatus.
func (in *ChiStatus) DeepCopy() *ChiStatus {
	if in == nil {
		return nil
	}
	out := new(ChiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplateNames) DeepCopyInto(out *ChiTemplateNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplateNames.
func (in *ChiTemplateNames) DeepCopy() *ChiTemplateNames {
	if in == nil {
		return nil
	}
	out := new(ChiTemplateNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplates) DeepCopyInto(out *ChiTemplates) {
	*out = *in
	if in.HostTemplates != nil {
		in, out := &in.HostTemplates, &out.HostTemplates
		*out = make([]ChiHostTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodTemplates != nil {
		in, out := &in.PodTemplates, &out.PodTemplates
		*out = make([]ChiPodTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]ChiVolumeClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceTemplates != nil {
		in, out := &in.ServiceTemplates, &out.ServiceTemplates
		*out = make([]ChiServiceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostTemplatesIndex != nil {
		in, out := &in.HostTemplatesIndex, &out.HostTemplatesIndex
		*out = new(HostTemplatesIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.PodTemplatesIndex != nil {
		in, out := &in.PodTemplatesIndex, &out.PodTemplatesIndex
		*out = new(PodTemplatesIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplatesIndex != nil {
		in, out := &in.VolumeClaimTemplatesIndex, &out.VolumeClaimTemplatesIndex
		*out = new(VolumeClaimTemplatesIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceTemplatesIndex != nil {
		in, out := &in.ServiceTemplatesIndex, &out.ServiceTemplatesIndex
		*out = new(ServiceTemplatesIndex)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplates.
func (in *ChiTemplates) DeepCopy() *ChiTemplates {
	if in == nil {
		return nil
	}
	out := new(ChiTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiTemplating) DeepCopyInto(out *ChiTemplating) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiTemplating.
func (in *ChiTemplating) DeepCopy() *ChiTemplating {
	if in == nil {
		return nil
	}
	out := new(ChiTemplating)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiUseTemplate) DeepCopyInto(out *ChiUseTemplate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiUseTemplate.
func (in *ChiUseTemplate) DeepCopy() *ChiUseTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiUseTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiVolumeClaimTemplate) DeepCopyInto(out *ChiVolumeClaimTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiVolumeClaimTemplate.
func (in *ChiVolumeClaimTemplate) DeepCopy() *ChiVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ChiVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperConfig) DeepCopyInto(out *ChiZookeeperConfig) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]ChiZookeeperNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperConfig.
func (in *ChiZookeeperConfig) DeepCopy() *ChiZookeeperConfig {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiZookeeperNode) DeepCopyInto(out *ChiZookeeperNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiZookeeperNode.
func (in *ChiZookeeperNode) DeepCopy() *ChiZookeeperNode {
	if in == nil {
		return nil
	}
	out := new(ChiZookeeperNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallation) DeepCopyInto(out *ClickHouseInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallation.
func (in *ClickHouseInstallation) DeepCopy() *ClickHouseInstallation {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationList) DeepCopyInto(out *ClickHouseInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationList.
func (in *ClickHouseInstallationList) DeepCopy() *ClickHouseInstallationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationTemplate) DeepCopyInto(out *ClickHouseInstallationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationTemplate.
func (in *ClickHouseInstallationTemplate) DeepCopy() *ClickHouseInstallationTemplate {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseInstallationTemplateList) DeepCopyInto(out *ClickHouseInstallationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseInstallationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseInstallationTemplateList.
func (in *ClickHouseInstallationTemplateList) DeepCopy() *ClickHouseInstallationTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseInstallationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseInstallationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseOperatorConfiguration) DeepCopyInto(out *ClickHouseOperatorConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseOperatorConfiguration.
func (in *ClickHouseOperatorConfiguration) DeepCopy() *ClickHouseOperatorConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClickHouseOperatorConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseOperatorConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseOperatorConfigurationList) DeepCopyInto(out *ClickHouseOperatorConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseOperatorConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseOperatorConfigurationList.
func (in *ClickHouseOperatorConfigurationList) DeepCopy() *ClickHouseOperatorConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseOperatorConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseOperatorConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	if in.Zookeeper != nil {
		in, out := &in.Zookeeper, &out.Zookeeper
		*out = new(ChiZookeeperConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]*ChiCluster, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ChiCluster)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostTemplatesIndex) DeepCopyInto(out *HostTemplatesIndex) {
	*out = *in
	if in.v != nil {
		in, out := &in.v, &out.v
		*out = make(map[string]*ChiHostTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiHostTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiHostTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostTemplatesIndex.
func (in *HostTemplatesIndex) DeepCopy() *HostTemplatesIndex {
	if in == nil {
		return nil
	}
	out := new(HostTemplatesIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostsField) DeepCopyInto(out *HostsField) {
	*out = *in
	if in.Field != nil {
		in, out := &in.Field, &out.Field
		*out = make([][]*ChiHost, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]*ChiHost, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(ChiHost)
						(*in).DeepCopyInto(*out)
					}
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostsField.
func (in *HostsField) DeepCopy() *HostsField {
	if in == nil {
		return nil
	}
	out := new(HostsField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfig) DeepCopyInto(out *OperatorConfig) {
	*out = *in
	if in.WatchNamespaces != nil {
		in, out := &in.WatchNamespaces, &out.WatchNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CHCommonConfigs != nil {
		in, out := &in.CHCommonConfigs, &out.CHCommonConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CHHostConfigs != nil {
		in, out := &in.CHHostConfigs, &out.CHHostConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CHUsersConfigs != nil {
		in, out := &in.CHUsersConfigs, &out.CHUsersConfigs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CHITemplateFiles != nil {
		in, out := &in.CHITemplateFiles, &out.CHITemplateFiles
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CHITemplates != nil {
		in, out := &in.CHITemplates, &out.CHITemplates
		*out = make([]*ClickHouseInstallation, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClickHouseInstallation)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CHITemplate != nil {
		in, out := &in.CHITemplate, &out.CHITemplate
		*out = new(ClickHouseInstallation)
		(*in).DeepCopyInto(*out)
	}
	if in.CHConfigUserDefaultNetworksIP != nil {
		in, out := &in.CHConfigUserDefaultNetworksIP, &out.CHConfigUserDefaultNetworksIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeIntoPropagationAnnotations != nil {
		in, out := &in.IncludeIntoPropagationAnnotations, &out.IncludeIntoPropagationAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeFromPropagationAnnotations != nil {
		in, out := &in.ExcludeFromPropagationAnnotations, &out.ExcludeFromPropagationAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeIntoPropagationLabels != nil {
		in, out := &in.IncludeIntoPropagationLabels, &out.IncludeIntoPropagationLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeFromPropagationLabels != nil {
		in, out := &in.ExcludeFromPropagationLabels, &out.ExcludeFromPropagationLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfig.
func (in *OperatorConfig) DeepCopy() *OperatorConfig {
	if in == nil {
		return nil
	}
	out := new(OperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplatesIndex) DeepCopyInto(out *PodTemplatesIndex) {
	*out = *in
	if in.v != nil {
		in, out := &in.v, &out.v
		*out = make(map[string]*ChiPodTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiPodTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiPodTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTemplatesIndex.
func (in *PodTemplatesIndex) DeepCopy() *PodTemplatesIndex {
	if in == nil {
		return nil
	}
	out := new(PodTemplatesIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTemplatesIndex) DeepCopyInto(out *ServiceTemplatesIndex) {
	*out = *in
	if in.v != nil {
		in, out := &in.v, &out.v
		*out = make(map[string]*ChiServiceTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiServiceTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiServiceTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTemplatesIndex.
func (in *ServiceTemplatesIndex) DeepCopy() *ServiceTemplatesIndex {
	if in == nil {
		return nil
	}
	out := new(ServiceTemplatesIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Setting) DeepCopyInto(out *Setting) {
	*out = *in
	if in.vector != nil {
		in, out := &in.vector, &out.vector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Setting.
func (in *Setting) DeepCopy() *Setting {
	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	if in.m != nil {
		in, out := &in.m, &out.m
		*out = make(map[string]*Setting, len(*in))
		for key, val := range *in {
			var outVal *Setting
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Setting)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClaimTemplatesIndex) DeepCopyInto(out *VolumeClaimTemplatesIndex) {
	*out = *in
	if in.v != nil {
		in, out := &in.v, &out.v
		*out = make(map[string]*ChiVolumeClaimTemplate, len(*in))
		for key, val := range *in {
			var outVal *ChiVolumeClaimTemplate
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ChiVolumeClaimTemplate)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClaimTemplatesIndex.
func (in *VolumeClaimTemplatesIndex) DeepCopy() *VolumeClaimTemplatesIndex {
	if in == nil {
		return nil
	}
	out := new(VolumeClaimTemplatesIndex)
	in.DeepCopyInto(out)
	return out
}

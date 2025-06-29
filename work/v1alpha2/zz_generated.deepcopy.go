//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Karmada Authors.

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

package v1alpha2

import (
	v1alpha1 "github.com/karmada-io/api/policy/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatedStatusItem) DeepCopyInto(out *AggregatedStatusItem) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatedStatusItem.
func (in *AggregatedStatusItem) DeepCopy() *AggregatedStatusItem {
	if in == nil {
		return nil
	}
	out := new(AggregatedStatusItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingSnapshot) DeepCopyInto(out *BindingSnapshot) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]TargetCluster, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingSnapshot.
func (in *BindingSnapshot) DeepCopy() *BindingSnapshot {
	if in == nil {
		return nil
	}
	out := new(BindingSnapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceBinding) DeepCopyInto(out *ClusterResourceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceBinding.
func (in *ClusterResourceBinding) DeepCopy() *ClusterResourceBinding {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceBindingList) DeepCopyInto(out *ClusterResourceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceBindingList.
func (in *ClusterResourceBindingList) DeepCopy() *ClusterResourceBindingList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracefulEvictionTask) DeepCopyInto(out *GracefulEvictionTask) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.GracePeriodSeconds != nil {
		in, out := &in.GracePeriodSeconds, &out.GracePeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SuppressDeletion != nil {
		in, out := &in.SuppressDeletion, &out.SuppressDeletion
		*out = new(bool)
		**out = **in
	}
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracefulEvictionTask.
func (in *GracefulEvictionTask) DeepCopy() *GracefulEvictionTask {
	if in == nil {
		return nil
	}
	out := new(GracefulEvictionTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeClaim) DeepCopyInto(out *NodeClaim) {
	*out = *in
	if in.HardNodeAffinity != nil {
		in, out := &in.HardNodeAffinity, &out.HardNodeAffinity
		*out = new(v1.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeClaim.
func (in *NodeClaim) DeepCopy() *NodeClaim {
	if in == nil {
		return nil
	}
	out := new(NodeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaRequirements) DeepCopyInto(out *ReplicaRequirements) {
	*out = *in
	if in.NodeClaim != nil {
		in, out := &in.NodeClaim, &out.NodeClaim
		*out = new(NodeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceRequest != nil {
		in, out := &in.ResourceRequest, &out.ResourceRequest
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaRequirements.
func (in *ReplicaRequirements) DeepCopy() *ReplicaRequirements {
	if in == nil {
		return nil
	}
	out := new(ReplicaRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBinding) DeepCopyInto(out *ResourceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBinding.
func (in *ResourceBinding) DeepCopy() *ResourceBinding {
	if in == nil {
		return nil
	}
	out := new(ResourceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingList) DeepCopyInto(out *ResourceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingList.
func (in *ResourceBindingList) DeepCopy() *ResourceBindingList {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingSpec) DeepCopyInto(out *ResourceBindingSpec) {
	*out = *in
	out.Resource = in.Resource
	if in.ReplicaRequirements != nil {
		in, out := &in.ReplicaRequirements, &out.ReplicaRequirements
		*out = new(ReplicaRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]TargetCluster, len(*in))
		copy(*out, *in)
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(v1alpha1.Placement)
		(*in).DeepCopyInto(*out)
	}
	if in.GracefulEvictionTasks != nil {
		in, out := &in.GracefulEvictionTasks, &out.GracefulEvictionTasks
		*out = make([]GracefulEvictionTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequiredBy != nil {
		in, out := &in.RequiredBy, &out.RequiredBy
		*out = make([]BindingSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = new(v1alpha1.FailoverBehavior)
		(*in).DeepCopyInto(*out)
	}
	if in.RescheduleTriggeredAt != nil {
		in, out := &in.RescheduleTriggeredAt, &out.RescheduleTriggeredAt
		*out = (*in).DeepCopy()
	}
	if in.Suspension != nil {
		in, out := &in.Suspension, &out.Suspension
		*out = new(v1alpha1.Suspension)
		(*in).DeepCopyInto(*out)
	}
	if in.PreserveResourcesOnDeletion != nil {
		in, out := &in.PreserveResourcesOnDeletion, &out.PreserveResourcesOnDeletion
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingSpec.
func (in *ResourceBindingSpec) DeepCopy() *ResourceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingStatus) DeepCopyInto(out *ResourceBindingStatus) {
	*out = *in
	if in.LastScheduledTime != nil {
		in, out := &in.LastScheduledTime, &out.LastScheduledTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AggregatedStatus != nil {
		in, out := &in.AggregatedStatus, &out.AggregatedStatus
		*out = make([]AggregatedStatusItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingStatus.
func (in *ResourceBindingStatus) DeepCopy() *ResourceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetCluster) DeepCopyInto(out *TargetCluster) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetCluster.
func (in *TargetCluster) DeepCopy() *TargetCluster {
	if in == nil {
		return nil
	}
	out := new(TargetCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskOptions) DeepCopyInto(out *TaskOptions) {
	*out = *in
	if in.gracePeriodSeconds != nil {
		in, out := &in.gracePeriodSeconds, &out.gracePeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.suppressDeletion != nil {
		in, out := &in.suppressDeletion, &out.suppressDeletion
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskOptions.
func (in *TaskOptions) DeepCopy() *TaskOptions {
	if in == nil {
		return nil
	}
	out := new(TaskOptions)
	in.DeepCopyInto(out)
	return out
}

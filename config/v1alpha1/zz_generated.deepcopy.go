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

package v1alpha1

import (
	v1alpha2 "github.com/karmada-io/api/work/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationRules) DeepCopyInto(out *CustomizationRules) {
	*out = *in
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(LocalValueRetention)
		**out = **in
	}
	if in.ReplicaResource != nil {
		in, out := &in.ReplicaResource, &out.ReplicaResource
		*out = new(ReplicaResourceRequirement)
		**out = **in
	}
	if in.ReplicaRevision != nil {
		in, out := &in.ReplicaRevision, &out.ReplicaRevision
		*out = new(ReplicaRevision)
		**out = **in
	}
	if in.StatusReflection != nil {
		in, out := &in.StatusReflection, &out.StatusReflection
		*out = new(StatusReflection)
		**out = **in
	}
	if in.StatusAggregation != nil {
		in, out := &in.StatusAggregation, &out.StatusAggregation
		*out = new(StatusAggregation)
		**out = **in
	}
	if in.HealthInterpretation != nil {
		in, out := &in.HealthInterpretation, &out.HealthInterpretation
		*out = new(HealthInterpretation)
		**out = **in
	}
	if in.DependencyInterpretation != nil {
		in, out := &in.DependencyInterpretation, &out.DependencyInterpretation
		*out = new(DependencyInterpretation)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationRules.
func (in *CustomizationRules) DeepCopy() *CustomizationRules {
	if in == nil {
		return nil
	}
	out := new(CustomizationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizationTarget) DeepCopyInto(out *CustomizationTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizationTarget.
func (in *CustomizationTarget) DeepCopy() *CustomizationTarget {
	if in == nil {
		return nil
	}
	out := new(CustomizationTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencyInterpretation) DeepCopyInto(out *DependencyInterpretation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencyInterpretation.
func (in *DependencyInterpretation) DeepCopy() *DependencyInterpretation {
	if in == nil {
		return nil
	}
	out := new(DependencyInterpretation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependentObjectReference) DeepCopyInto(out *DependentObjectReference) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependentObjectReference.
func (in *DependentObjectReference) DeepCopy() *DependentObjectReference {
	if in == nil {
		return nil
	}
	out := new(DependentObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthInterpretation) DeepCopyInto(out *HealthInterpretation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthInterpretation.
func (in *HealthInterpretation) DeepCopy() *HealthInterpretation {
	if in == nil {
		return nil
	}
	out := new(HealthInterpretation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalValueRetention) DeepCopyInto(out *LocalValueRetention) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalValueRetention.
func (in *LocalValueRetention) DeepCopy() *LocalValueRetention {
	if in == nil {
		return nil
	}
	out := new(LocalValueRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaResourceRequirement) DeepCopyInto(out *ReplicaResourceRequirement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaResourceRequirement.
func (in *ReplicaResourceRequirement) DeepCopy() *ReplicaResourceRequirement {
	if in == nil {
		return nil
	}
	out := new(ReplicaResourceRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaRevision) DeepCopyInto(out *ReplicaRevision) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaRevision.
func (in *ReplicaRevision) DeepCopy() *ReplicaRevision {
	if in == nil {
		return nil
	}
	out := new(ReplicaRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestStatus) DeepCopyInto(out *RequestStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestStatus.
func (in *RequestStatus) DeepCopy() *RequestStatus {
	if in == nil {
		return nil
	}
	out := new(RequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterContext) DeepCopyInto(out *ResourceInterpreterContext) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(ResourceInterpreterRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(ResourceInterpreterResponse)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterContext.
func (in *ResourceInterpreterContext) DeepCopy() *ResourceInterpreterContext {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInterpreterContext) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterCustomization) DeepCopyInto(out *ResourceInterpreterCustomization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterCustomization.
func (in *ResourceInterpreterCustomization) DeepCopy() *ResourceInterpreterCustomization {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterCustomization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInterpreterCustomization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterCustomizationList) DeepCopyInto(out *ResourceInterpreterCustomizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceInterpreterCustomization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterCustomizationList.
func (in *ResourceInterpreterCustomizationList) DeepCopy() *ResourceInterpreterCustomizationList {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterCustomizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInterpreterCustomizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterCustomizationSpec) DeepCopyInto(out *ResourceInterpreterCustomizationSpec) {
	*out = *in
	out.Target = in.Target
	in.Customizations.DeepCopyInto(&out.Customizations)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterCustomizationSpec.
func (in *ResourceInterpreterCustomizationSpec) DeepCopy() *ResourceInterpreterCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterRequest) DeepCopyInto(out *ResourceInterpreterRequest) {
	*out = *in
	out.Kind = in.Kind
	in.Object.DeepCopyInto(&out.Object)
	if in.ObservedObject != nil {
		in, out := &in.ObservedObject, &out.ObservedObject
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int32)
		**out = **in
	}
	if in.AggregatedStatus != nil {
		in, out := &in.AggregatedStatus, &out.AggregatedStatus
		*out = make([]v1alpha2.AggregatedStatusItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterRequest.
func (in *ResourceInterpreterRequest) DeepCopy() *ResourceInterpreterRequest {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterResponse) DeepCopyInto(out *ResourceInterpreterResponse) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RequestStatus)
		**out = **in
	}
	if in.Patch != nil {
		in, out := &in.Patch, &out.Patch
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.PatchType != nil {
		in, out := &in.PatchType, &out.PatchType
		*out = new(PatchType)
		**out = **in
	}
	if in.ReplicaRequirements != nil {
		in, out := &in.ReplicaRequirements, &out.ReplicaRequirements
		*out = new(v1alpha2.ReplicaRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]DependentObjectReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RawStatus != nil {
		in, out := &in.RawStatus, &out.RawStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterResponse.
func (in *ResourceInterpreterResponse) DeepCopy() *ResourceInterpreterResponse {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterWebhook) DeepCopyInto(out *ResourceInterpreterWebhook) {
	*out = *in
	in.ClientConfig.DeepCopyInto(&out.ClientConfig)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleWithOperations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.InterpreterContextVersions != nil {
		in, out := &in.InterpreterContextVersions, &out.InterpreterContextVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterWebhook.
func (in *ResourceInterpreterWebhook) DeepCopy() *ResourceInterpreterWebhook {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterWebhookConfiguration) DeepCopyInto(out *ResourceInterpreterWebhookConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]ResourceInterpreterWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterWebhookConfiguration.
func (in *ResourceInterpreterWebhookConfiguration) DeepCopy() *ResourceInterpreterWebhookConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterWebhookConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInterpreterWebhookConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInterpreterWebhookConfigurationList) DeepCopyInto(out *ResourceInterpreterWebhookConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceInterpreterWebhookConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInterpreterWebhookConfigurationList.
func (in *ResourceInterpreterWebhookConfigurationList) DeepCopy() *ResourceInterpreterWebhookConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ResourceInterpreterWebhookConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceInterpreterWebhookConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIVersions != nil {
		in, out := &in.APIVersions, &out.APIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Kinds != nil {
		in, out := &in.Kinds, &out.Kinds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWithOperations) DeepCopyInto(out *RuleWithOperations) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]InterpreterOperation, len(*in))
		copy(*out, *in)
	}
	in.Rule.DeepCopyInto(&out.Rule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWithOperations.
func (in *RuleWithOperations) DeepCopy() *RuleWithOperations {
	if in == nil {
		return nil
	}
	out := new(RuleWithOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusAggregation) DeepCopyInto(out *StatusAggregation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusAggregation.
func (in *StatusAggregation) DeepCopy() *StatusAggregation {
	if in == nil {
		return nil
	}
	out := new(StatusAggregation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusReflection) DeepCopyInto(out *StatusReflection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusReflection.
func (in *StatusReflection) DeepCopy() *StatusReflection {
	if in == nil {
		return nil
	}
	out := new(StatusReflection)
	in.DeepCopyInto(out)
	return out
}

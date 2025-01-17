//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAttachmentDefinition) DeepCopyInto(out *NetworkAttachmentDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAttachmentDefinition.
func (in *NetworkAttachmentDefinition) DeepCopy() *NetworkAttachmentDefinition {
	if in == nil {
		return nil
	}
	out := new(NetworkAttachmentDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAttachmentDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAttachmentDefinitionList) DeepCopyInto(out *NetworkAttachmentDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkAttachmentDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAttachmentDefinitionList.
func (in *NetworkAttachmentDefinitionList) DeepCopy() *NetworkAttachmentDefinitionList {
	if in == nil {
		return nil
	}
	out := new(NetworkAttachmentDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAttachmentDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAttachmentDefinitionSpec) DeepCopyInto(out *NetworkAttachmentDefinitionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAttachmentDefinitionSpec.
func (in *NetworkAttachmentDefinitionSpec) DeepCopy() *NetworkAttachmentDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkAttachmentDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

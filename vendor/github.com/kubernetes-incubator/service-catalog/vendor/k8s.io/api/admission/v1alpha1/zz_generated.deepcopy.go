// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AdmissionReview).DeepCopyInto(out.(*AdmissionReview))
			return nil
		}, InType: reflect.TypeOf(&AdmissionReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AdmissionReviewSpec).DeepCopyInto(out.(*AdmissionReviewSpec))
			return nil
		}, InType: reflect.TypeOf(&AdmissionReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AdmissionReviewStatus).DeepCopyInto(out.(*AdmissionReviewStatus))
			return nil
		}, InType: reflect.TypeOf(&AdmissionReviewStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionReview) DeepCopyInto(out *AdmissionReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionReview.
func (in *AdmissionReview) DeepCopy() *AdmissionReview {
	if in == nil {
		return nil
	}
	out := new(AdmissionReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdmissionReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionReviewSpec) DeepCopyInto(out *AdmissionReviewSpec) {
	*out = *in
	out.Kind = in.Kind
	in.Object.DeepCopyInto(&out.Object)
	in.OldObject.DeepCopyInto(&out.OldObject)
	out.Resource = in.Resource
	in.UserInfo.DeepCopyInto(&out.UserInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionReviewSpec.
func (in *AdmissionReviewSpec) DeepCopy() *AdmissionReviewSpec {
	if in == nil {
		return nil
	}
	out := new(AdmissionReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionReviewStatus) DeepCopyInto(out *AdmissionReviewStatus) {
	*out = *in
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Status)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionReviewStatus.
func (in *AdmissionReviewStatus) DeepCopy() *AdmissionReviewStatus {
	if in == nil {
		return nil
	}
	out := new(AdmissionReviewStatus)
	in.DeepCopyInto(out)
	return out
}

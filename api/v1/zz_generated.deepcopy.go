//go:build !ignore_autogenerated

/*
Copyright 2025.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifacts) DeepCopyInto(out *Artifacts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifacts.
func (in *Artifacts) DeepCopy() *Artifacts {
	if in == nil {
		return nil
	}
	out := new(Artifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLink) DeepCopyInto(out *ExternalLink) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLink.
func (in *ExternalLink) DeepCopy() *ExternalLink {
	if in == nil {
		return nil
	}
	out := new(ExternalLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gandelm) DeepCopyInto(out *Gandelm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gandelm.
func (in *Gandelm) DeepCopy() *Gandelm {
	if in == nil {
		return nil
	}
	out := new(Gandelm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gandelm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmArtifact) DeepCopyInto(out *GandelmArtifact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmArtifact.
func (in *GandelmArtifact) DeepCopy() *GandelmArtifact {
	if in == nil {
		return nil
	}
	out := new(GandelmArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmArtifact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmArtifactList) DeepCopyInto(out *GandelmArtifactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GandelmArtifact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmArtifactList.
func (in *GandelmArtifactList) DeepCopy() *GandelmArtifactList {
	if in == nil {
		return nil
	}
	out := new(GandelmArtifactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmArtifactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmArtifactSpec) DeepCopyInto(out *GandelmArtifactSpec) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmArtifactSpec.
func (in *GandelmArtifactSpec) DeepCopy() *GandelmArtifactSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmArtifactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmArtifactStatus) DeepCopyInto(out *GandelmArtifactStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmArtifactStatus.
func (in *GandelmArtifactStatus) DeepCopy() *GandelmArtifactStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmArtifactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmCatalog) DeepCopyInto(out *GandelmCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmCatalog.
func (in *GandelmCatalog) DeepCopy() *GandelmCatalog {
	if in == nil {
		return nil
	}
	out := new(GandelmCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmCatalogList) DeepCopyInto(out *GandelmCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GandelmCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmCatalogList.
func (in *GandelmCatalogList) DeepCopy() *GandelmCatalogList {
	if in == nil {
		return nil
	}
	out := new(GandelmCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmCatalogSpec) DeepCopyInto(out *GandelmCatalogSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ArtifactTags != nil {
		in, out := &in.ArtifactTags, &out.ArtifactTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmCatalogSpec.
func (in *GandelmCatalogSpec) DeepCopy() *GandelmCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmCatalogStatus) DeepCopyInto(out *GandelmCatalogStatus) {
	*out = *in
	in.Timestamp.DeepCopyInto(&out.Timestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmCatalogStatus.
func (in *GandelmCatalogStatus) DeepCopy() *GandelmCatalogStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmCatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmLabel) DeepCopyInto(out *GandelmLabel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmLabel.
func (in *GandelmLabel) DeepCopy() *GandelmLabel {
	if in == nil {
		return nil
	}
	out := new(GandelmLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmLabel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmLabelList) DeepCopyInto(out *GandelmLabelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GandelmLabel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmLabelList.
func (in *GandelmLabelList) DeepCopy() *GandelmLabelList {
	if in == nil {
		return nil
	}
	out := new(GandelmLabelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmLabelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmLabelSpec) DeepCopyInto(out *GandelmLabelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmLabelSpec.
func (in *GandelmLabelSpec) DeepCopy() *GandelmLabelSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmLabelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmLabelStatus) DeepCopyInto(out *GandelmLabelStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmLabelStatus.
func (in *GandelmLabelStatus) DeepCopy() *GandelmLabelStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmLabelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmList) DeepCopyInto(out *GandelmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gandelm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmList.
func (in *GandelmList) DeepCopy() *GandelmList {
	if in == nil {
		return nil
	}
	out := new(GandelmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmSpec) DeepCopyInto(out *GandelmSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmSpec.
func (in *GandelmSpec) DeepCopy() *GandelmSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmStatus) DeepCopyInto(out *GandelmStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmStatus.
func (in *GandelmStatus) DeepCopy() *GandelmStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmVersion) DeepCopyInto(out *GandelmVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmVersion.
func (in *GandelmVersion) DeepCopy() *GandelmVersion {
	if in == nil {
		return nil
	}
	out := new(GandelmVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmVersionList) DeepCopyInto(out *GandelmVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GandelmVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmVersionList.
func (in *GandelmVersionList) DeepCopy() *GandelmVersionList {
	if in == nil {
		return nil
	}
	out := new(GandelmVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmVersionSpec) DeepCopyInto(out *GandelmVersionSpec) {
	*out = *in
	if in.BaseArtifacts != nil {
		in, out := &in.BaseArtifacts, &out.BaseArtifacts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.UpdatedAt.DeepCopyInto(&out.UpdatedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmVersionSpec.
func (in *GandelmVersionSpec) DeepCopy() *GandelmVersionSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmVersionStatus) DeepCopyInto(out *GandelmVersionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmVersionStatus.
func (in *GandelmVersionStatus) DeepCopy() *GandelmVersionStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmWorkload) DeepCopyInto(out *GandelmWorkload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmWorkload.
func (in *GandelmWorkload) DeepCopy() *GandelmWorkload {
	if in == nil {
		return nil
	}
	out := new(GandelmWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmWorkload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmWorkloadList) DeepCopyInto(out *GandelmWorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GandelmWorkload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmWorkloadList.
func (in *GandelmWorkloadList) DeepCopy() *GandelmWorkloadList {
	if in == nil {
		return nil
	}
	out := new(GandelmWorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GandelmWorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmWorkloadSpec) DeepCopyInto(out *GandelmWorkloadSpec) {
	*out = *in
	if in.ExternalLinks != nil {
		in, out := &in.ExternalLinks, &out.ExternalLinks
		*out = make([]ExternalLink, len(*in))
		copy(*out, *in)
	}
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]Artifacts, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmWorkloadSpec.
func (in *GandelmWorkloadSpec) DeepCopy() *GandelmWorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(GandelmWorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GandelmWorkloadStatus) DeepCopyInto(out *GandelmWorkloadStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GandelmWorkloadStatus.
func (in *GandelmWorkloadStatus) DeepCopy() *GandelmWorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(GandelmWorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

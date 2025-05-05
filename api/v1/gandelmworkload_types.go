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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GandelmWorkloadSpec defines the desired state of GandelmWorkload.
type GandelmWorkloadSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Endpoint   string `json:"endpoint"`
	Entrypoint string `json:"entrypoint"`

	ExternalLinks []ExternalLink `json:"external_links,omitempty"`
	Artifacts     []Artifacts    `json:"artifacts,omitempty"`
}

type ExternalLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Artifacts struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// GandelmWorkloadStatus defines the observed state of GandelmWorkload.
type GandelmWorkloadStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GandelmWorkload is the Schema for the gandelmworkloads API.
type GandelmWorkload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GandelmWorkloadSpec   `json:"spec,omitempty"`
	Status GandelmWorkloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GandelmWorkloadList contains a list of GandelmWorkload.
type GandelmWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GandelmWorkload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GandelmWorkload{}, &GandelmWorkloadList{})
}

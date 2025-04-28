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

// GandelmVersionSpec defines the desired state of GandelmVersion.
type GandelmVersionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	BaseArtifacts map[string]string `json:"base_artifacts"`
	BasePriority  uint32            `json:"base_priority"`
	CreatedAt     metav1.Time       `json:"created_at"`
	UpdatedAt     metav1.Time       `json:"updated_at"`
}

// GandelmVersionStatus defines the observed state of GandelmVersion.
type GandelmVersionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GandelmVersion is the Schema for the gandelmversions API.
type GandelmVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GandelmVersionSpec   `json:"spec,omitempty"`
	Status GandelmVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GandelmVersionList contains a list of GandelmVersion.
type GandelmVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GandelmVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GandelmVersion{}, &GandelmVersionList{})
}

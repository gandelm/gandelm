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

// GandelmCatalogSpec defines the desired state of GandelmCatalog.
type GandelmCatalogSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ID       string `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Priority uint32 `json:"priority"`

	Description  string            `json:"description,omitempty"`
	Labels       []string          `json:"labels,omitempty"`
	ArtifactTags map[string]string `json:"artifactTags,omitempty"`
}

// GandelmCatalogStatus defines the observed state of GandelmCatalog.
type GandelmCatalogStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Phase              string      `json:"phase,omitempty"`
	Message            string      `json:"message,omitempty"`
	ObservedGeneration int64       `json:"observedGeneration,omitempty"`
	Timestamp          metav1.Time `json:"timestamp,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GandelmCatalog is the Schema for the gandelmcatalogs API.
type GandelmCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GandelmCatalogSpec   `json:"spec,omitempty"`
	Status GandelmCatalogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GandelmCatalogList contains a list of GandelmCatalog.
type GandelmCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GandelmCatalog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GandelmCatalog{}, &GandelmCatalogList{})
}

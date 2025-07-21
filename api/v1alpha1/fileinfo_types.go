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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FileInfoSpec defines the desired state of FileInfo.
type FileInfoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of FileInfo. Edit fileinfo_types.go to remove/update

	FilePath string `json:"filePath"`
}

type FileDetails struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// FileInfoStatus defines the observed state of FileInfo.
type FileInfoStatus struct {
	NodeName string        `json:"nodeName"`
	Files    []FileDetails `json:"files"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// FileInfo is the Schema for the fileinfoes API.
type FileInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FileInfoSpec   `json:"spec,omitempty"`
	Status FileInfoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileInfoList contains a list of FileInfo.
type FileInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileInfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FileInfo{}, &FileInfoList{})
}

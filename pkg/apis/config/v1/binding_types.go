/*
 * Copyright (C) 2019  Rohith Jayawardene <info@appvia.io>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package v1

import (
	corev1 "github.com/appvia/hub-apis/pkg/apis/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BindingType defines a binding resource type
type BindingType string

const (
	// ResourceType indicates a resource binding type
	ResourceType BindingType = "resource"
	// ClassType indicates a class binding type
	ClassType BindingType = "class"
)

// BindingSpec defines the desired state of Binding
// +k8s:openapi-gen=true
type BindingSpec struct {
	// Class is the reference to the provider of this class
	// +kubebuilder:validation:Optional
	Class corev1.Ownership `json:"class"`
	// Ref is a reference to the configuration object
	// +kubebuilder:validation:Required
	Ref corev1.Ownership `json:"ref"`
}

// BindingStatus defines the observed state of Binding
// +k8s:openapi-gen=true
type BindingStatus struct {
	// Conditions is a set of condition which has caused an error
	// +listType
	Conditions []corev1.Condition `json:"conditions"`
	// Status is overall status of the workspace
	Status corev1.Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Binding is the Schema for the class API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bindings,scope=Namespaced
type Binding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BindingSpec   `json:"spec,omitempty"`
	Status BindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BindingList contains a list of Binding
type BindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Binding `json:"items"`
}

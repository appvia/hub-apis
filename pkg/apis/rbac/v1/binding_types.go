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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// APIGroupReference is the annotation on the binding reference
	APIGroupReference = "rbac.hub.appia.io/group"
	// KindReference is a reference to the kind
	KindReference = "rbac.hub.appia.io/kind"
)

// BindingSpec defines the desired state of Binding
// +k8s:openapi-gen=true
type BindingSpec struct {
	// RoleRef is the reference role we are associated to
	// +kubebuilder:validation:Required
	RoleRef RoleRef `json:"roleRef"`
	// Subjects is a collection of subjects to are binding to the role
	// +listType
	Subjects []Subject `json:"subjects"`
}

// RoleRef provides a reference to the role type
type RoleRef struct {
	// Kind is the type of role we are referencing
	// +kubebuilder:validation:MinLength=1
	Kind string `json:"kind"`
	// Name is the name of the role
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// APIGroup is the api group the role is situated
	// +kubebuilder:validation:MinLength=1
	APIGroup string `json:"apiGroup"`
}

// Subject is user or team we are referencing
type Subject struct {
	// APIGroup is the apigroup if the resoruce we are binding to
	// +kubebuilder:validation:MinLength=1
	APIGroup string `json:"apiGroup"`
	// Kind is the type of resource we are binding to User, Team or Workspace
	Kind string `json:"kind"`
	// Name is the name of the resource in the apigroup
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
}

// BindingStatus defines the observed state of Binding
// +k8s:openapi-gen=true
type BindingStatus struct {
	// Conditions is collection of potentials error causes
	// +listType
	Condiitions []metav1.Status `json:"condiitions"`
	// Status provides an overview of the user status
	Status metav1.StatusReason `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Binding is the Schema for the bindings API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bindings
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

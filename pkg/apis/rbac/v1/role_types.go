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

// RoleSpec defines the desired state of Role
// +k8s:openapi-gen=true
type RoleSpec struct {
	// Rules is a collection of rules defined by this role
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	// +listType
	Rules []Rule `json:"rules"`
}

// Rule defines access to one of more resources and apigroups
type Rule struct {
	// Group is the apigroup we are referencing
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	// +listType
	APIGroups []string `json:"apiGroups"`
	// Resources is a reference to a object kind
	// +kubebuilder:validation:MinItems=1
	// +listType
	Resources []string `json:"resources"`
	// ResourceNames provides the association to map an integration
	// and a provider to team / user
	// +listType
	// +optional
	ResourceNames []string `json:"resourceNames"`
	// Verbs is a collection of actions which can be performed on the kind / resources
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	// +listType
	Verbs []string `json:"verbs"`
}

// RoleStatus defines the observed state of Role
// +k8s:openapi-gen=true
type RoleStatus struct {
	// Status provides a description of the state of this resource
	Status metav1.StatusReason `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Role is the Schema for the roles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=roles
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoleSpec   `json:"spec,omitempty"`
	Status RoleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleList contains a list of Role
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

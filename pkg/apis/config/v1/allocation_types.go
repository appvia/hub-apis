/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
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

const (
	// AllTeams is a special group name
	AllTeams = "*"
)

// AllocationSpec defines the desired state of Allocation
// +k8s:openapi-gen=true
type AllocationSpec struct {
	// Class is the reference to the provider of this class
	// +kubebuilder:validation:Required
	ClassRef corev1.Ownership `json:"classRef,omitempty"`
	// Instance is a reference to the configuration object
	// +kubebuilder:validation:Required
	InstanceRef corev1.Ownership `json:"instanceRef,omitempty"`
	// Teams is a collection of teams the allocation is to reside
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	// +listType
	Teams []string `json:"teams,omitempty"`
}

// AllocationStatus defines the observed state of Allocation
// +k8s:openapi-gen=true
type AllocationStatus struct {
	// Status is the general status of the resource
	Status corev1.Status `json:"status,omitempty"`
	// Conditions is a collection of potential issues
	// +listType
	Conditions []corev1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Allocation is the Schema for the allocations API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=allocations,scope=Namespaced
type Allocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AllocationSpec   `json:"spec,omitempty"`
	Status AllocationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AllocationList contains a list of Allocation
type AllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Allocation `json:"items"`
}

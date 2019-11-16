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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AllocationSpec defines the desired state of Allocation
// +k8s:openapi-gen=true
type AllocationSpec struct {
	// ClassRef is the reference to the provider of this class
	// +kubebuilder:validation:Required
	ClassRef metav1.GroupVersion `json:"classRef"`
	// InstanceRef is a reference to the configuration object
	// +kubebuilder:validation:Required
	InstanceRef Ownership `json:"instanceRef"`
	// Ownership provides optional ownership to the binding
	Owner *Ownership `json:"owner"`
}

// AllocationStatus defines the observed state of Allocation
// +k8s:openapi-gen=true
type AllocationStatus struct{}

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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AllocationList contains a list of Allocation
type AllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Allocation `json:"items"`
}

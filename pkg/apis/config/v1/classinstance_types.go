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

// ClassInstanceSpec defines the desired state of ClassInstance
// +k8s:openapi-gen=true
type ClassInstanceSpec struct {
	// ClassRef is the reference to the provider of this class
	// +kubebuilder:validation:Required
	ClassRef metav1.GroupKind `json:"classRef"`
	// ConfigurationRef is a reference to the configuration object
	ConfigurationRef metav1.GroupKind `json:"configurationRef"`
}

// ClassInstanceStatus defines the observed state of ClassInstance
// +k8s:openapi-gen=true
type ClassInstanceStatus struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClassInstance is the Schema for the classinstances API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=classinstances,scope=Namespaced
type ClassInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClassInstanceSpec   `json:"spec,omitempty"`
	Status ClassInstanceStatus `json:"status,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClassInstanceList contains a list of ClassInstance
type ClassInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClassInstance `json:"items"`
}

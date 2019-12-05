/*
 * Copyright (C) 2019 Rohith Jayawardene <gambol99@gmail.com>
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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// AuthProviderSpec defines the desired state of AuthProvider
// +k8s:openapi-gen=true
type AuthProviderSpec struct {
	// ClientID
	// +kubebuilder:validation:Optional
	ClientID string `json:"clientID"`
	// ClientSecret
	// +kubebuilder:validation:Optional
	ClientSecret string `json:"clientSecret"`
	// DisplayName
	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName"`
}

// AuthProviderStatus defines the observed state of OAuthProvider
// +k8s:openapi-gen=true
type AuthProviderStatus struct {
	// Conditions is a set of condition which has caused an error
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []Condition `json:"conditions"`
	// Status is overall status of the workspace
	Status Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthProvider is the Schema for the class API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=auth,scope=Namespaced
type AuthProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthProviderSpec   `json:"spec,omitempty"`
	Status AuthProviderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthProviderList contains a list of AuthProvider
type AuthProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthProvider `json:"items"`
}

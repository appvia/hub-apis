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

// CredentialsSpec defines the desired state of Credentials
// +k8s:openapi-gen=true
type CredentialsSpec struct {
	// CaCertificate is the base64 encoded cluster certificate
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	CaCertificate string `json:"caCertificate"`
	// Endpoint is the kubernetes endpoint url
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`
	// Token is the hub-admin service account token which is bound to cluster-admin
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Token string `json:"token"`
}

// CredentialsStatus defines the observed state of Credentials
// +k8s:openapi-gen=true
type CredentialsStatus struct{}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Credentials is the Schema for the roles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clustercredentials
type Credentials struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CredentialsSpec   `json:"spec,omitempty"`
	Status CredentialsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CredentialsList contains a list of Credentials
type CredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credentials `json:"items"`
}

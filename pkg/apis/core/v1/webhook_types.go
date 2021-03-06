/*
 * Copyright (C) 2019 Rohith Jayawardene <info@appvia.io>
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

// WebHookSpec defines the desired state of WebHook
// +k8s:openapi-gen=true
type WebHookSpec struct {
	// CaBundle is a ca bundle if required
	// +kubebuilder:validation:Optional
	CaBundle string `json:"caBundle"`
	// Description is a description for the team
	// +kubebuilder:validation:Required
	Description string `json:"description"`
	// Endpoint is the endpoint the service should proxy on
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`
}

// WebHookStatus defines the observed state of WebHook
// +k8s:openapi-gen=true
type WebHookStatus struct {
	// Conditions is a collection of possible errors
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []Condition `json:"conditions"`
	// Status is the status of the resource
	Status Status `json:"status"`
	// URL is the generate webhook which the application can use
	URL string `json:"url"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebHook is the Schema for the teams API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=webhooks
type WebHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebHookSpec   `json:"spec,omitempty"`
	Status WebHookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebHookList contains a list of WebHook
type WebHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebHook `json:"items"`
}

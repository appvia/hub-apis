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

import (
	corev1 "github.com/appvia/hub-apis/pkg/apis/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OrganizationSpec defines the desired state of Organization
// +k8s:openapi-gen=true
type OrganizationSpec struct {
	// Name is the name of the github organization.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	Name string `json:"name"`
	// ApplicationID is the github application id.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	ApplicationID string `json:"applicationID"`
	// ApplicationPrivateKey is the generate private key for the Github application.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	ApplicationPrivateKey string `json:"applicationPrivateKey"`
	// ApplicationInstallID is the installation id of the Github application.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	ApplicationInstallID string `json:"applicationInstallID"`
	// ApplicationClientID is the application client id for the installed application.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	ApplicationClientID string `json:"applicationClientID"`
	// ApplicationClientSecret is the application client secret from the installed applicaton.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=3
	ApplicationClientSecret string `json:"applicationClientSecret"`
}

// OrganizationStatus defines the observed state of Organization
// +k8s:openapi-gen=true
type OrganizationStatus struct {
	// Conditions is a collection of possible errors
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []corev1.Condition `json:"conditions"`
	// Status is the status of the resource
	Status corev1.Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Organization is the Schema for the teams API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=organizations
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationSpec   `json:"spec,omitempty"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrganizationList contains a list of github organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

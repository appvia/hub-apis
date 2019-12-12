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

/*
 * TODO: Need to list the identity providers back to the client by creating
 * an endpoint with the list by reflecting the types used by the
 */

// IDPConfig represents a configuration required for any Identity Provider available
// Only a single identity provider config should be set
type IDPConfig struct {
	// Google represents a Google IDP config
	// +optional
	Google *GoogleIDP `json:"google,omitempty"`
	SAML   *SAMLIDP   `json:"saml,omitempty"`
}

// IDPSpec defines the spec for a configured instance of an IDP
// +k8s:openapi-gen=true
type IDPSpec struct {
	// DisplayName
	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName"`
	// IDPConfig
	// +kubebuilder:validation:Required
	Config IDPConfig `json:"config"`
}

// IDPStatus defines the observed state of an IDP (ID Providers)
// +k8s:openapi-gen=true
type IDPStatus struct {
	// Conditions is a set of condition which has caused an error
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []Condition `json:"conditions"`
	// Status is overall status of the IDP configuration
	Status Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IDP is the Schema for the class API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=idp,scope=Namespaced
type IDP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IDPSpec   `json:"spec,omitempty"`
	Status IDPStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IDPList contains a list of IDProvider
type IDPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IDP `json:"items"`
}

// GoogleIDP provides config for a Google Identity provider
type GoogleIDP struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	Domains      []string `json:"domains"`
}

// SAMLIDP provides configuration for a generic SAML Identity provider
type SAMLIDP struct {
	SSOURL       string `json:"ssoURL"`
	CAPEM        string `json:"caPEM"`
	UsernameAttr string `json:"usernameAttr"`
	EmailAttr    string `json:"emailAttr"`
	GroupsAttr   string `json:"groupsAttr"`
}

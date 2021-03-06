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

// TeamMembershipSpec defines the desired state of Team
// +k8s:openapi-gen=true
type TeamMembershipSpec struct {
	// Username is the user being bound to the team
	// +kubebuilder:validation:Required
	Username string `json:"username"`
	// Team is the name of the team
	// +kubebuilder:validation:Required
	Team string `json:"team"`
}

// TeamMembershipStatus defines the observed state of Team
// +k8s:openapi-gen=true
type TeamMembershipStatus struct {
	// Conditions is a collection of possible errors
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []corev1.Condition `json:"conditions"`
	// Status is the status of the resource
	Status corev1.Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TeamMembership is the Schema for the teams API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=teammemberships
// +kubebuilder:printcolumn:name="Team",type="string",JSONPath=".spec.team",description="The name of the hub team"
// +kubebuilder:printcolumn:name="User",type="string",JSONPath=".spec.username",description="The name of the user whom has been granted access"
type TeamMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamMembershipSpec   `json:"spec,omitempty"`
	Status TeamMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TeamMembershipList contains a list of Team
type TeamMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamMembership `json:"items"`
}

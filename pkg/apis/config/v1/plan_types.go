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

const (
	// ClassAPILabel is the class label
	ClassAPILabel = "config.hub.appvia.io/class"
	// ClassKindLabel is the class kind label
	ClassKindLabel = "config.hub.appvia.io/kind"
)

// PlanSpec defines the desired state of Plan
// +k8s:openapi-gen=true
type PlanSpec struct {
	// Description provides a summary of the configuration provided by this plan
	// +kubebuilder:validation:MinLength=1
	Description string `json:"description,omitempty"`
	// Group is the api group this plan is providing values for
	Group metav1.GroupVersionKind `json:"group,omitempty"`
	// Summary provides a short title summary for the plan
	// +kubebuilder:validation:MinLength=1
	Summary string `json:"summary,omitempty"`
	// Values is a set of default values
	Values map[string]string `json:"values,omitempty"`
}

// PlanStatus defines the observed state of Plan
// +k8s:openapi-gen=true
type PlanStatus struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Plan is the Schema for the plans API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=plans
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanSpec   `json:"spec,omitempty"`
	Status PlanStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanList contains a list of Plan
type PlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plan `json:"items"`
}

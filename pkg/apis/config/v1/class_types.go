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
	common "k8s.io/kube-openapi/pkg/common"
)

// ClassScope defines the scope of a resources from a provider
type ClassScope string

const (
	// TeamScope indicates the resources in the hub namespace i.e toplevel
	TeamScope ClassScope = "team"
	// WorkspaceScope indicates a resource in a space
	WorkspaceScope ClassScope = "workspace"
)

// ClassSpec defines the desired state of Class
// +k8s:openapi-gen=true
type ClassSpec struct {
	// Category provides a category for this class / resurce offering
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// +required
	Category string `json:"category"`
	// Description provides a summary of what the class it offering
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// +required
	Description string `json:"description"`
	// DefaultPlan defines a default values plan for this offering
	// +optional
	DefaultPlan string `json:"defaultPlan"`
	// Requires provides a means to idenitity a relationship between the class and
	// the configuration
	// +optional
	Requires metav1.GroupVersionKind `json:"requires"`
	// Plans is a collection of default values for this class the initial one being
	// default in plans.config.hub.appvia.io/v1
	// +kubebuilder:validation:MinItems=1
	// +optional
	// +listType
	Plans []string `json:"plans"`
	// Resources is a list of any resources this class provides
	// +kubebuilder:validation:MinItems=1
	// +listType
	Resources []ClassResource `json:"resources"`
	// Schemas is OpenAPI schema for the resources
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	Schemas map[string]common.OpenAPIDefinition `json:"schemas"`
	// Summary provides a one one summary of the offering
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// +required
	Summary string `json:"summary"`
}

// ClassResource is a type of resource a class provides
type ClassResource struct {
	// Group is the apigroup the resource lives under
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// +required
	APIGroup string `json:"apiGroup"`
	// Kind is the name of the resource under the group
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// +required
	Kind string `json:"kind"`
	// Version is the apigroup version of the kind
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Version string `json:"version"`
	// Scope is the scope the resource lives under
	// +kubebuilder:validation:Required
	Scope ClassScope `json:"scope"`
	// Plans is a collection of plans for this resource
	// +kubebuilder:validation:Optional
	// +optional
	// +listType
	Plans []string `json:"plans"`
}

// ClassStatus defines the observed state of Class
// +k8s:openapi-gen=true
type ClassStatus struct{}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Class is the Schema for the classes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=classes
type Class struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClassSpec   `json:"spec,omitempty"`
	Status ClassStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClassList contains a list of Class
type ClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Class `json:"items"`
}

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
	corev1 "github.com/appvia/hub-apis/pkg/apis/core/v1"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClassScope defines the scope of a resources from a provider
type ClassScope string

const (
	// AllScope indicates no scope at all
	AllScope = ""
	// TeamScope indicates the resources in the hub namespace i.e toplevel
	TeamScope ClassScope = "team"
	// WorkspaceScope indicates a resource in a space
	WorkspaceScope ClassScope = "workspace"
)

// ClassSpec defines the desired state of Class
// +k8s:openapi-gen=true
type ClassSpec struct {
	// APIVersion is the api group and version for this class
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:Required
	APIVersion string `json:"apiVersion"`
	// AutoProvision indicates this class can be auto-provisioned based on
	// on a team pre-existing team binding
	AutoProvision bool `json:"autoProvision,omitempty"`
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
	// DisplayName is the title of the provider
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName"`
	// Requires provides a means to idenitity a relationship between the class and
	// the configuration
	// +optional
	Requires metav1.GroupVersionKind `json:"requires"`
	// Plans is a collection of default values for this class the initial one being
	// default in plans.config.hub.appvia.io/v1
	// +kubebuilder:validation:Optional
	// +listType
	Plans []string `json:"plans"`
	// Resources is a list of any resources this class provides
	// +kubebuilder:validation:MinItems=1
	// +listType
	Resources []ClassResource `json:"resources"`
	// Schemas is OpenAPI schema for the resources
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=object
	Schemas apiextv1.JSON `json:"schemas"`
}

// ClassResource is a type of resource a class provides
type ClassResource struct {
	// Group is the apigroup the resource lives under
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Group string `json:"group"`
	// DisplayName is a short name for the resource
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	DisplayName string `json:"displayName"`
	// Kind is the name of the resource under the group
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// LongDescription provides longer details as to what the resource
	// does, how it's used etc etc
	LongDescription string `json:"longDescription"`
	// Scope is the scope the resource lives under
	// +kubebuilder:validation:Required
	Scope ClassScope `json:"scope"`
	// ShortDescription provides a paragraph to description was the
	// resource is used for
	ShortDescription string `json:"shortDescription"`
	// Plans is a collection of plans for this resource
	// +kubebuilder:validation:Optional
	// +listType
	Plans []string `json:"plans"`
	// Version is the apigroup version of the kind
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Version string `json:"version"`
}

// ClassStatus defines the observed state of Class
// +k8s:openapi-gen=true
type ClassStatus struct {
	// Conditions is a set of condition which has caused an error
	// +kubebuilder:validation:Optional
	// +listType
	Conditions []corev1.Condition `json:"conditions"`
	// Status is overall status of the workspace
	Status corev1.Status `json:"status"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Class is the Schema for the classes API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=classes
// +kubebuilder:printcolumn:name="API Group",type="string",JSONPath=".spec.apiVersion",description="The API Group of the class provider"
// +kubebuilder:printcolumn:name="Provisioned",type="boolean",JSONPath=".spec.autoProvision",description="Indicates if the resource if auto provisioned"
// +kubebuilder:printcolumn:name="Display Name",type="string",JSONPath=".spec.displayName",description="The display name of the class"
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

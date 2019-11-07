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
	configv1 "github.com/appvia/hub-apis/pkg/apis/config/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubernetesSpec defines the desired state of Cluster
// +k8s:openapi-gen=true
type KubernetesSpec struct {
	// Ownership is the cloud cluster provider type for this kubernetes
	// +kubebuilder:validation:Required
	Ownership configv1.Ownership `json:"ownership"`
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

// KubernetesStatus defines the observed state of Cluster
// +k8s:openapi-gen=true
type KubernetesStatus struct{}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kubernetes is the Schema for the roles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=kubernetes
// +kubebuilder:printcolumn:name="Provider",type="string",JSONPath=".spec.ownership.group",description="The cloud provider apigroup used to provision the cluster"
// +kubebuilder:printcolumn:name="Endpoint",type="string" ,JSONPath=".spec.endpoint",description="The kube-apiserver endpoint for this cluster"
type Kubernetes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubernetesSpec   `json:"spec,omitempty"`
	Status KubernetesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubernetesList contains a list of Cluster
type KubernetesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kubernetes `json:"items"`
}

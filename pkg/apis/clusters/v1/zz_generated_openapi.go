// +build !ignore_autogenerated

/*
Copyright 2019 Rohith Jayawardene <info@appvia.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appvia/hub-apis/pkg/apis/clusters/v1.Kubernetes":       schema_pkg_apis_clusters_v1_Kubernetes(ref),
		"github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesSpec":   schema_pkg_apis_clusters_v1_KubernetesSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesStatus": schema_pkg_apis_clusters_v1_KubernetesStatus(ref),
	}
}

func schema_pkg_apis_clusters_v1_Kubernetes(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kubernetes is the Schema for the roles API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesSpec", "github.com/appvia/hub-apis/pkg/apis/clusters/v1.KubernetesStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_clusters_v1_KubernetesSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubernetesSpec defines the desired state of Cluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"use": {
						SchemaProps: spec.SchemaProps{
							Description: "Use is the cloud cluster provider type for this kubernetes",
							Ref:         ref("github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"),
						},
					},
					"caCertificate": {
						SchemaProps: spec.SchemaProps{
							Description: "CaCertificate is the base64 encoded cluster certificate",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"domain": {
						SchemaProps: spec.SchemaProps{
							Description: "Domain is the domain of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint is the kubernetes endpoint url",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"token": {
						SchemaProps: spec.SchemaProps{
							Description: "Token is the hub-admin service account token which is bound to cluster-admin",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"use"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_clusters_v1_KubernetesStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubernetesStatus defines the observed state of Cluster",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a set of condition which has caused an error",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"),
									},
								},
							},
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase indicates the phase of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "phase", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

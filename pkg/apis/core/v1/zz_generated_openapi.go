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
		"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition":     schema_pkg_apis_core_v1_Condition(ref),
		"github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership":     schema_pkg_apis_core_v1_Ownership(ref),
		"github.com/appvia/hub-apis/pkg/apis/core/v1.WebHook":       schema_pkg_apis_core_v1_WebHook(ref),
		"github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookSpec":   schema_pkg_apis_core_v1_WebHookSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookStatus": schema_pkg_apis_core_v1_WebHookStatus(ref),
	}
}

func schema_pkg_apis_core_v1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Condition is a reason why something failed",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is a human readable message",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"detail": {
						SchemaProps: spec.SchemaProps{
							Description: "Detail is a actual error which might contain technical reference",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"code": {
						SchemaProps: spec.SchemaProps{
							Description: "Code is machine readable code of the error type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"message", "detail", "code"},
			},
		},
	}
}

func schema_pkg_apis_core_v1_Ownership(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Ownership indicates the ownership of a resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion is the apigroup the resource lives under",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is the name of the resource under the group",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace is the location of the object",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is name of the resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"apiVersion", "kind", "namespace", "name"},
			},
		},
	}
}

func schema_pkg_apis_core_v1_WebHook(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebHook is the Schema for the teams API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookSpec", "github.com/appvia/hub-apis/pkg/apis/core/v1.WebHookStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_core_v1_WebHookSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebHookSpec defines the desired state of WebHook",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary is a summary name for this team",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a description for the team",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Description: "Service is the kubernetes namespaces",
							Ref:         ref("github.com/appvia/hub-apis/pkg/apis/core/v1.ServiceSpec"),
						},
					},
				},
				Required: []string{"summary", "description", "service"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.ServiceSpec"},
	}
}

func schema_pkg_apis_core_v1_WebHookStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WebHookStatus defines the observed state of WebHook",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a collection of possible errors",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the status of the resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Description: "URL is the generate webhook which the application can use",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status", "url"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}
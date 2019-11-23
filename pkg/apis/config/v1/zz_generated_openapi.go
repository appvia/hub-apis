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
		"github.com/appvia/hub-apis/pkg/apis/config/v1.Binding":       schema_pkg_apis_config_v1_Binding(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.BindingSpec":   schema_pkg_apis_config_v1_BindingSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.BindingStatus": schema_pkg_apis_config_v1_BindingStatus(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.Class":         schema_pkg_apis_config_v1_Class(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.ClassSpec":     schema_pkg_apis_config_v1_ClassSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.ClassStatus":   schema_pkg_apis_config_v1_ClassStatus(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.Plan":          schema_pkg_apis_config_v1_Plan(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.PlanSpec":      schema_pkg_apis_config_v1_PlanSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/config/v1.PlanStatus":    schema_pkg_apis_config_v1_PlanStatus(ref),
	}
}

func schema_pkg_apis_config_v1_Binding(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Binding is the Schema for the class API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.BindingSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.BindingStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/config/v1.BindingSpec", "github.com/appvia/hub-apis/pkg/apis/config/v1.BindingStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_BindingSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BindingSpec defines the desired state of Binding",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of binding",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"class": {
						SchemaProps: spec.SchemaProps{
							Description: "Class is the reference to the provider of this class",
							Ref:         ref("github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"),
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource is a reference to a resource",
							Ref:         ref("github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"),
						},
					},
					"ref": {
						SchemaProps: spec.SchemaProps{
							Description: "Ref is a reference to the configuration object",
							Ref:         ref("github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"),
						},
					},
				},
				Required: []string{"type", "class", "resource", "ref"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_config_v1_BindingStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BindingStatus defines the observed state of Binding",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_config_v1_Class(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Class is the Schema for the classes API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.ClassSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.ClassStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/config/v1.ClassSpec", "github.com/appvia/hub-apis/pkg/apis/config/v1.ClassStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_ClassSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClassSpec defines the desired state of Class",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion is the api group and version for this class",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"autoProvision": {
						SchemaProps: spec.SchemaProps{
							Description: "AutoProvision indicates this class can be auto-provisioned based on on a team pre-existing team binding",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"category": {
						SchemaProps: spec.SchemaProps{
							Description: "Category provides a category for this class / resurce offering",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description provides a summary of what the class it offering",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName is the title of the provider",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"requires": {
						SchemaProps: spec.SchemaProps{
							Description: "Requires provides a means to idenitity a relationship between the class and the configuration",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionKind"),
						},
					},
					"plans": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Plans is a collection of default values for this class the initial one being default in plans.config.hub.appvia.io/v1",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"resources": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Resources is a list of any resources this class provides",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.ClassResource"),
									},
								},
							},
						},
					},
					"schemas": {
						SchemaProps: spec.SchemaProps{
							Description: "Schemas is OpenAPI schema for the resources",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON"),
						},
					},
				},
				Required: []string{"apiVersion", "category", "description", "displayName", "resources", "schemas"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/config/v1.ClassResource", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON", "k8s.io/apimachinery/pkg/apis/meta/v1.GroupVersionKind"},
	}
}

func schema_pkg_apis_config_v1_ClassStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClassStatus defines the observed state of Class",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_config_v1_Plan(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Plan is the Schema for the plans API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.PlanSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/config/v1.PlanStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/config/v1.PlanSpec", "github.com/appvia/hub-apis/pkg/apis/config/v1.PlanStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_PlanSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanSpec defines the desired state of Plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description provides a summary of the configuration provided by this plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"values": {
						SchemaProps: spec.SchemaProps{
							Description: "Values are the key values to the plan",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON"),
						},
					},
				},
				Required: []string{"values"},
			},
		},
		Dependencies: []string{
			"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON"},
	}
}

func schema_pkg_apis_config_v1_PlanStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanStatus defines the observed state of Plan",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

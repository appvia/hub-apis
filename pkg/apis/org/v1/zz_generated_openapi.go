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
		"github.com/appvia/hub-apis/pkg/apis/org/v1.Team":                 schema_pkg_apis_org_v1_Team(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembership":       schema_pkg_apis_org_v1_TeamMembership(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipSpec":   schema_pkg_apis_org_v1_TeamMembershipSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipStatus": schema_pkg_apis_org_v1_TeamMembershipStatus(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamSpec":             schema_pkg_apis_org_v1_TeamSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamStatus":           schema_pkg_apis_org_v1_TeamStatus(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.User":                 schema_pkg_apis_org_v1_User(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.UserSpec":             schema_pkg_apis_org_v1_UserSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.UserStatus":           schema_pkg_apis_org_v1_UserStatus(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.Workspace":            schema_pkg_apis_org_v1_Workspace(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceSpec":        schema_pkg_apis_org_v1_WorkspaceSpec(ref),
		"github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceStatus":      schema_pkg_apis_org_v1_WorkspaceStatus(ref),
	}
}

func schema_pkg_apis_org_v1_Team(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Team is the Schema for the teams API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.TeamSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.TeamStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamSpec", "github.com/appvia/hub-apis/pkg/apis/org/v1.TeamStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_org_v1_TeamMembership(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TeamMembership is the Schema for the teams API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipSpec", "github.com/appvia/hub-apis/pkg/apis/org/v1.TeamMembershipStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_org_v1_TeamMembershipSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TeamMembershipSpec defines the desired state of Team",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"username": {
						SchemaProps: spec.SchemaProps{
							Description: "Username is the user being bound to the team",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"team": {
						SchemaProps: spec.SchemaProps{
							Description: "Team is the name of the team",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"username", "team"},
			},
		},
	}
}

func schema_pkg_apis_org_v1_TeamMembershipStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TeamMembershipStatus defines the observed state of Team",
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
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_org_v1_TeamSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TeamSpec defines the desired state of Team",
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
				},
				Required: []string{"summary", "description"},
			},
		},
	}
}

func schema_pkg_apis_org_v1_TeamStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TeamStatus defines the observed state of Team",
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
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace is the namespace the team is mapped to",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status", "namespace"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_org_v1_User(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "User is the Schema for the users API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.UserSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.UserStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/org/v1.UserSpec", "github.com/appvia/hub-apis/pkg/apis/org/v1.UserStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_org_v1_UserSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserSpec defines the desired state of User",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"username": {
						SchemaProps: spec.SchemaProps{
							Description: "Username is the userame or identity for this user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"username"},
			},
		},
	}
}

func schema_pkg_apis_org_v1_UserStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserStatus defines the observed state of User",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is collection of potentials error causes",
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
							Description: "Status provides an overview of the user status",
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

func schema_pkg_apis_org_v1_Workspace(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Workspace is the Schema for the workspaces API",
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
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceSpec", "github.com/appvia/hub-apis/pkg/apis/org/v1.WorkspaceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_org_v1_WorkspaceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceSpec defines the desired state of Workspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary is a short title summary to the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a desciption for the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"summary", "description"},
			},
		},
	}
}

func schema_pkg_apis_org_v1_WorkspaceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceStatus defines the observed state of Workspace",
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

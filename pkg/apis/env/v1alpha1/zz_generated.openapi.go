// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnv":       schema_pkg_apis_env_v1alpha1_ClusterEnv(ref),
		"github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvSpec":   schema_pkg_apis_env_v1alpha1_ClusterEnvSpec(ref),
		"github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvStatus": schema_pkg_apis_env_v1alpha1_ClusterEnvStatus(ref),
	}
}

func schema_pkg_apis_env_v1alpha1_ClusterEnv(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterEnv is the Schema for the clusterenvs API",
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
							Ref: ref("github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvSpec", "github.com/jmckind/env-operator/pkg/apis/env/v1alpha1.ClusterEnvStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_env_v1alpha1_ClusterEnvSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterEnvSpec defines the desired state of ClusterEnv",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_env_v1alpha1_ClusterEnvStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ClusterEnvStatus defines the observed state of ClusterEnv",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

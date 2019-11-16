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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/appvia/hub-apis/pkg/apis/config/v1"
	"github.com/appvia/hub-apis/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ConfigV1Interface interface {
	RESTClient() rest.Interface
	AllocationListsGetter
	BindingListsGetter
	ClassesGetter
}

// ConfigV1Client is used to interact with features provided by the config.hub.appvia.io group.
type ConfigV1Client struct {
	restClient rest.Interface
}

func (c *ConfigV1Client) AllocationLists(namespace string) AllocationListInterface {
	return newAllocationLists(c, namespace)
}

func (c *ConfigV1Client) BindingLists(namespace string) BindingListInterface {
	return newBindingLists(c, namespace)
}

func (c *ConfigV1Client) Classes() ClassInterface {
	return newClasses(c)
}

// NewForConfig creates a new ConfigV1Client for the given config.
func NewForConfig(c *rest.Config) (*ConfigV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ConfigV1Client{client}, nil
}

// NewForConfigOrDie creates a new ConfigV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ConfigV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ConfigV1Client for the given RESTClient.
func New(c rest.Interface) *ConfigV1Client {
	return &ConfigV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ConfigV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

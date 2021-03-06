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

package fake

import (
	v1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/core/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) AuthProviders(namespace string) v1.AuthProviderInterface {
	return &FakeAuthProviders{c, namespace}
}

func (c *FakeCoreV1) IDPs(namespace string) v1.IDPInterface {
	return &FakeIDPs{c, namespace}
}

func (c *FakeCoreV1) IDPClients(namespace string) v1.IDPClientInterface {
	return &FakeIDPClients{c, namespace}
}

func (c *FakeCoreV1) WebHooks(namespace string) v1.WebHookInterface {
	return &FakeWebHooks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

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
	clientset "github.com/appvia/hub-apis/pkg/client/clientset/versioned"
	clustersv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/clusters/v1"
	fakeclustersv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/clusters/v1/fake"
	configv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/config/v1"
	fakeconfigv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/config/v1/fake"
	corev1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/core/v1"
	fakecorev1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/core/v1/fake"
	githubv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/github/v1"
	fakegithubv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/github/v1/fake"
	orgv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/org/v1"
	fakeorgv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/org/v1/fake"
	rbacv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/rbac/v1"
	fakerbacv1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/rbac/v1/fake"
	storev1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/store/v1"
	fakestorev1 "github.com/appvia/hub-apis/pkg/client/clientset/versioned/typed/store/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ClustersV1 retrieves the ClustersV1Client
func (c *Clientset) ClustersV1() clustersv1.ClustersV1Interface {
	return &fakeclustersv1.FakeClustersV1{Fake: &c.Fake}
}

// ConfigV1 retrieves the ConfigV1Client
func (c *Clientset) ConfigV1() configv1.ConfigV1Interface {
	return &fakeconfigv1.FakeConfigV1{Fake: &c.Fake}
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// GithubV1 retrieves the GithubV1Client
func (c *Clientset) GithubV1() githubv1.GithubV1Interface {
	return &fakegithubv1.FakeGithubV1{Fake: &c.Fake}
}

// OrgV1 retrieves the OrgV1Client
func (c *Clientset) OrgV1() orgv1.OrgV1Interface {
	return &fakeorgv1.FakeOrgV1{Fake: &c.Fake}
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return &fakerbacv1.FakeRbacV1{Fake: &c.Fake}
}

// StoreV1 retrieves the StoreV1Client
func (c *Clientset) StoreV1() storev1.StoreV1Interface {
	return &fakestorev1.FakeStoreV1{Fake: &c.Fake}
}

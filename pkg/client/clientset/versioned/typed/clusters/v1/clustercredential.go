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
	"time"

	v1 "github.com/appvia/hub-apis/pkg/apis/clusters/v1"
	scheme "github.com/appvia/hub-apis/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterCredentialsGetter has a method to return a ClusterCredentialInterface.
// A group's client should implement this interface.
type ClusterCredentialsGetter interface {
	ClusterCredentials(namespace string) ClusterCredentialInterface
}

// ClusterCredentialInterface has methods to work with ClusterCredential resources.
type ClusterCredentialInterface interface {
	Create(*v1.ClusterCredential) (*v1.ClusterCredential, error)
	Update(*v1.ClusterCredential) (*v1.ClusterCredential, error)
	UpdateStatus(*v1.ClusterCredential) (*v1.ClusterCredential, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterCredential, error)
	List(opts metav1.ListOptions) (*v1.ClusterCredentialList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterCredential, err error)
	ClusterCredentialExpansion
}

// clusterCredentials implements ClusterCredentialInterface
type clusterCredentials struct {
	client rest.Interface
	ns     string
}

// newClusterCredentials returns a ClusterCredentials
func newClusterCredentials(c *ClustersV1Client, namespace string) *clusterCredentials {
	return &clusterCredentials{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterCredential, and returns the corresponding clusterCredential object, and an error if there is any.
func (c *clusterCredentials) Get(name string, options metav1.GetOptions) (result *v1.ClusterCredential, err error) {
	result = &v1.ClusterCredential{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustercredentials").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterCredentials that match those selectors.
func (c *clusterCredentials) List(opts metav1.ListOptions) (result *v1.ClusterCredentialList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterCredentialList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustercredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterCredentials.
func (c *clusterCredentials) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustercredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterCredential and creates it.  Returns the server's representation of the clusterCredential, and an error, if there is any.
func (c *clusterCredentials) Create(clusterCredential *v1.ClusterCredential) (result *v1.ClusterCredential, err error) {
	result = &v1.ClusterCredential{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustercredentials").
		Body(clusterCredential).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterCredential and updates it. Returns the server's representation of the clusterCredential, and an error, if there is any.
func (c *clusterCredentials) Update(clusterCredential *v1.ClusterCredential) (result *v1.ClusterCredential, err error) {
	result = &v1.ClusterCredential{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustercredentials").
		Name(clusterCredential.Name).
		Body(clusterCredential).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterCredentials) UpdateStatus(clusterCredential *v1.ClusterCredential) (result *v1.ClusterCredential, err error) {
	result = &v1.ClusterCredential{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustercredentials").
		Name(clusterCredential.Name).
		SubResource("status").
		Body(clusterCredential).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterCredential and deletes it. Returns an error if one occurs.
func (c *clusterCredentials) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustercredentials").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterCredentials) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustercredentials").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterCredential.
func (c *clusterCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterCredential, err error) {
	result = &v1.ClusterCredential{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustercredentials").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

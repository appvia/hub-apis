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

// CredentialsesGetter has a method to return a CredentialsInterface.
// A group's client should implement this interface.
type CredentialsesGetter interface {
	Credentialses(namespace string) CredentialsInterface
}

// CredentialsInterface has methods to work with Credentials resources.
type CredentialsInterface interface {
	Create(*v1.Credentials) (*v1.Credentials, error)
	Update(*v1.Credentials) (*v1.Credentials, error)
	UpdateStatus(*v1.Credentials) (*v1.Credentials, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Credentials, error)
	List(opts metav1.ListOptions) (*v1.CredentialsList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Credentials, err error)
	CredentialsExpansion
}

// credentialses implements CredentialsInterface
type credentialses struct {
	client rest.Interface
	ns     string
}

// newCredentialses returns a Credentialses
func newCredentialses(c *ClustersV1Client, namespace string) *credentialses {
	return &credentialses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the credentials, and returns the corresponding credentials object, and an error if there is any.
func (c *credentialses) Get(name string, options metav1.GetOptions) (result *v1.Credentials, err error) {
	result = &v1.Credentials{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("credentialses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Credentialses that match those selectors.
func (c *credentialses) List(opts metav1.ListOptions) (result *v1.CredentialsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CredentialsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("credentialses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested credentialses.
func (c *credentialses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("credentialses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a credentials and creates it.  Returns the server's representation of the credentials, and an error, if there is any.
func (c *credentialses) Create(credentials *v1.Credentials) (result *v1.Credentials, err error) {
	result = &v1.Credentials{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("credentialses").
		Body(credentials).
		Do().
		Into(result)
	return
}

// Update takes the representation of a credentials and updates it. Returns the server's representation of the credentials, and an error, if there is any.
func (c *credentialses) Update(credentials *v1.Credentials) (result *v1.Credentials, err error) {
	result = &v1.Credentials{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("credentialses").
		Name(credentials.Name).
		Body(credentials).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *credentialses) UpdateStatus(credentials *v1.Credentials) (result *v1.Credentials, err error) {
	result = &v1.Credentials{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("credentialses").
		Name(credentials.Name).
		SubResource("status").
		Body(credentials).
		Do().
		Into(result)
	return
}

// Delete takes name of the credentials and deletes it. Returns an error if one occurs.
func (c *credentialses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("credentialses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *credentialses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("credentialses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched credentials.
func (c *credentialses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Credentials, err error) {
	result = &v1.Credentials{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("credentialses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

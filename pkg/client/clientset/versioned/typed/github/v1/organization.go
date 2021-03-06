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

	v1 "github.com/appvia/hub-apis/pkg/apis/github/v1"
	scheme "github.com/appvia/hub-apis/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrganizationsGetter has a method to return a OrganizationInterface.
// A group's client should implement this interface.
type OrganizationsGetter interface {
	Organizations(namespace string) OrganizationInterface
}

// OrganizationInterface has methods to work with Organization resources.
type OrganizationInterface interface {
	Create(*v1.Organization) (*v1.Organization, error)
	Update(*v1.Organization) (*v1.Organization, error)
	UpdateStatus(*v1.Organization) (*v1.Organization, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Organization, error)
	List(opts metav1.ListOptions) (*v1.OrganizationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Organization, err error)
	OrganizationExpansion
}

// organizations implements OrganizationInterface
type organizations struct {
	client rest.Interface
	ns     string
}

// newOrganizations returns a Organizations
func newOrganizations(c *GithubV1Client, namespace string) *organizations {
	return &organizations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the organization, and returns the corresponding organization object, and an error if there is any.
func (c *organizations) Get(name string, options metav1.GetOptions) (result *v1.Organization, err error) {
	result = &v1.Organization{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Organizations that match those selectors.
func (c *organizations) List(opts metav1.ListOptions) (result *v1.OrganizationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OrganizationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested organizations.
func (c *organizations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("organizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a organization and creates it.  Returns the server's representation of the organization, and an error, if there is any.
func (c *organizations) Create(organization *v1.Organization) (result *v1.Organization, err error) {
	result = &v1.Organization{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("organizations").
		Body(organization).
		Do().
		Into(result)
	return
}

// Update takes the representation of a organization and updates it. Returns the server's representation of the organization, and an error, if there is any.
func (c *organizations) Update(organization *v1.Organization) (result *v1.Organization, err error) {
	result = &v1.Organization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizations").
		Name(organization.Name).
		Body(organization).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *organizations) UpdateStatus(organization *v1.Organization) (result *v1.Organization, err error) {
	result = &v1.Organization{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizations").
		Name(organization.Name).
		SubResource("status").
		Body(organization).
		Do().
		Into(result)
	return
}

// Delete takes name of the organization and deletes it. Returns an error if one occurs.
func (c *organizations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *organizations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched organization.
func (c *organizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Organization, err error) {
	result = &v1.Organization{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("organizations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

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

	v1 "github.com/appvia/hub-apis/pkg/apis/core/v1"
	scheme "github.com/appvia/hub-apis/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WebHooksGetter has a method to return a WebHookInterface.
// A group's client should implement this interface.
type WebHooksGetter interface {
	WebHooks(namespace string) WebHookInterface
}

// WebHookInterface has methods to work with WebHook resources.
type WebHookInterface interface {
	Create(*v1.WebHook) (*v1.WebHook, error)
	Update(*v1.WebHook) (*v1.WebHook, error)
	UpdateStatus(*v1.WebHook) (*v1.WebHook, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.WebHook, error)
	List(opts metav1.ListOptions) (*v1.WebHookList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.WebHook, err error)
	WebHookExpansion
}

// webHooks implements WebHookInterface
type webHooks struct {
	client rest.Interface
	ns     string
}

// newWebHooks returns a WebHooks
func newWebHooks(c *CoreV1Client, namespace string) *webHooks {
	return &webHooks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the webHook, and returns the corresponding webHook object, and an error if there is any.
func (c *webHooks) Get(name string, options metav1.GetOptions) (result *v1.WebHook, err error) {
	result = &v1.WebHook{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webhooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WebHooks that match those selectors.
func (c *webHooks) List(opts metav1.ListOptions) (result *v1.WebHookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.WebHookList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("webhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested webHooks.
func (c *webHooks) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("webhooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a webHook and creates it.  Returns the server's representation of the webHook, and an error, if there is any.
func (c *webHooks) Create(webHook *v1.WebHook) (result *v1.WebHook, err error) {
	result = &v1.WebHook{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("webhooks").
		Body(webHook).
		Do().
		Into(result)
	return
}

// Update takes the representation of a webHook and updates it. Returns the server's representation of the webHook, and an error, if there is any.
func (c *webHooks) Update(webHook *v1.WebHook) (result *v1.WebHook, err error) {
	result = &v1.WebHook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webhooks").
		Name(webHook.Name).
		Body(webHook).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *webHooks) UpdateStatus(webHook *v1.WebHook) (result *v1.WebHook, err error) {
	result = &v1.WebHook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("webhooks").
		Name(webHook.Name).
		SubResource("status").
		Body(webHook).
		Do().
		Into(result)
	return
}

// Delete takes name of the webHook and deletes it. Returns an error if one occurs.
func (c *webHooks) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webhooks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *webHooks) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("webhooks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched webHook.
func (c *webHooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.WebHook, err error) {
	result = &v1.WebHook{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("webhooks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

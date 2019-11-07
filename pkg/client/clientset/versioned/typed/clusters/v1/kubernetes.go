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

// KubernetesesGetter has a method to return a KubernetesInterface.
// A group's client should implement this interface.
type KubernetesesGetter interface {
	Kuberneteses(namespace string) KubernetesInterface
}

// KubernetesInterface has methods to work with Kubernetes resources.
type KubernetesInterface interface {
	Create(*v1.Kubernetes) (*v1.Kubernetes, error)
	Update(*v1.Kubernetes) (*v1.Kubernetes, error)
	UpdateStatus(*v1.Kubernetes) (*v1.Kubernetes, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Kubernetes, error)
	List(opts metav1.ListOptions) (*v1.KubernetesList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Kubernetes, err error)
	KubernetesExpansion
}

// kuberneteses implements KubernetesInterface
type kuberneteses struct {
	client rest.Interface
	ns     string
}

// newKuberneteses returns a Kuberneteses
func newKuberneteses(c *ClustersV1Client, namespace string) *kuberneteses {
	return &kuberneteses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubernetes, and returns the corresponding kubernetes object, and an error if there is any.
func (c *kuberneteses) Get(name string, options metav1.GetOptions) (result *v1.Kubernetes, err error) {
	result = &v1.Kubernetes{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kuberneteses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Kuberneteses that match those selectors.
func (c *kuberneteses) List(opts metav1.ListOptions) (result *v1.KubernetesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KubernetesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kuberneteses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kuberneteses.
func (c *kuberneteses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kuberneteses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kubernetes and creates it.  Returns the server's representation of the kubernetes, and an error, if there is any.
func (c *kuberneteses) Create(kubernetes *v1.Kubernetes) (result *v1.Kubernetes, err error) {
	result = &v1.Kubernetes{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kuberneteses").
		Body(kubernetes).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubernetes and updates it. Returns the server's representation of the kubernetes, and an error, if there is any.
func (c *kuberneteses) Update(kubernetes *v1.Kubernetes) (result *v1.Kubernetes, err error) {
	result = &v1.Kubernetes{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kuberneteses").
		Name(kubernetes.Name).
		Body(kubernetes).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kuberneteses) UpdateStatus(kubernetes *v1.Kubernetes) (result *v1.Kubernetes, err error) {
	result = &v1.Kubernetes{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kuberneteses").
		Name(kubernetes.Name).
		SubResource("status").
		Body(kubernetes).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubernetes and deletes it. Returns an error if one occurs.
func (c *kuberneteses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kuberneteses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kuberneteses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kuberneteses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubernetes.
func (c *kuberneteses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Kubernetes, err error) {
	result = &v1.Kubernetes{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kuberneteses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

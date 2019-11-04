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
	configv1 "github.com/gambol99/hub-apis/pkg/apis/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClasses implements ClassInterface
type FakeClasses struct {
	Fake *FakeConfigV1
}

var classesResource = schema.GroupVersionResource{Group: "config.hub.appvia.io", Version: "v1", Resource: "classes"}

var classesKind = schema.GroupVersionKind{Group: "config.hub.appvia.io", Version: "v1", Kind: "Class"}

// Get takes name of the class, and returns the corresponding class object, and an error if there is any.
func (c *FakeClasses) Get(name string, options v1.GetOptions) (result *configv1.Class, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(classesResource, name), &configv1.Class{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Class), err
}

// List takes label and field selectors, and returns the list of Classes that match those selectors.
func (c *FakeClasses) List(opts v1.ListOptions) (result *configv1.ClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(classesResource, classesKind, opts), &configv1.ClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.ClassList{ListMeta: obj.(*configv1.ClassList).ListMeta}
	for _, item := range obj.(*configv1.ClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested classes.
func (c *FakeClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(classesResource, opts))
}

// Create takes the representation of a class and creates it.  Returns the server's representation of the class, and an error, if there is any.
func (c *FakeClasses) Create(class *configv1.Class) (result *configv1.Class, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(classesResource, class), &configv1.Class{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Class), err
}

// Update takes the representation of a class and updates it. Returns the server's representation of the class, and an error, if there is any.
func (c *FakeClasses) Update(class *configv1.Class) (result *configv1.Class, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(classesResource, class), &configv1.Class{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Class), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClasses) UpdateStatus(class *configv1.Class) (*configv1.Class, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(classesResource, "status", class), &configv1.Class{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Class), err
}

// Delete takes name of the class and deletes it. Returns an error if one occurs.
func (c *FakeClasses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(classesResource, name), &configv1.Class{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(classesResource, listOptions)

	_, err := c.Fake.Invokes(action, &configv1.ClassList{})
	return err
}

// Patch applies the patch and returns the patched class.
func (c *FakeClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *configv1.Class, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(classesResource, name, pt, data, subresources...), &configv1.Class{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Class), err
}

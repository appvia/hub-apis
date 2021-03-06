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
	storev1 "github.com/appvia/hub-apis/pkg/apis/store/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEntities implements EntityInterface
type FakeEntities struct {
	Fake *FakeStoreV1
	ns   string
}

var entitiesResource = schema.GroupVersionResource{Group: "store.hub.appvia.io", Version: "v1", Resource: "entities"}

var entitiesKind = schema.GroupVersionKind{Group: "store.hub.appvia.io", Version: "v1", Kind: "Entity"}

// Get takes name of the entity, and returns the corresponding entity object, and an error if there is any.
func (c *FakeEntities) Get(name string, options v1.GetOptions) (result *storev1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(entitiesResource, c.ns, name), &storev1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*storev1.Entity), err
}

// List takes label and field selectors, and returns the list of Entities that match those selectors.
func (c *FakeEntities) List(opts v1.ListOptions) (result *storev1.EntityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(entitiesResource, entitiesKind, c.ns, opts), &storev1.EntityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storev1.EntityList{ListMeta: obj.(*storev1.EntityList).ListMeta}
	for _, item := range obj.(*storev1.EntityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested entities.
func (c *FakeEntities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(entitiesResource, c.ns, opts))

}

// Create takes the representation of a entity and creates it.  Returns the server's representation of the entity, and an error, if there is any.
func (c *FakeEntities) Create(entity *storev1.Entity) (result *storev1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(entitiesResource, c.ns, entity), &storev1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*storev1.Entity), err
}

// Update takes the representation of a entity and updates it. Returns the server's representation of the entity, and an error, if there is any.
func (c *FakeEntities) Update(entity *storev1.Entity) (result *storev1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(entitiesResource, c.ns, entity), &storev1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*storev1.Entity), err
}

// Delete takes name of the entity and deletes it. Returns an error if one occurs.
func (c *FakeEntities) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(entitiesResource, c.ns, name), &storev1.Entity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEntities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(entitiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &storev1.EntityList{})
	return err
}

// Patch applies the patch and returns the patched entity.
func (c *FakeEntities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *storev1.Entity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entitiesResource, c.ns, name, pt, data, subresources...), &storev1.Entity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*storev1.Entity), err
}

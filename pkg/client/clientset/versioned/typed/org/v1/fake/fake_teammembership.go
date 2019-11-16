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
	orgv1 "github.com/appvia/hub-apis/pkg/apis/org/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTeamMemberships implements TeamMembershipInterface
type FakeTeamMemberships struct {
	Fake *FakeOrgV1
	ns   string
}

var teammembershipsResource = schema.GroupVersionResource{Group: "org.hub.appvia.io", Version: "v1", Resource: "teammemberships"}

var teammembershipsKind = schema.GroupVersionKind{Group: "org.hub.appvia.io", Version: "v1", Kind: "TeamMembership"}

// Get takes name of the teamMembership, and returns the corresponding teamMembership object, and an error if there is any.
func (c *FakeTeamMemberships) Get(name string, options v1.GetOptions) (result *orgv1.TeamMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(teammembershipsResource, c.ns, name), &orgv1.TeamMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*orgv1.TeamMembership), err
}

// List takes label and field selectors, and returns the list of TeamMemberships that match those selectors.
func (c *FakeTeamMemberships) List(opts v1.ListOptions) (result *orgv1.TeamMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(teammembershipsResource, teammembershipsKind, c.ns, opts), &orgv1.TeamMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &orgv1.TeamMembershipList{ListMeta: obj.(*orgv1.TeamMembershipList).ListMeta}
	for _, item := range obj.(*orgv1.TeamMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested teamMemberships.
func (c *FakeTeamMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(teammembershipsResource, c.ns, opts))

}

// Create takes the representation of a teamMembership and creates it.  Returns the server's representation of the teamMembership, and an error, if there is any.
func (c *FakeTeamMemberships) Create(teamMembership *orgv1.TeamMembership) (result *orgv1.TeamMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(teammembershipsResource, c.ns, teamMembership), &orgv1.TeamMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*orgv1.TeamMembership), err
}

// Update takes the representation of a teamMembership and updates it. Returns the server's representation of the teamMembership, and an error, if there is any.
func (c *FakeTeamMemberships) Update(teamMembership *orgv1.TeamMembership) (result *orgv1.TeamMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(teammembershipsResource, c.ns, teamMembership), &orgv1.TeamMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*orgv1.TeamMembership), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTeamMemberships) UpdateStatus(teamMembership *orgv1.TeamMembership) (*orgv1.TeamMembership, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(teammembershipsResource, "status", c.ns, teamMembership), &orgv1.TeamMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*orgv1.TeamMembership), err
}

// Delete takes name of the teamMembership and deletes it. Returns an error if one occurs.
func (c *FakeTeamMemberships) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(teammembershipsResource, c.ns, name), &orgv1.TeamMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTeamMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(teammembershipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &orgv1.TeamMembershipList{})
	return err
}

// Patch applies the patch and returns the patched teamMembership.
func (c *FakeTeamMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *orgv1.TeamMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(teammembershipsResource, c.ns, name, pt, data, subresources...), &orgv1.TeamMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*orgv1.TeamMembership), err
}
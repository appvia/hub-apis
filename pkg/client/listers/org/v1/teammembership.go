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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/appvia/hub-apis/pkg/apis/org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TeamMembershipLister helps list TeamMemberships.
type TeamMembershipLister interface {
	// List lists all TeamMemberships in the indexer.
	List(selector labels.Selector) (ret []*v1.TeamMembership, err error)
	// TeamMemberships returns an object that can list and get TeamMemberships.
	TeamMemberships(namespace string) TeamMembershipNamespaceLister
	TeamMembershipListerExpansion
}

// teamMembershipLister implements the TeamMembershipLister interface.
type teamMembershipLister struct {
	indexer cache.Indexer
}

// NewTeamMembershipLister returns a new TeamMembershipLister.
func NewTeamMembershipLister(indexer cache.Indexer) TeamMembershipLister {
	return &teamMembershipLister{indexer: indexer}
}

// List lists all TeamMemberships in the indexer.
func (s *teamMembershipLister) List(selector labels.Selector) (ret []*v1.TeamMembership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TeamMembership))
	})
	return ret, err
}

// TeamMemberships returns an object that can list and get TeamMemberships.
func (s *teamMembershipLister) TeamMemberships(namespace string) TeamMembershipNamespaceLister {
	return teamMembershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TeamMembershipNamespaceLister helps list and get TeamMemberships.
type TeamMembershipNamespaceLister interface {
	// List lists all TeamMemberships in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TeamMembership, err error)
	// Get retrieves the TeamMembership from the indexer for a given namespace and name.
	Get(name string) (*v1.TeamMembership, error)
	TeamMembershipNamespaceListerExpansion
}

// teamMembershipNamespaceLister implements the TeamMembershipNamespaceLister
// interface.
type teamMembershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TeamMemberships in the indexer for a given namespace.
func (s teamMembershipNamespaceLister) List(selector labels.Selector) (ret []*v1.TeamMembership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TeamMembership))
	})
	return ret, err
}

// Get retrieves the TeamMembership from the indexer for a given namespace and name.
func (s teamMembershipNamespaceLister) Get(name string) (*v1.TeamMembership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("teammembership"), name)
	}
	return obj.(*v1.TeamMembership), nil
}

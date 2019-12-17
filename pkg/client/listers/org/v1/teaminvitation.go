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

// TeamInvitationLister helps list TeamInvitations.
type TeamInvitationLister interface {
	// List lists all TeamInvitations in the indexer.
	List(selector labels.Selector) (ret []*v1.TeamInvitation, err error)
	// TeamInvitations returns an object that can list and get TeamInvitations.
	TeamInvitations(namespace string) TeamInvitationNamespaceLister
	TeamInvitationListerExpansion
}

// teamInvitationLister implements the TeamInvitationLister interface.
type teamInvitationLister struct {
	indexer cache.Indexer
}

// NewTeamInvitationLister returns a new TeamInvitationLister.
func NewTeamInvitationLister(indexer cache.Indexer) TeamInvitationLister {
	return &teamInvitationLister{indexer: indexer}
}

// List lists all TeamInvitations in the indexer.
func (s *teamInvitationLister) List(selector labels.Selector) (ret []*v1.TeamInvitation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TeamInvitation))
	})
	return ret, err
}

// TeamInvitations returns an object that can list and get TeamInvitations.
func (s *teamInvitationLister) TeamInvitations(namespace string) TeamInvitationNamespaceLister {
	return teamInvitationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TeamInvitationNamespaceLister helps list and get TeamInvitations.
type TeamInvitationNamespaceLister interface {
	// List lists all TeamInvitations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TeamInvitation, err error)
	// Get retrieves the TeamInvitation from the indexer for a given namespace and name.
	Get(name string) (*v1.TeamInvitation, error)
	TeamInvitationNamespaceListerExpansion
}

// teamInvitationNamespaceLister implements the TeamInvitationNamespaceLister
// interface.
type teamInvitationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TeamInvitations in the indexer for a given namespace.
func (s teamInvitationNamespaceLister) List(selector labels.Selector) (ret []*v1.TeamInvitation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TeamInvitation))
	})
	return ret, err
}

// Get retrieves the TeamInvitation from the indexer for a given namespace and name.
func (s teamInvitationNamespaceLister) Get(name string) (*v1.TeamInvitation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("teaminvitation"), name)
	}
	return obj.(*v1.TeamInvitation), nil
}
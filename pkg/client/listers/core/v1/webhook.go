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
	v1 "github.com/appvia/hub-apis/pkg/apis/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebHookLister helps list WebHooks.
type WebHookLister interface {
	// List lists all WebHooks in the indexer.
	List(selector labels.Selector) (ret []*v1.WebHook, err error)
	// WebHooks returns an object that can list and get WebHooks.
	WebHooks(namespace string) WebHookNamespaceLister
	WebHookListerExpansion
}

// webHookLister implements the WebHookLister interface.
type webHookLister struct {
	indexer cache.Indexer
}

// NewWebHookLister returns a new WebHookLister.
func NewWebHookLister(indexer cache.Indexer) WebHookLister {
	return &webHookLister{indexer: indexer}
}

// List lists all WebHooks in the indexer.
func (s *webHookLister) List(selector labels.Selector) (ret []*v1.WebHook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WebHook))
	})
	return ret, err
}

// WebHooks returns an object that can list and get WebHooks.
func (s *webHookLister) WebHooks(namespace string) WebHookNamespaceLister {
	return webHookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebHookNamespaceLister helps list and get WebHooks.
type WebHookNamespaceLister interface {
	// List lists all WebHooks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.WebHook, err error)
	// Get retrieves the WebHook from the indexer for a given namespace and name.
	Get(name string) (*v1.WebHook, error)
	WebHookNamespaceListerExpansion
}

// webHookNamespaceLister implements the WebHookNamespaceLister
// interface.
type webHookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebHooks in the indexer for a given namespace.
func (s webHookNamespaceLister) List(selector labels.Selector) (ret []*v1.WebHook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WebHook))
	})
	return ret, err
}

// Get retrieves the WebHook from the indexer for a given namespace and name.
func (s webHookNamespaceLister) Get(name string) (*v1.WebHook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("webhook"), name)
	}
	return obj.(*v1.WebHook), nil
}
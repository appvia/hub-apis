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
	v1 "github.com/appvia/hub-apis/pkg/apis/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BindingLister helps list Bindings.
type BindingLister interface {
	// List lists all Bindings in the indexer.
	List(selector labels.Selector) (ret []*v1.Binding, err error)
	// Get retrieves the Binding from the index for a given name.
	Get(name string) (*v1.Binding, error)
	BindingListerExpansion
}

// bindingLister implements the BindingLister interface.
type bindingLister struct {
	indexer cache.Indexer
}

// NewBindingLister returns a new BindingLister.
func NewBindingLister(indexer cache.Indexer) BindingLister {
	return &bindingLister{indexer: indexer}
}

// List lists all Bindings in the indexer.
func (s *bindingLister) List(selector labels.Selector) (ret []*v1.Binding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Binding))
	})
	return ret, err
}

// Get retrieves the Binding from the index for a given name.
func (s *bindingLister) Get(name string) (*v1.Binding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("binding"), name)
	}
	return obj.(*v1.Binding), nil
}

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	configv1 "github.com/gambol99/hub-apis/pkg/apis/config/v1"
	versioned "github.com/gambol99/hub-apis/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gambol99/hub-apis/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/gambol99/hub-apis/pkg/client/listers/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClassInformer provides access to a shared informer and lister for
// Classes.
type ClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClassLister
}

type classInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClassInformer constructs a new informer for Class type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClassInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClassInformer constructs a new informer for Class type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClassInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().Classes().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().Classes().Watch(options)
			},
		},
		&configv1.Class{},
		resyncPeriod,
		indexers,
	)
}

func (f *classInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *classInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1.Class{}, f.defaultInformer)
}

func (f *classInformer) Lister() v1.ClassLister {
	return v1.NewClassLister(f.Informer().GetIndexer())
}

/*
Copyright The Kubernetes Authors.

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

package v1alpha

import (
	appsv1alpha "edgenet/pkg/apis/apps/v1alpha"
	versioned "edgenet/pkg/client/clientset/versioned"
	internalinterfaces "edgenet/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha "edgenet/pkg/client/listers/apps/v1alpha"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SliceInformer provides access to a shared informer and lister for
// Slices.
type SliceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha.SliceLister
}

type sliceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSliceInformer constructs a new informer for Slice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSliceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSliceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSliceInformer constructs a new informer for Slice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSliceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha().Slices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha().Slices(namespace).Watch(options)
			},
		},
		&appsv1alpha.Slice{},
		resyncPeriod,
		indexers,
	)
}

func (f *sliceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSliceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sliceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha.Slice{}, f.defaultInformer)
}

func (f *sliceInformer) Lister() v1alpha.SliceLister {
	return v1alpha.NewSliceLister(f.Informer().GetIndexer())
}

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	internalclientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ServiceBindingInformer provides access to a shared informer and lister for
// ServiceBindings.
type ServiceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ServiceBindingLister
}

type serviceBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewServiceBindingInformer constructs a new informer for ServiceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceBindingInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Servicecatalog().ServiceBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Servicecatalog().ServiceBindings(namespace).Watch(options)
			},
		},
		&servicecatalog.ServiceBinding{},
		resyncPeriod,
		indexers,
	)
}

func defaultServiceBindingInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewServiceBindingInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *serviceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalog.ServiceBinding{}, defaultServiceBindingInformer)
}

func (f *serviceBindingInformer) Lister() internalversion.ServiceBindingLister {
	return internalversion.NewServiceBindingLister(f.Informer().GetIndexer())
}

/*
Copyright 2023 The OpenShift Pipelines Authors

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

package v1alpha1

import (
	"context"
	time "time"

	operatorv1alpha1 "github.com/openshift-pipelines/operator/pkg/apis/operator/v1alpha1"
	versioned "github.com/openshift-pipelines/operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift-pipelines/operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift-pipelines/operator/pkg/client/listers/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PipelinesAsCodeInformer provides access to a shared informer and lister for
// PipelinesAsCodes.
type PipelinesAsCodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PipelinesAsCodeLister
}

type pipelinesAsCodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPipelinesAsCodeInformer constructs a new informer for PipelinesAsCode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelinesAsCodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelinesAsCodeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPipelinesAsCodeInformer constructs a new informer for PipelinesAsCode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelinesAsCodeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().PipelinesAsCodes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().PipelinesAsCodes().Watch(context.TODO(), options)
			},
		},
		&operatorv1alpha1.PipelinesAsCode{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelinesAsCodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelinesAsCodeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelinesAsCodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1alpha1.PipelinesAsCode{}, f.defaultInformer)
}

func (f *pipelinesAsCodeInformer) Lister() v1alpha1.PipelinesAsCodeLister {
	return v1alpha1.NewPipelinesAsCodeLister(f.Informer().GetIndexer())
}
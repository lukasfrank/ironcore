// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apiipamv1alpha1 "github.com/ironcore-dev/ironcore/api/ipam/v1alpha1"
	internalinterfaces "github.com/ironcore-dev/ironcore/client-go/informers/externalversions/internalinterfaces"
	versioned "github.com/ironcore-dev/ironcore/client-go/ironcore/versioned"
	ipamv1alpha1 "github.com/ironcore-dev/ironcore/client-go/listers/ipam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PrefixInformer provides access to a shared informer and lister for
// Prefixes.
type PrefixInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() ipamv1alpha1.PrefixLister
}

type prefixInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPrefixInformer constructs a new informer for Prefix type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPrefixInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPrefixInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPrefixInformer constructs a new informer for Prefix type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPrefixInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IpamV1alpha1().Prefixes(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IpamV1alpha1().Prefixes(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IpamV1alpha1().Prefixes(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IpamV1alpha1().Prefixes(namespace).Watch(ctx, options)
			},
		},
		&apiipamv1alpha1.Prefix{},
		resyncPeriod,
		indexers,
	)
}

func (f *prefixInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPrefixInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *prefixInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiipamv1alpha1.Prefix{}, f.defaultInformer)
}

func (f *prefixInformer) Lister() ipamv1alpha1.PrefixLister {
	return ipamv1alpha1.NewPrefixLister(f.Informer().GetIndexer())
}

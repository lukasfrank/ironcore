/*
 * Copyright (c) 2021 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	storage "github.com/onmetal/onmetal-api/apis/storage"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeClaimLister helps list VolumeClaims.
// All objects returned here must be treated as read-only.
type VolumeClaimLister interface {
	// List lists all VolumeClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storage.VolumeClaim, err error)
	// VolumeClaims returns an object that can list and get VolumeClaims.
	VolumeClaims(namespace string) VolumeClaimNamespaceLister
	VolumeClaimListerExpansion
}

// volumeClaimLister implements the VolumeClaimLister interface.
type volumeClaimLister struct {
	indexer cache.Indexer
}

// NewVolumeClaimLister returns a new VolumeClaimLister.
func NewVolumeClaimLister(indexer cache.Indexer) VolumeClaimLister {
	return &volumeClaimLister{indexer: indexer}
}

// List lists all VolumeClaims in the indexer.
func (s *volumeClaimLister) List(selector labels.Selector) (ret []*storage.VolumeClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storage.VolumeClaim))
	})
	return ret, err
}

// VolumeClaims returns an object that can list and get VolumeClaims.
func (s *volumeClaimLister) VolumeClaims(namespace string) VolumeClaimNamespaceLister {
	return volumeClaimNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeClaimNamespaceLister helps list and get VolumeClaims.
// All objects returned here must be treated as read-only.
type VolumeClaimNamespaceLister interface {
	// List lists all VolumeClaims in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storage.VolumeClaim, err error)
	// Get retrieves the VolumeClaim from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storage.VolumeClaim, error)
	VolumeClaimNamespaceListerExpansion
}

// volumeClaimNamespaceLister implements the VolumeClaimNamespaceLister
// interface.
type volumeClaimNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeClaims in the indexer for a given namespace.
func (s volumeClaimNamespaceLister) List(selector labels.Selector) (ret []*storage.VolumeClaim, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*storage.VolumeClaim))
	})
	return ret, err
}

// Get retrieves the VolumeClaim from the indexer for a given namespace and name.
func (s volumeClaimNamespaceLister) Get(name string) (*storage.VolumeClaim, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storage.Resource("volumeclaim"), name)
	}
	return obj.(*storage.VolumeClaim), nil
}
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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/onmetal/onmetal-api/apis/ipam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrefixAllocations implements PrefixAllocationInterface
type FakePrefixAllocations struct {
	Fake *FakeIpamV1alpha1
	ns   string
}

var prefixallocationsResource = schema.GroupVersionResource{Group: "ipam.api.onmetal.de", Version: "v1alpha1", Resource: "prefixallocations"}

var prefixallocationsKind = schema.GroupVersionKind{Group: "ipam.api.onmetal.de", Version: "v1alpha1", Kind: "PrefixAllocation"}

// Get takes name of the prefixAllocation, and returns the corresponding prefixAllocation object, and an error if there is any.
func (c *FakePrefixAllocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrefixAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prefixallocationsResource, c.ns, name), &v1alpha1.PrefixAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// List takes label and field selectors, and returns the list of PrefixAllocations that match those selectors.
func (c *FakePrefixAllocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrefixAllocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prefixallocationsResource, prefixallocationsKind, c.ns, opts), &v1alpha1.PrefixAllocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrefixAllocationList{ListMeta: obj.(*v1alpha1.PrefixAllocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrefixAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prefixAllocations.
func (c *FakePrefixAllocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prefixallocationsResource, c.ns, opts))

}

// Create takes the representation of a prefixAllocation and creates it.  Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *FakePrefixAllocations) Create(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.CreateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prefixallocationsResource, c.ns, prefixAllocation), &v1alpha1.PrefixAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// Update takes the representation of a prefixAllocation and updates it. Returns the server's representation of the prefixAllocation, and an error, if there is any.
func (c *FakePrefixAllocations) Update(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (result *v1alpha1.PrefixAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prefixallocationsResource, c.ns, prefixAllocation), &v1alpha1.PrefixAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrefixAllocations) UpdateStatus(ctx context.Context, prefixAllocation *v1alpha1.PrefixAllocation, opts v1.UpdateOptions) (*v1alpha1.PrefixAllocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(prefixallocationsResource, "status", c.ns, prefixAllocation), &v1alpha1.PrefixAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

// Delete takes name of the prefixAllocation and deletes it. Returns an error if one occurs.
func (c *FakePrefixAllocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prefixallocationsResource, c.ns, name, opts), &v1alpha1.PrefixAllocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrefixAllocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prefixallocationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrefixAllocationList{})
	return err
}

// Patch applies the patch and returns the patched prefixAllocation.
func (c *FakePrefixAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrefixAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prefixallocationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrefixAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrefixAllocation), err
}

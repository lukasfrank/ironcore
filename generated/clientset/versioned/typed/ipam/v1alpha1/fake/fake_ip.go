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

// FakeIPs implements IPInterface
type FakeIPs struct {
	Fake *FakeIpamV1alpha1
	ns   string
}

var ipsResource = schema.GroupVersionResource{Group: "ipam.api.onmetal.de", Version: "v1alpha1", Resource: "ips"}

var ipsKind = schema.GroupVersionKind{Group: "ipam.api.onmetal.de", Version: "v1alpha1", Kind: "IP"}

// Get takes name of the iP, and returns the corresponding iP object, and an error if there is any.
func (c *FakeIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipsResource, c.ns, name), &v1alpha1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IP), err
}

// List takes label and field selectors, and returns the list of IPs that match those selectors.
func (c *FakeIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipsResource, ipsKind, c.ns, opts), &v1alpha1.IPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IPList{ListMeta: obj.(*v1alpha1.IPList).ListMeta}
	for _, item := range obj.(*v1alpha1.IPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPs.
func (c *FakeIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipsResource, c.ns, opts))

}

// Create takes the representation of a iP and creates it.  Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Create(ctx context.Context, iP *v1alpha1.IP, opts v1.CreateOptions) (result *v1alpha1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipsResource, c.ns, iP), &v1alpha1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IP), err
}

// Update takes the representation of a iP and updates it. Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Update(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (result *v1alpha1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipsResource, c.ns, iP), &v1alpha1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPs) UpdateStatus(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (*v1alpha1.IP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipsResource, "status", c.ns, iP), &v1alpha1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IP), err
}

// Delete takes name of the iP and deletes it. Returns an error if one occurs.
func (c *FakeIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipsResource, c.ns, name, opts), &v1alpha1.IP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IPList{})
	return err
}

// Patch applies the patch and returns the patched iP.
func (c *FakeIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IP), err
}

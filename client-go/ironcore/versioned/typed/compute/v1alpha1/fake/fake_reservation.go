// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	computev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReservations implements ReservationInterface
type FakeReservations struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var reservationsResource = v1alpha1.SchemeGroupVersion.WithResource("reservations")

var reservationsKind = v1alpha1.SchemeGroupVersion.WithKind("Reservation")

// Get takes name of the reservation, and returns the corresponding reservation object, and an error if there is any.
func (c *FakeReservations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Reservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reservationsResource, c.ns, name), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// List takes label and field selectors, and returns the list of Reservations that match those selectors.
func (c *FakeReservations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReservationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reservationsResource, reservationsKind, c.ns, opts), &v1alpha1.ReservationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReservationList{ListMeta: obj.(*v1alpha1.ReservationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReservationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reservations.
func (c *FakeReservations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reservationsResource, c.ns, opts))

}

// Create takes the representation of a reservation and creates it.  Returns the server's representation of the reservation, and an error, if there is any.
func (c *FakeReservations) Create(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.CreateOptions) (result *v1alpha1.Reservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reservationsResource, c.ns, reservation), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// Update takes the representation of a reservation and updates it. Returns the server's representation of the reservation, and an error, if there is any.
func (c *FakeReservations) Update(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (result *v1alpha1.Reservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reservationsResource, c.ns, reservation), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReservations) UpdateStatus(ctx context.Context, reservation *v1alpha1.Reservation, opts v1.UpdateOptions) (*v1alpha1.Reservation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reservationsResource, "status", c.ns, reservation), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// Delete takes name of the reservation and deletes it. Returns an error if one occurs.
func (c *FakeReservations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(reservationsResource, c.ns, name, opts), &v1alpha1.Reservation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReservations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reservationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReservationList{})
	return err
}

// Patch applies the patch and returns the patched reservation.
func (c *FakeReservations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Reservation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reservationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied reservation.
func (c *FakeReservations) Apply(ctx context.Context, reservation *computev1alpha1.ReservationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Reservation, err error) {
	if reservation == nil {
		return nil, fmt.Errorf("reservation provided to Apply must not be nil")
	}
	data, err := json.Marshal(reservation)
	if err != nil {
		return nil, err
	}
	name := reservation.Name
	if name == nil {
		return nil, fmt.Errorf("reservation.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reservationsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeReservations) ApplyStatus(ctx context.Context, reservation *computev1alpha1.ReservationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Reservation, err error) {
	if reservation == nil {
		return nil, fmt.Errorf("reservation provided to Apply must not be nil")
	}
	data, err := json.Marshal(reservation)
	if err != nil {
		return nil, err
	}
	name := reservation.Name
	if name == nil {
		return nil, fmt.Errorf("reservation.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reservationsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Reservation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Reservation), err
}
// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	computev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/compute/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore/client-go/ironcore/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineClassesGetter has a method to return a MachineClassInterface.
// A group's client should implement this interface.
type MachineClassesGetter interface {
	MachineClasses() MachineClassInterface
}

// MachineClassInterface has methods to work with MachineClass resources.
type MachineClassInterface interface {
	Create(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.CreateOptions) (*v1alpha1.MachineClass, error)
	Update(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.UpdateOptions) (*v1alpha1.MachineClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MachineClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MachineClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineClass, err error)
	Apply(ctx context.Context, machineClass *computev1alpha1.MachineClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MachineClass, err error)
	MachineClassExpansion
}

// machineClasses implements MachineClassInterface
type machineClasses struct {
	client rest.Interface
}

// newMachineClasses returns a MachineClasses
func newMachineClasses(c *ComputeV1alpha1Client) *machineClasses {
	return &machineClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineClass, and returns the corresponding machineClass object, and an error if there is any.
func (c *machineClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MachineClass, err error) {
	result = &v1alpha1.MachineClass{}
	err = c.client.Get().
		Resource("machineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineClasses that match those selectors.
func (c *machineClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MachineClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MachineClassList{}
	err = c.client.Get().
		Resource("machineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineClasses.
func (c *machineClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("machineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineClass and creates it.  Returns the server's representation of the machineClass, and an error, if there is any.
func (c *machineClasses) Create(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.CreateOptions) (result *v1alpha1.MachineClass, err error) {
	result = &v1alpha1.MachineClass{}
	err = c.client.Post().
		Resource("machineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineClass and updates it. Returns the server's representation of the machineClass, and an error, if there is any.
func (c *machineClasses) Update(ctx context.Context, machineClass *v1alpha1.MachineClass, opts v1.UpdateOptions) (result *v1alpha1.MachineClass, err error) {
	result = &v1alpha1.MachineClass{}
	err = c.client.Put().
		Resource("machineclasses").
		Name(machineClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineClass and deletes it. Returns an error if one occurs.
func (c *machineClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machineclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("machineclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineClass.
func (c *machineClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MachineClass, err error) {
	result = &v1alpha1.MachineClass{}
	err = c.client.Patch(pt).
		Resource("machineclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied machineClass.
func (c *machineClasses) Apply(ctx context.Context, machineClass *computev1alpha1.MachineClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.MachineClass, err error) {
	if machineClass == nil {
		return nil, fmt.Errorf("machineClass provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(machineClass)
	if err != nil {
		return nil, err
	}
	name := machineClass.Name
	if name == nil {
		return nil, fmt.Errorf("machineClass.Name must be provided to Apply")
	}
	result = &v1alpha1.MachineClass{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("machineclasses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/ironcore-dev/ironcore/api/compute/v1alpha1"
	computev1alpha1 "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/compute/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore/client-go/ironcore/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MachinesGetter has a method to return a MachineInterface.
// A group's client should implement this interface.
type MachinesGetter interface {
	Machines(namespace string) MachineInterface
}

// MachineInterface has methods to work with Machine resources.
type MachineInterface interface {
	Create(ctx context.Context, machine *v1alpha1.Machine, opts v1.CreateOptions) (*v1alpha1.Machine, error)
	Update(ctx context.Context, machine *v1alpha1.Machine, opts v1.UpdateOptions) (*v1alpha1.Machine, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, machine *v1alpha1.Machine, opts v1.UpdateOptions) (*v1alpha1.Machine, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Machine, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MachineList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Machine, err error)
	Apply(ctx context.Context, machine *computev1alpha1.MachineApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Machine, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, machine *computev1alpha1.MachineApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Machine, err error)
	MachineExpansion
}

// machines implements MachineInterface
type machines struct {
	*gentype.ClientWithListAndApply[*v1alpha1.Machine, *v1alpha1.MachineList, *computev1alpha1.MachineApplyConfiguration]
}

// newMachines returns a Machines
func newMachines(c *ComputeV1alpha1Client, namespace string) *machines {
	return &machines{
		gentype.NewClientWithListAndApply[*v1alpha1.Machine, *v1alpha1.MachineList, *computev1alpha1.MachineApplyConfiguration](
			"machines",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Machine { return &v1alpha1.Machine{} },
			func() *v1alpha1.MachineList { return &v1alpha1.MachineList{} }),
	}
}

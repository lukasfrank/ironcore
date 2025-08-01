// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
)

// NetworkStatusApplyConfiguration represents a declarative configuration of the NetworkStatus type for use
// with apply.
type NetworkStatusApplyConfiguration struct {
	State    *networkingv1alpha1.NetworkState         `json:"state,omitempty"`
	Peerings []NetworkPeeringStatusApplyConfiguration `json:"peerings,omitempty"`
}

// NetworkStatusApplyConfiguration constructs a declarative configuration of the NetworkStatus type for use with
// apply.
func NetworkStatus() *NetworkStatusApplyConfiguration {
	return &NetworkStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *NetworkStatusApplyConfiguration) WithState(value networkingv1alpha1.NetworkState) *NetworkStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithPeerings adds the given value to the Peerings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Peerings field.
func (b *NetworkStatusApplyConfiguration) WithPeerings(values ...*NetworkPeeringStatusApplyConfiguration) *NetworkStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeerings")
		}
		b.Peerings = append(b.Peerings, *values[i])
	}
	return b
}

// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// NetworkInterfaceSourceApplyConfiguration represents a declarative configuration of the NetworkInterfaceSource type for use
// with apply.
type NetworkInterfaceSourceApplyConfiguration struct {
	NetworkInterfaceRef *v1.LocalObjectReference                           `json:"networkInterfaceRef,omitempty"`
	Ephemeral           *EphemeralNetworkInterfaceSourceApplyConfiguration `json:"ephemeral,omitempty"`
}

// NetworkInterfaceSourceApplyConfiguration constructs a declarative configuration of the NetworkInterfaceSource type for use with
// apply.
func NetworkInterfaceSource() *NetworkInterfaceSourceApplyConfiguration {
	return &NetworkInterfaceSourceApplyConfiguration{}
}

// WithNetworkInterfaceRef sets the NetworkInterfaceRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkInterfaceRef field is set to the value of the last call.
func (b *NetworkInterfaceSourceApplyConfiguration) WithNetworkInterfaceRef(value v1.LocalObjectReference) *NetworkInterfaceSourceApplyConfiguration {
	b.NetworkInterfaceRef = &value
	return b
}

// WithEphemeral sets the Ephemeral field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ephemeral field is set to the value of the last call.
func (b *NetworkInterfaceSourceApplyConfiguration) WithEphemeral(value *EphemeralNetworkInterfaceSourceApplyConfiguration) *NetworkInterfaceSourceApplyConfiguration {
	b.Ephemeral = value
	return b
}

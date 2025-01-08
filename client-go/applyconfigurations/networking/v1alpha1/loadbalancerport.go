// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// LoadBalancerPortApplyConfiguration represents a declarative configuration of the LoadBalancerPort type for use
// with apply.
type LoadBalancerPortApplyConfiguration struct {
	Protocol *v1.Protocol `json:"protocol,omitempty"`
	Port     *int32       `json:"port,omitempty"`
	EndPort  *int32       `json:"endPort,omitempty"`
}

// LoadBalancerPortApplyConfiguration constructs a declarative configuration of the LoadBalancerPort type for use with
// apply.
func LoadBalancerPort() *LoadBalancerPortApplyConfiguration {
	return &LoadBalancerPortApplyConfiguration{}
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithProtocol(value v1.Protocol) *LoadBalancerPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithPort(value int32) *LoadBalancerPortApplyConfiguration {
	b.Port = &value
	return b
}

// WithEndPort sets the EndPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndPort field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithEndPort(value int32) *LoadBalancerPortApplyConfiguration {
	b.EndPort = &value
	return b
}

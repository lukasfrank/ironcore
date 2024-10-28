// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	networkingv1alpha1 "github.com/ironcore-dev/ironcore/api/networking/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// NetworkPolicySpecApplyConfiguration represents an declarative configuration of the NetworkPolicySpec type for use
// with apply.
type NetworkPolicySpecApplyConfiguration struct {
	NetworkRef               *v1.LocalObjectReference                     `json:"networkRef,omitempty"`
	NetworkInterfaceSelector *metav1.LabelSelectorApplyConfiguration      `json:"networkInterfaceSelector,omitempty"`
	Ingress                  []NetworkPolicyIngressRuleApplyConfiguration `json:"ingress,omitempty"`
	Egress                   []NetworkPolicyEgressRuleApplyConfiguration  `json:"egress,omitempty"`
	PolicyTypes              []networkingv1alpha1.PolicyType              `json:"policyTypes,omitempty"`
}

// NetworkPolicySpecApplyConfiguration constructs an declarative configuration of the NetworkPolicySpec type for use with
// apply.
func NetworkPolicySpec() *NetworkPolicySpecApplyConfiguration {
	return &NetworkPolicySpecApplyConfiguration{}
}

// WithNetworkRef sets the NetworkRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkRef field is set to the value of the last call.
func (b *NetworkPolicySpecApplyConfiguration) WithNetworkRef(value v1.LocalObjectReference) *NetworkPolicySpecApplyConfiguration {
	b.NetworkRef = &value
	return b
}

// WithNetworkInterfaceSelector sets the NetworkInterfaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkInterfaceSelector field is set to the value of the last call.
func (b *NetworkPolicySpecApplyConfiguration) WithNetworkInterfaceSelector(value *metav1.LabelSelectorApplyConfiguration) *NetworkPolicySpecApplyConfiguration {
	b.NetworkInterfaceSelector = value
	return b
}

// WithIngress adds the given value to the Ingress field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ingress field.
func (b *NetworkPolicySpecApplyConfiguration) WithIngress(values ...*NetworkPolicyIngressRuleApplyConfiguration) *NetworkPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIngress")
		}
		b.Ingress = append(b.Ingress, *values[i])
	}
	return b
}

// WithEgress adds the given value to the Egress field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Egress field.
func (b *NetworkPolicySpecApplyConfiguration) WithEgress(values ...*NetworkPolicyEgressRuleApplyConfiguration) *NetworkPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEgress")
		}
		b.Egress = append(b.Egress, *values[i])
	}
	return b
}

// WithPolicyTypes adds the given value to the PolicyTypes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PolicyTypes field.
func (b *NetworkPolicySpecApplyConfiguration) WithPolicyTypes(values ...networkingv1alpha1.PolicyType) *NetworkPolicySpecApplyConfiguration {
	for i := range values {
		b.PolicyTypes = append(b.PolicyTypes, values[i])
	}
	return b
}

// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/ironcore-dev/ironcore/api/core/v1alpha1"
	internal "github.com/ironcore-dev/ironcore/client-go/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ResourceQuotaApplyConfiguration represents a declarative configuration of the ResourceQuota type for use
// with apply.
type ResourceQuotaApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ResourceQuotaSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *ResourceQuotaStatusApplyConfiguration `json:"status,omitempty"`
}

// ResourceQuota constructs a declarative configuration of the ResourceQuota type for use with
// apply.
func ResourceQuota(name, namespace string) *ResourceQuotaApplyConfiguration {
	b := &ResourceQuotaApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ResourceQuota")
	b.WithAPIVersion("core.ironcore.dev/v1alpha1")
	return b
}

// ExtractResourceQuota extracts the applied configuration owned by fieldManager from
// resourceQuota. If no managedFields are found in resourceQuota for fieldManager, a
// ResourceQuotaApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// resourceQuota must be a unmodified ResourceQuota API object that was retrieved from the Kubernetes API.
// ExtractResourceQuota provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractResourceQuota(resourceQuota *corev1alpha1.ResourceQuota, fieldManager string) (*ResourceQuotaApplyConfiguration, error) {
	return extractResourceQuota(resourceQuota, fieldManager, "")
}

// ExtractResourceQuotaStatus is the same as ExtractResourceQuota except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractResourceQuotaStatus(resourceQuota *corev1alpha1.ResourceQuota, fieldManager string) (*ResourceQuotaApplyConfiguration, error) {
	return extractResourceQuota(resourceQuota, fieldManager, "status")
}

func extractResourceQuota(resourceQuota *corev1alpha1.ResourceQuota, fieldManager string, subresource string) (*ResourceQuotaApplyConfiguration, error) {
	b := &ResourceQuotaApplyConfiguration{}
	err := managedfields.ExtractInto(resourceQuota, internal.Parser().Type("com.github.ironcore-dev.ironcore.api.core.v1alpha1.ResourceQuota"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(resourceQuota.Name)
	b.WithNamespace(resourceQuota.Namespace)

	b.WithKind("ResourceQuota")
	b.WithAPIVersion("core.ironcore.dev/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithKind(value string) *ResourceQuotaApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithAPIVersion(value string) *ResourceQuotaApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithName(value string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithGenerateName(value string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithNamespace(value string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithUID(value types.UID) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithResourceVersion(value string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithGeneration(value int64) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ResourceQuotaApplyConfiguration) WithLabels(entries map[string]string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ResourceQuotaApplyConfiguration) WithAnnotations(entries map[string]string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ResourceQuotaApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ResourceQuotaApplyConfiguration) WithFinalizers(values ...string) *ResourceQuotaApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *ResourceQuotaApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithSpec(value *ResourceQuotaSpecApplyConfiguration) *ResourceQuotaApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ResourceQuotaApplyConfiguration) WithStatus(value *ResourceQuotaStatusApplyConfiguration) *ResourceQuotaApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ResourceQuotaApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	ipamv1alpha1 "github.com/onmetal/onmetal-api/apis/ipam/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&NetworkInterface{}, func(obj interface{}) { SetObjectDefaults_NetworkInterface(obj.(*NetworkInterface)) })
	scheme.AddTypeDefaultingFunc(&NetworkInterfaceList{}, func(obj interface{}) { SetObjectDefaults_NetworkInterfaceList(obj.(*NetworkInterfaceList)) })
	return nil
}

func SetObjectDefaults_NetworkInterface(in *NetworkInterface) {
	SetDefaults_NetworkInterfaceSpec(&in.Spec)
	for i := range in.Spec.IPs {
		a := &in.Spec.IPs[i]
		if a.EphemeralPrefix != nil {
			if a.EphemeralPrefix.PrefixTemplate != nil {
				ipamv1alpha1.SetDefaults_PrefixSpec(&a.EphemeralPrefix.PrefixTemplate.Spec)
			}
		}
	}
}

func SetObjectDefaults_NetworkInterfaceList(in *NetworkInterfaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_NetworkInterface(a)
	}
}
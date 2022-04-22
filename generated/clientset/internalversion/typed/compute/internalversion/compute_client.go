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

package internalversion

import (
	"net/http"

	"github.com/onmetal/onmetal-api/generated/clientset/internalversion/scheme"
	rest "k8s.io/client-go/rest"
)

type ComputeInterface interface {
	RESTClient() rest.Interface
	MachinesGetter
	MachineClassesGetter
	MachinePoolsGetter
}

// ComputeClient is used to interact with features provided by the compute.api.onmetal.de group.
type ComputeClient struct {
	restClient rest.Interface
}

func (c *ComputeClient) Machines(namespace string) MachineInterface {
	return newMachines(c, namespace)
}

func (c *ComputeClient) MachineClasses(namespace string) MachineClassInterface {
	return newMachineClasses(c, namespace)
}

func (c *ComputeClient) MachinePools() MachinePoolInterface {
	return newMachinePools(c)
}

// NewForConfig creates a new ComputeClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ComputeClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ComputeClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ComputeClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ComputeClient{client}, nil
}

// NewForConfigOrDie creates a new ComputeClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ComputeClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ComputeClient for the given RESTClient.
func New(c rest.Interface) *ComputeClient {
	return &ComputeClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("compute.api.onmetal.de")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("compute.api.onmetal.de")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ComputeClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT DIRECTLY.
package versioned

import (
	glog "github.com/golang/glog"
	ianlewisv1alpha1 "github.com/ianlewis/memcached-operator/internal/client/clientset/versioned/typed/ianlewis.org/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	IanlewisV1alpha1() ianlewisv1alpha1.IanlewisV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Ianlewis() ianlewisv1alpha1.IanlewisV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	ianlewisV1alpha1 *ianlewisv1alpha1.IanlewisV1alpha1Client
}

// IanlewisV1alpha1 retrieves the IanlewisV1alpha1Client
func (c *Clientset) IanlewisV1alpha1() ianlewisv1alpha1.IanlewisV1alpha1Interface {
	return c.ianlewisV1alpha1
}

// Deprecated: Ianlewis retrieves the default version of IanlewisClient.
// Please explicitly pick a version.
func (c *Clientset) Ianlewis() ianlewisv1alpha1.IanlewisV1alpha1Interface {
	return c.ianlewisV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.ianlewisV1alpha1, err = ianlewisv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.ianlewisV1alpha1 = ianlewisv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.ianlewisV1alpha1 = ianlewisv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}

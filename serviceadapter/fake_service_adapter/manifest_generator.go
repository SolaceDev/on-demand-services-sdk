// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated by counterfeiter
package fake_service_adapter

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
)

type FakeManifestGenerator struct {
	GenerateManifestStub        func(serviceDeployment serviceadapter.ServiceDeployment, plan serviceadapter.Plan, requestParams serviceadapter.RequestParameters, previousManifest *bosh.BoshManifest, previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error)
	generateManifestMutex       sync.RWMutex
	generateManifestArgsForCall []struct {
		serviceDeployment serviceadapter.ServiceDeployment
		plan              serviceadapter.Plan
		requestParams     serviceadapter.RequestParameters
		previousManifest  *bosh.BoshManifest
		previousPlan      *serviceadapter.Plan
	}
	generateManifestReturns struct {
		result1 bosh.BoshManifest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestGenerator) GenerateManifest(serviceDeployment serviceadapter.ServiceDeployment, plan serviceadapter.Plan, requestParams serviceadapter.RequestParameters, previousManifest *bosh.BoshManifest, previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error) {
	fake.generateManifestMutex.Lock()
	fake.generateManifestArgsForCall = append(fake.generateManifestArgsForCall, struct {
		serviceDeployment serviceadapter.ServiceDeployment
		plan              serviceadapter.Plan
		requestParams     serviceadapter.RequestParameters
		previousManifest  *bosh.BoshManifest
		previousPlan      *serviceadapter.Plan
	}{serviceDeployment, plan, requestParams, previousManifest, previousPlan})
	fake.recordInvocation("GenerateManifest", []interface{}{serviceDeployment, plan, requestParams, previousManifest, previousPlan})
	fake.generateManifestMutex.Unlock()
	if fake.GenerateManifestStub != nil {
		return fake.GenerateManifestStub(serviceDeployment, plan, requestParams, previousManifest, previousPlan)
	} else {
		return fake.generateManifestReturns.result1, fake.generateManifestReturns.result2
	}
}

func (fake *FakeManifestGenerator) GenerateManifestCallCount() int {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return len(fake.generateManifestArgsForCall)
}

func (fake *FakeManifestGenerator) GenerateManifestArgsForCall(i int) (serviceadapter.ServiceDeployment, serviceadapter.Plan, serviceadapter.RequestParameters, *bosh.BoshManifest, *serviceadapter.Plan) {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return fake.generateManifestArgsForCall[i].serviceDeployment, fake.generateManifestArgsForCall[i].plan, fake.generateManifestArgsForCall[i].requestParams, fake.generateManifestArgsForCall[i].previousManifest, fake.generateManifestArgsForCall[i].previousPlan
}

func (fake *FakeManifestGenerator) GenerateManifestReturns(result1 bosh.BoshManifest, result2 error) {
	fake.GenerateManifestStub = nil
	fake.generateManifestReturns = struct {
		result1 bosh.BoshManifest
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManifestGenerator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ serviceadapter.ManifestGenerator = new(FakeManifestGenerator)

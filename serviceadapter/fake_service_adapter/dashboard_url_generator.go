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

type FakeDashboardUrlGenerator struct {
	DashboardUrlStub        func(instanceID string, plan serviceadapter.Plan, manifest bosh.BoshManifest) (serviceadapter.DashboardUrl, error)
	dashboardUrlMutex       sync.RWMutex
	dashboardUrlArgsForCall []struct {
		instanceID string
		plan       serviceadapter.Plan
		manifest   bosh.BoshManifest
	}
	dashboardUrlReturns struct {
		result1 serviceadapter.DashboardUrl
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDashboardUrlGenerator) DashboardUrl(instanceID string, plan serviceadapter.Plan, manifest bosh.BoshManifest) (serviceadapter.DashboardUrl, error) {
	fake.dashboardUrlMutex.Lock()
	fake.dashboardUrlArgsForCall = append(fake.dashboardUrlArgsForCall, struct {
		instanceID string
		plan       serviceadapter.Plan
		manifest   bosh.BoshManifest
	}{instanceID, plan, manifest})
	fake.recordInvocation("DashboardUrl", []interface{}{instanceID, plan, manifest})
	fake.dashboardUrlMutex.Unlock()
	if fake.DashboardUrlStub != nil {
		return fake.DashboardUrlStub(instanceID, plan, manifest)
	} else {
		return fake.dashboardUrlReturns.result1, fake.dashboardUrlReturns.result2
	}
}

func (fake *FakeDashboardUrlGenerator) DashboardUrlCallCount() int {
	fake.dashboardUrlMutex.RLock()
	defer fake.dashboardUrlMutex.RUnlock()
	return len(fake.dashboardUrlArgsForCall)
}

func (fake *FakeDashboardUrlGenerator) DashboardUrlArgsForCall(i int) (string, serviceadapter.Plan, bosh.BoshManifest) {
	fake.dashboardUrlMutex.RLock()
	defer fake.dashboardUrlMutex.RUnlock()
	return fake.dashboardUrlArgsForCall[i].instanceID, fake.dashboardUrlArgsForCall[i].plan, fake.dashboardUrlArgsForCall[i].manifest
}

func (fake *FakeDashboardUrlGenerator) DashboardUrlReturns(result1 serviceadapter.DashboardUrl, result2 error) {
	fake.DashboardUrlStub = nil
	fake.dashboardUrlReturns = struct {
		result1 serviceadapter.DashboardUrl
		result2 error
	}{result1, result2}
}

func (fake *FakeDashboardUrlGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dashboardUrlMutex.RLock()
	defer fake.dashboardUrlMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDashboardUrlGenerator) recordInvocation(key string, args []interface{}) {
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

var _ serviceadapter.DashboardUrlGenerator = new(FakeDashboardUrlGenerator)

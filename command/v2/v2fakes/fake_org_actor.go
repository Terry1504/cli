// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeOrgActor struct {
	GetOrganizationByNameStub        func(orgName string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		orgName string
	}
	getOrganizationByNameReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationDomainsStub        func(orgGUID string) ([]v2action.Domain, v2action.Warnings, error)
	getOrganizationDomainsMutex       sync.RWMutex
	getOrganizationDomainsArgsForCall []struct {
		orgGUID string
	}
	getOrganizationDomainsReturns struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationQuotaStub        func(quotaGUID string) (v2action.OrganizationQuota, v2action.Warnings, error)
	getOrganizationQuotaMutex       sync.RWMutex
	getOrganizationQuotaArgsForCall []struct {
		quotaGUID string
	}
	getOrganizationQuotaReturns struct {
		result1 v2action.OrganizationQuota
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationSpacesStub        func(orgGUID string) ([]v2action.Space, v2action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		orgGUID string
	}
	getOrganizationSpacesReturns struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgActor) GetOrganizationByName(orgName string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("GetOrganizationByName", []interface{}{orgName})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(orgName)
	} else {
		return fake.getOrganizationByNameReturns.result1, fake.getOrganizationByNameReturns.result2, fake.getOrganizationByNameReturns.result3
	}
}

func (fake *FakeOrgActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return fake.getOrganizationByNameArgsForCall[i].orgName
}

func (fake *FakeOrgActor) GetOrganizationByNameReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationDomains(orgGUID string) ([]v2action.Domain, v2action.Warnings, error) {
	fake.getOrganizationDomainsMutex.Lock()
	fake.getOrganizationDomainsArgsForCall = append(fake.getOrganizationDomainsArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationDomains", []interface{}{orgGUID})
	fake.getOrganizationDomainsMutex.Unlock()
	if fake.GetOrganizationDomainsStub != nil {
		return fake.GetOrganizationDomainsStub(orgGUID)
	} else {
		return fake.getOrganizationDomainsReturns.result1, fake.getOrganizationDomainsReturns.result2, fake.getOrganizationDomainsReturns.result3
	}
}

func (fake *FakeOrgActor) GetOrganizationDomainsCallCount() int {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return len(fake.getOrganizationDomainsArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationDomainsArgsForCall(i int) string {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return fake.getOrganizationDomainsArgsForCall[i].orgGUID
}

func (fake *FakeOrgActor) GetOrganizationDomainsReturns(result1 []v2action.Domain, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationDomainsStub = nil
	fake.getOrganizationDomainsReturns = struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationQuota(quotaGUID string) (v2action.OrganizationQuota, v2action.Warnings, error) {
	fake.getOrganizationQuotaMutex.Lock()
	fake.getOrganizationQuotaArgsForCall = append(fake.getOrganizationQuotaArgsForCall, struct {
		quotaGUID string
	}{quotaGUID})
	fake.recordInvocation("GetOrganizationQuota", []interface{}{quotaGUID})
	fake.getOrganizationQuotaMutex.Unlock()
	if fake.GetOrganizationQuotaStub != nil {
		return fake.GetOrganizationQuotaStub(quotaGUID)
	} else {
		return fake.getOrganizationQuotaReturns.result1, fake.getOrganizationQuotaReturns.result2, fake.getOrganizationQuotaReturns.result3
	}
}

func (fake *FakeOrgActor) GetOrganizationQuotaCallCount() int {
	fake.getOrganizationQuotaMutex.RLock()
	defer fake.getOrganizationQuotaMutex.RUnlock()
	return len(fake.getOrganizationQuotaArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationQuotaArgsForCall(i int) string {
	fake.getOrganizationQuotaMutex.RLock()
	defer fake.getOrganizationQuotaMutex.RUnlock()
	return fake.getOrganizationQuotaArgsForCall[i].quotaGUID
}

func (fake *FakeOrgActor) GetOrganizationQuotaReturns(result1 v2action.OrganizationQuota, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationQuotaStub = nil
	fake.getOrganizationQuotaReturns = struct {
		result1 v2action.OrganizationQuota
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) GetOrganizationSpaces(orgGUID string) ([]v2action.Space, v2action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{orgGUID})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(orgGUID)
	} else {
		return fake.getOrganizationSpacesReturns.result1, fake.getOrganizationSpacesReturns.result2, fake.getOrganizationSpacesReturns.result3
	}
}

func (fake *FakeOrgActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakeOrgActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return fake.getOrganizationSpacesArgsForCall[i].orgGUID
}

func (fake *FakeOrgActor) GetOrganizationSpacesReturns(result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	fake.getOrganizationQuotaMutex.RLock()
	defer fake.getOrganizationQuotaMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOrgActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.OrgActor = new(FakeOrgActor)
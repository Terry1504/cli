// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeEnvActor struct {
	GetEnvironmentVariablesByApplicationNameAndSpaceStub        func(string, string) (v7action.EnvironmentVariableGroups, v7action.Warnings, error)
	getEnvironmentVariablesByApplicationNameAndSpaceMutex       sync.RWMutex
	getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getEnvironmentVariablesByApplicationNameAndSpaceReturns struct {
		result1 v7action.EnvironmentVariableGroups
		result2 v7action.Warnings
		result3 error
	}
	getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.EnvironmentVariableGroups
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpace(arg1 string, arg2 string) (v7action.EnvironmentVariableGroups, v7action.Warnings, error) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall[len(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall)]
	fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall = append(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetEnvironmentVariablesByApplicationNameAndSpace", []interface{}{arg1, arg2})
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Unlock()
	if fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub != nil {
		return fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceCallCount() int {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceCalls(stub func(string, string) (v7action.EnvironmentVariableGroups, v7action.Warnings, error)) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Lock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Unlock()
	fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub = stub
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceReturns(result1 v7action.EnvironmentVariableGroups, result2 v7action.Warnings, result3 error) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Lock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Unlock()
	fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub = nil
	fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns = struct {
		result1 v7action.EnvironmentVariableGroups
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall(i int, result1 v7action.EnvironmentVariableGroups, result2 v7action.Warnings, result3 error) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Lock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Unlock()
	fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub = nil
	if fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.EnvironmentVariableGroups
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.EnvironmentVariableGroups
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEnvActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.EnvActor = new(FakeEnvActor)

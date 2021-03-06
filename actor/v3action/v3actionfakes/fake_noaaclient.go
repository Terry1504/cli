// Code generated by counterfeiter. DO NOT EDIT.
package v3actionfakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	events "github.com/cloudfoundry/sonde-go/events"
)

type FakeNOAAClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	RecentLogsStub        func(string, string) ([]*events.LogMessage, error)
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	recentLogsReturns struct {
		result1 []*events.LogMessage
		result2 error
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 []*events.LogMessage
		result2 error
	}
	SetOnConnectCallbackStub        func(func())
	setOnConnectCallbackMutex       sync.RWMutex
	setOnConnectCallbackArgsForCall []struct {
		arg1 func()
	}
	TailingLogsStub        func(string, string) (<-chan *events.LogMessage, <-chan error)
	tailingLogsMutex       sync.RWMutex
	tailingLogsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	tailingLogsReturns struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	tailingLogsReturnsOnCall map[int]struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNOAAClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeNOAAClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeNOAAClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeNOAAClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNOAAClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNOAAClient) RecentLogs(arg1 string, arg2 string) ([]*events.LogMessage, error) {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("RecentLogs", []interface{}{arg1, arg2})
	fake.recentLogsMutex.Unlock()
	if fake.RecentLogsStub != nil {
		return fake.RecentLogsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recentLogsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNOAAClient) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeNOAAClient) RecentLogsCalls(stub func(string, string) ([]*events.LogMessage, error)) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = stub
}

func (fake *FakeNOAAClient) RecentLogsArgsForCall(i int) (string, string) {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	argsForCall := fake.recentLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNOAAClient) RecentLogsReturns(result1 []*events.LogMessage, result2 error) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeNOAAClient) RecentLogsReturnsOnCall(i int, result1 []*events.LogMessage, result2 error) {
	fake.recentLogsMutex.Lock()
	defer fake.recentLogsMutex.Unlock()
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 []*events.LogMessage
			result2 error
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeNOAAClient) SetOnConnectCallback(arg1 func()) {
	fake.setOnConnectCallbackMutex.Lock()
	fake.setOnConnectCallbackArgsForCall = append(fake.setOnConnectCallbackArgsForCall, struct {
		arg1 func()
	}{arg1})
	fake.recordInvocation("SetOnConnectCallback", []interface{}{arg1})
	fake.setOnConnectCallbackMutex.Unlock()
	if fake.SetOnConnectCallbackStub != nil {
		fake.SetOnConnectCallbackStub(arg1)
	}
}

func (fake *FakeNOAAClient) SetOnConnectCallbackCallCount() int {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	return len(fake.setOnConnectCallbackArgsForCall)
}

func (fake *FakeNOAAClient) SetOnConnectCallbackCalls(stub func(func())) {
	fake.setOnConnectCallbackMutex.Lock()
	defer fake.setOnConnectCallbackMutex.Unlock()
	fake.SetOnConnectCallbackStub = stub
}

func (fake *FakeNOAAClient) SetOnConnectCallbackArgsForCall(i int) func() {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	argsForCall := fake.setOnConnectCallbackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNOAAClient) TailingLogs(arg1 string, arg2 string) (<-chan *events.LogMessage, <-chan error) {
	fake.tailingLogsMutex.Lock()
	ret, specificReturn := fake.tailingLogsReturnsOnCall[len(fake.tailingLogsArgsForCall)]
	fake.tailingLogsArgsForCall = append(fake.tailingLogsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("TailingLogs", []interface{}{arg1, arg2})
	fake.tailingLogsMutex.Unlock()
	if fake.TailingLogsStub != nil {
		return fake.TailingLogsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tailingLogsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNOAAClient) TailingLogsCallCount() int {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return len(fake.tailingLogsArgsForCall)
}

func (fake *FakeNOAAClient) TailingLogsCalls(stub func(string, string) (<-chan *events.LogMessage, <-chan error)) {
	fake.tailingLogsMutex.Lock()
	defer fake.tailingLogsMutex.Unlock()
	fake.TailingLogsStub = stub
}

func (fake *FakeNOAAClient) TailingLogsArgsForCall(i int) (string, string) {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	argsForCall := fake.tailingLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNOAAClient) TailingLogsReturns(result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.tailingLogsMutex.Lock()
	defer fake.tailingLogsMutex.Unlock()
	fake.TailingLogsStub = nil
	fake.tailingLogsReturns = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNOAAClient) TailingLogsReturnsOnCall(i int, result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.tailingLogsMutex.Lock()
	defer fake.tailingLogsMutex.Unlock()
	fake.TailingLogsStub = nil
	if fake.tailingLogsReturnsOnCall == nil {
		fake.tailingLogsReturnsOnCall = make(map[int]struct {
			result1 <-chan *events.LogMessage
			result2 <-chan error
		})
	}
	fake.tailingLogsReturnsOnCall[i] = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeNOAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNOAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v3action.NOAAClient = new(FakeNOAAClient)

// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3CreateAppActor struct {
	CreateApplicationInSpaceStub        func(v3action.Application, string) (v3action.Application, v3action.Warnings, error)
	createApplicationInSpaceMutex       sync.RWMutex
	createApplicationInSpaceArgsForCall []struct {
		arg1 v3action.Application
		arg2 string
	}
	createApplicationInSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	createApplicationInSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpace(arg1 v3action.Application, arg2 string) (v3action.Application, v3action.Warnings, error) {
	fake.createApplicationInSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationInSpaceReturnsOnCall[len(fake.createApplicationInSpaceArgsForCall)]
	fake.createApplicationInSpaceArgsForCall = append(fake.createApplicationInSpaceArgsForCall, struct {
		arg1 v3action.Application
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateApplicationInSpace", []interface{}{arg1, arg2})
	fake.createApplicationInSpaceMutex.Unlock()
	if fake.CreateApplicationInSpaceStub != nil {
		return fake.CreateApplicationInSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createApplicationInSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpaceCallCount() int {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return len(fake.createApplicationInSpaceArgsForCall)
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpaceCalls(stub func(v3action.Application, string) (v3action.Application, v3action.Warnings, error)) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = stub
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpaceArgsForCall(i int) (v3action.Application, string) {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	argsForCall := fake.createApplicationInSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = nil
	fake.createApplicationInSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3CreateAppActor) CreateApplicationInSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = nil
	if fake.createApplicationInSpaceReturnsOnCall == nil {
		fake.createApplicationInSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.createApplicationInSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3CreateAppActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3CreateAppActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3CreateAppActor = new(FakeV3CreateAppActor)

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type RequestService struct {
	InvokeStub        func(api.RequestServiceInvokeInput) (api.RequestServiceInvokeOutput, error)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		arg1 api.RequestServiceInvokeInput
	}
	invokeReturns struct {
		result1 api.RequestServiceInvokeOutput
		result2 error
	}
	invokeReturnsOnCall map[int]struct {
		result1 api.RequestServiceInvokeOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RequestService) Invoke(arg1 api.RequestServiceInvokeInput) (api.RequestServiceInvokeOutput, error) {
	fake.invokeMutex.Lock()
	ret, specificReturn := fake.invokeReturnsOnCall[len(fake.invokeArgsForCall)]
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		arg1 api.RequestServiceInvokeInput
	}{arg1})
	fake.recordInvocation("Invoke", []interface{}{arg1})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.invokeReturns.result1, fake.invokeReturns.result2
}

func (fake *RequestService) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *RequestService) InvokeArgsForCall(i int) api.RequestServiceInvokeInput {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invokeArgsForCall[i].arg1
}

func (fake *RequestService) InvokeReturns(result1 api.RequestServiceInvokeOutput, result2 error) {
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 api.RequestServiceInvokeOutput
		result2 error
	}{result1, result2}
}

func (fake *RequestService) InvokeReturnsOnCall(i int, result1 api.RequestServiceInvokeOutput, result2 error) {
	fake.InvokeStub = nil
	if fake.invokeReturnsOnCall == nil {
		fake.invokeReturnsOnCall = make(map[int]struct {
			result1 api.RequestServiceInvokeOutput
			result2 error
		})
	}
	fake.invokeReturnsOnCall[i] = struct {
		result1 api.RequestServiceInvokeOutput
		result2 error
	}{result1, result2}
}

func (fake *RequestService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RequestService) recordInvocation(key string, args []interface{}) {
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
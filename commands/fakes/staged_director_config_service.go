// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StagedDirectorConfigService struct {
	GetStagedProductByNameStub        func(product string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		product string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductPropertiesStub        func(product string) (map[string]api.ResponseProperty, error)
	getStagedProductPropertiesMutex       sync.RWMutex
	getStagedProductPropertiesArgsForCall []struct {
		product string
	}
	getStagedProductPropertiesReturns struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	getStagedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	GetStagedDirectorAZStub        func() ([]interface{}, error)
	getStagedDirectorAZMutex       sync.RWMutex
	getStagedDirectorAZArgsForCall []struct{}
	getStagedDirectorAZReturns     struct {
		result1 []interface{}
		result2 error
	}
	getStagedDirectorAZReturnsOnCall map[int]struct {
		result1 []interface{}
		result2 error
	}
	GetStagedDirectorPropertiesStub        func() (map[string]interface{}, error)
	getStagedDirectorPropertiesMutex       sync.RWMutex
	getStagedDirectorPropertiesArgsForCall []struct{}
	getStagedDirectorPropertiesReturns     struct {
		result1 map[string]interface{}
		result2 error
	}
	getStagedDirectorPropertiesReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedDirectorConfigService) GetStagedProductByName(product string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetStagedProductByName", []interface{}{product})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductByNameReturns.result1, fake.getStagedProductByNameReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return fake.getStagedProductByNameArgsForCall[i].product
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductProperties(product string) (map[string]api.ResponseProperty, error) {
	fake.getStagedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedProductPropertiesReturnsOnCall[len(fake.getStagedProductPropertiesArgsForCall)]
	fake.getStagedProductPropertiesArgsForCall = append(fake.getStagedProductPropertiesArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetStagedProductProperties", []interface{}{product})
	fake.getStagedProductPropertiesMutex.Unlock()
	if fake.GetStagedProductPropertiesStub != nil {
		return fake.GetStagedProductPropertiesStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductPropertiesReturns.result1, fake.getStagedProductPropertiesReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedProductPropertiesCallCount() int {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return len(fake.getStagedProductPropertiesArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedProductPropertiesArgsForCall(i int) string {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return fake.getStagedProductPropertiesArgsForCall[i].product
}

func (fake *StagedDirectorConfigService) GetStagedProductPropertiesReturns(result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetStagedProductPropertiesStub = nil
	fake.getStagedProductPropertiesReturns = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedProductPropertiesReturnsOnCall(i int, result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetStagedProductPropertiesStub = nil
	if fake.getStagedProductPropertiesReturnsOnCall == nil {
		fake.getStagedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]api.ResponseProperty
			result2 error
		})
	}
	fake.getStagedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAZ() ([]interface{}, error) {
	fake.getStagedDirectorAZMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorAZReturnsOnCall[len(fake.getStagedDirectorAZArgsForCall)]
	fake.getStagedDirectorAZArgsForCall = append(fake.getStagedDirectorAZArgsForCall, struct{}{})
	fake.recordInvocation("GetStagedDirectorAZ", []interface{}{})
	fake.getStagedDirectorAZMutex.Unlock()
	if fake.GetStagedDirectorAZStub != nil {
		return fake.GetStagedDirectorAZStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedDirectorAZReturns.result1, fake.getStagedDirectorAZReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAZCallCount() int {
	fake.getStagedDirectorAZMutex.RLock()
	defer fake.getStagedDirectorAZMutex.RUnlock()
	return len(fake.getStagedDirectorAZArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAZReturns(result1 []interface{}, result2 error) {
	fake.GetStagedDirectorAZStub = nil
	fake.getStagedDirectorAZReturns = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorAZReturnsOnCall(i int, result1 []interface{}, result2 error) {
	fake.GetStagedDirectorAZStub = nil
	if fake.getStagedDirectorAZReturnsOnCall == nil {
		fake.getStagedDirectorAZReturnsOnCall = make(map[int]struct {
			result1 []interface{}
			result2 error
		})
	}
	fake.getStagedDirectorAZReturnsOnCall[i] = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorProperties() (map[string]interface{}, error) {
	fake.getStagedDirectorPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedDirectorPropertiesReturnsOnCall[len(fake.getStagedDirectorPropertiesArgsForCall)]
	fake.getStagedDirectorPropertiesArgsForCall = append(fake.getStagedDirectorPropertiesArgsForCall, struct{}{})
	fake.recordInvocation("GetStagedDirectorProperties", []interface{}{})
	fake.getStagedDirectorPropertiesMutex.Unlock()
	if fake.GetStagedDirectorPropertiesStub != nil {
		return fake.GetStagedDirectorPropertiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedDirectorPropertiesReturns.result1, fake.getStagedDirectorPropertiesReturns.result2
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesCallCount() int {
	fake.getStagedDirectorPropertiesMutex.RLock()
	defer fake.getStagedDirectorPropertiesMutex.RUnlock()
	return len(fake.getStagedDirectorPropertiesArgsForCall)
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesReturns(result1 map[string]interface{}, result2 error) {
	fake.GetStagedDirectorPropertiesStub = nil
	fake.getStagedDirectorPropertiesReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) GetStagedDirectorPropertiesReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.GetStagedDirectorPropertiesStub = nil
	if fake.getStagedDirectorPropertiesReturnsOnCall == nil {
		fake.getStagedDirectorPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getStagedDirectorPropertiesReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedDirectorConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	fake.getStagedDirectorAZMutex.RLock()
	defer fake.getStagedDirectorAZMutex.RUnlock()
	fake.getStagedDirectorPropertiesMutex.RLock()
	defer fake.getStagedDirectorPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedDirectorConfigService) recordInvocation(key string, args []interface{}) {
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

// Code generated by counterfeiter. DO NOT EDIT.
package mysql_fake

import (
	"crypto/tls"
	"sync"

	"code.cloudfoundry.org/goshims/mysqlshim"
	"github.com/go-sql-driver/mysql"
)

type FakeMySQL struct {
	ParseDSNStub        func(string) (*mysql.Config, error)
	parseDSNMutex       sync.RWMutex
	parseDSNArgsForCall []struct {
		arg1 string
	}
	parseDSNReturns struct {
		result1 *mysql.Config
		result2 error
	}
	parseDSNReturnsOnCall map[int]struct {
		result1 *mysql.Config
		result2 error
	}
	RegisterTLSConfigStub        func(string, *tls.Config) error
	registerTLSConfigMutex       sync.RWMutex
	registerTLSConfigArgsForCall []struct {
		arg1 string
		arg2 *tls.Config
	}
	registerTLSConfigReturns struct {
		result1 error
	}
	registerTLSConfigReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMySQL) ParseDSN(arg1 string) (*mysql.Config, error) {
	fake.parseDSNMutex.Lock()
	ret, specificReturn := fake.parseDSNReturnsOnCall[len(fake.parseDSNArgsForCall)]
	fake.parseDSNArgsForCall = append(fake.parseDSNArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ParseDSN", []interface{}{arg1})
	fake.parseDSNMutex.Unlock()
	if fake.ParseDSNStub != nil {
		return fake.ParseDSNStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.parseDSNReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMySQL) ParseDSNCallCount() int {
	fake.parseDSNMutex.RLock()
	defer fake.parseDSNMutex.RUnlock()
	return len(fake.parseDSNArgsForCall)
}

func (fake *FakeMySQL) ParseDSNCalls(stub func(string) (*mysql.Config, error)) {
	fake.parseDSNMutex.Lock()
	defer fake.parseDSNMutex.Unlock()
	fake.ParseDSNStub = stub
}

func (fake *FakeMySQL) ParseDSNArgsForCall(i int) string {
	fake.parseDSNMutex.RLock()
	defer fake.parseDSNMutex.RUnlock()
	argsForCall := fake.parseDSNArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMySQL) ParseDSNReturns(result1 *mysql.Config, result2 error) {
	fake.parseDSNMutex.Lock()
	defer fake.parseDSNMutex.Unlock()
	fake.ParseDSNStub = nil
	fake.parseDSNReturns = struct {
		result1 *mysql.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeMySQL) ParseDSNReturnsOnCall(i int, result1 *mysql.Config, result2 error) {
	fake.parseDSNMutex.Lock()
	defer fake.parseDSNMutex.Unlock()
	fake.ParseDSNStub = nil
	if fake.parseDSNReturnsOnCall == nil {
		fake.parseDSNReturnsOnCall = make(map[int]struct {
			result1 *mysql.Config
			result2 error
		})
	}
	fake.parseDSNReturnsOnCall[i] = struct {
		result1 *mysql.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeMySQL) RegisterTLSConfig(arg1 string, arg2 *tls.Config) error {
	fake.registerTLSConfigMutex.Lock()
	ret, specificReturn := fake.registerTLSConfigReturnsOnCall[len(fake.registerTLSConfigArgsForCall)]
	fake.registerTLSConfigArgsForCall = append(fake.registerTLSConfigArgsForCall, struct {
		arg1 string
		arg2 *tls.Config
	}{arg1, arg2})
	fake.recordInvocation("RegisterTLSConfig", []interface{}{arg1, arg2})
	fake.registerTLSConfigMutex.Unlock()
	if fake.RegisterTLSConfigStub != nil {
		return fake.RegisterTLSConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.registerTLSConfigReturns
	return fakeReturns.result1
}

func (fake *FakeMySQL) RegisterTLSConfigCallCount() int {
	fake.registerTLSConfigMutex.RLock()
	defer fake.registerTLSConfigMutex.RUnlock()
	return len(fake.registerTLSConfigArgsForCall)
}

func (fake *FakeMySQL) RegisterTLSConfigCalls(stub func(string, *tls.Config) error) {
	fake.registerTLSConfigMutex.Lock()
	defer fake.registerTLSConfigMutex.Unlock()
	fake.RegisterTLSConfigStub = stub
}

func (fake *FakeMySQL) RegisterTLSConfigArgsForCall(i int) (string, *tls.Config) {
	fake.registerTLSConfigMutex.RLock()
	defer fake.registerTLSConfigMutex.RUnlock()
	argsForCall := fake.registerTLSConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMySQL) RegisterTLSConfigReturns(result1 error) {
	fake.registerTLSConfigMutex.Lock()
	defer fake.registerTLSConfigMutex.Unlock()
	fake.RegisterTLSConfigStub = nil
	fake.registerTLSConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMySQL) RegisterTLSConfigReturnsOnCall(i int, result1 error) {
	fake.registerTLSConfigMutex.Lock()
	defer fake.registerTLSConfigMutex.Unlock()
	fake.RegisterTLSConfigStub = nil
	if fake.registerTLSConfigReturnsOnCall == nil {
		fake.registerTLSConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerTLSConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMySQL) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseDSNMutex.RLock()
	defer fake.parseDSNMutex.RUnlock()
	fake.registerTLSConfigMutex.RLock()
	defer fake.registerTLSConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMySQL) recordInvocation(key string, args []interface{}) {
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

var _ mysqlshim.MySQL = new(FakeMySQL)

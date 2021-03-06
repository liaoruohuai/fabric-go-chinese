
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:37</date>
//</624456128985567232>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/protos/token"
	"github.com/hyperledger/fabric/token/server"
)

type PolicyChecker struct {
	CheckStub        func(sc *token.SignedCommand, c *token.Command) error
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		sc *token.SignedCommand
		c  *token.Command
	}
	checkReturns struct {
		result1 error
	}
	checkReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyChecker) Check(sc *token.SignedCommand, c *token.Command) error {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		sc *token.SignedCommand
		c  *token.Command
	}{sc, c})
	fake.recordInvocation("Check", []interface{}{sc, c})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(sc, c)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkReturns.result1
}

func (fake *PolicyChecker) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *PolicyChecker) CheckArgsForCall(i int) (*token.SignedCommand, *token.Command) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return fake.checkArgsForCall[i].sc, fake.checkArgsForCall[i].c
}

func (fake *PolicyChecker) CheckReturns(result1 error) {
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 error
	}{result1}
}

func (fake *PolicyChecker) CheckReturnsOnCall(i int, result1 error) {
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PolicyChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyChecker) recordInvocation(key string, args []interface{}) {
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

var _ server.PolicyChecker = new(PolicyChecker)



//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456006050516992>

//伪造者生成的代码。不要编辑。
package mock

import (
	io "io"
	sync "sync"
)

type Builder struct {
	BuildStub        func() (io.Reader, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
	}
	buildReturns struct {
		result1 io.Reader
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 io.Reader
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Builder) Build() (io.Reader, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
	}{})
	fake.recordInvocation("Build", []interface{}{})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.buildReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Builder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *Builder) BuildCalls(stub func() (io.Reader, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *Builder) BuildReturns(result1 io.Reader, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *Builder) BuildReturnsOnCall(i int, result1 io.Reader, result2 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 io.Reader
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *Builder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Builder) recordInvocation(key string, args []interface{}) {
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



//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:14</date>
//</624456031648354304>

//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
)

type MembershipInfoProvider struct {
	AmMemberOfStub        func(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (bool, error)
	amMemberOfMutex       sync.RWMutex
	amMemberOfArgsForCall []struct {
		channelName            string
		collectionPolicyConfig *common.CollectionPolicyConfig
	}
	amMemberOfReturns struct {
		result1 bool
		result2 error
	}
	amMemberOfReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MembershipInfoProvider) AmMemberOf(channelName string, collectionPolicyConfig *common.CollectionPolicyConfig) (bool, error) {
	fake.amMemberOfMutex.Lock()
	ret, specificReturn := fake.amMemberOfReturnsOnCall[len(fake.amMemberOfArgsForCall)]
	fake.amMemberOfArgsForCall = append(fake.amMemberOfArgsForCall, struct {
		channelName            string
		collectionPolicyConfig *common.CollectionPolicyConfig
	}{channelName, collectionPolicyConfig})
	fake.recordInvocation("AmMemberOf", []interface{}{channelName, collectionPolicyConfig})
	fake.amMemberOfMutex.Unlock()
	if fake.AmMemberOfStub != nil {
		return fake.AmMemberOfStub(channelName, collectionPolicyConfig)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.amMemberOfReturns.result1, fake.amMemberOfReturns.result2
}

func (fake *MembershipInfoProvider) AmMemberOfCallCount() int {
	fake.amMemberOfMutex.RLock()
	defer fake.amMemberOfMutex.RUnlock()
	return len(fake.amMemberOfArgsForCall)
}

func (fake *MembershipInfoProvider) AmMemberOfArgsForCall(i int) (string, *common.CollectionPolicyConfig) {
	fake.amMemberOfMutex.RLock()
	defer fake.amMemberOfMutex.RUnlock()
	return fake.amMemberOfArgsForCall[i].channelName, fake.amMemberOfArgsForCall[i].collectionPolicyConfig
}

func (fake *MembershipInfoProvider) AmMemberOfReturns(result1 bool, result2 error) {
	fake.AmMemberOfStub = nil
	fake.amMemberOfReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MembershipInfoProvider) AmMemberOfReturnsOnCall(i int, result1 bool, result2 error) {
	fake.AmMemberOfStub = nil
	if fake.amMemberOfReturnsOnCall == nil {
		fake.amMemberOfReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.amMemberOfReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *MembershipInfoProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.amMemberOfMutex.RLock()
	defer fake.amMemberOfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MembershipInfoProvider) recordInvocation(key string, args []interface{}) {
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

var _ ledger.MembershipInfoProvider = new(MembershipInfoProvider)


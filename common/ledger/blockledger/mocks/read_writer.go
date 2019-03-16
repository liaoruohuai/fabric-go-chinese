
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455960542318592>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import blockledger "github.com/hyperledger/fabric/common/ledger/blockledger"
import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//
type ReadWriter struct {
	mock.Mock
}

//append提供了一个具有给定字段的模拟函数：block
func (_m *ReadWriter) Append(block *common.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//Height provides a mock function with given fields:
func (_m *ReadWriter) Height() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

//
func (_m *ReadWriter) Iterator(startType *orderer.SeekPosition) (blockledger.Iterator, uint64) {
	ret := _m.Called(startType)

	var r0 blockledger.Iterator
	if rf, ok := ret.Get(0).(func(*orderer.SeekPosition) blockledger.Iterator); ok {
		r0 = rf(startType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockledger.Iterator)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*orderer.SeekPosition) uint64); ok {
		r1 = rf(startType)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}


//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455977579581440>

//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import cc "github.com/hyperledger/fabric/core/cclifecycle"
import mock "github.com/stretchr/testify/mock"

//QueryCreator是为QueryCreator类型自动生成的模拟类型
type QueryCreator struct {
	mock.Mock
}

//NewQuery提供了一个具有给定字段的模拟函数：
func (_m *QueryCreator) NewQuery() (cc.Query, error) {
	ret := _m.Called()

	var r0 cc.Query
	if rf, ok := ret.Get(0).(func() cc.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cc.Query)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}


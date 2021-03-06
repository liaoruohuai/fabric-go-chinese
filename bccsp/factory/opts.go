
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455926388101120>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/

package factory

//GetDefaultOpts为Opts提供默认实现
//每次返回一个新实例
func GetDefaultOpts() *FactoryOpts {
	return &FactoryOpts{
		ProviderName: "SW",
		SwOpts: &SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,

			Ephemeral: true,
		},
	}
}

//FactoryName返回提供程序的名称
func (o *FactoryOpts) FactoryName() string {
	return o.ProviderName
}


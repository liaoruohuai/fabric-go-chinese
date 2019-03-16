
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:00</date>
//</624455971510423552>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package sanitycheck

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	newchannelconfig "github.com/hyperledger/fabric/common/channelconfig"
	cb "github.com/hyperledger/fabric/protos/common"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
)

type Messages struct {
	GeneralErrors   []string          `json:"general_errors"`
	ElementWarnings []*ElementMessage `json:"element_warnings"`
	ElementErrors   []*ElementMessage `json:"element_errors"`
}

type ElementMessage struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

func Check(config *cb.Config) (*Messages, error) {
	result := &Messages{}

	bundle, err := newchannelconfig.NewBundle("sanitycheck", config)
	if err != nil {
		result.GeneralErrors = []string{err.Error()}
		return result, nil
	}

//这应该来自MSP经理，但出于某种原因
//如果没有组织，MSP管理器就不会初始化，所以，
//我们是手工收集的。
	mspMap := make(map[string]struct{})

	if ac, ok := bundle.ApplicationConfig(); ok {
		for _, org := range ac.Organizations() {
			mspMap[org.MSPID()] = struct{}{}
		}
	}

	if oc, ok := bundle.OrdererConfig(); ok {
		for _, org := range oc.Organizations() {
			mspMap[org.MSPID()] = struct{}{}
		}
	}

	policyWarnings, policyErrors := checkPolicyPrincipals(config.ChannelGroup, "", mspMap)

	result.ElementWarnings = policyWarnings
	result.ElementErrors = policyErrors

	return result, nil
}

func checkPolicyPrincipals(group *cb.ConfigGroup, basePath string, mspMap map[string]struct{}) (warnings []*ElementMessage, errors []*ElementMessage) {
	for policyName, configPolicy := range group.Policies {
		appendError := func(err string) {
			errors = append(errors, &ElementMessage{
				Path:    basePath + ".policies." + policyName,
				Message: err,
			})
		}

		appendWarning := func(err string) {
			warnings = append(errors, &ElementMessage{
				Path:    basePath + ".policies." + policyName,
				Message: err,
			})
		}

		if configPolicy.Policy == nil {
			appendError(fmt.Sprintf("no policy value set for %s", policyName))
			continue
		}

		if configPolicy.Policy.Type != int32(cb.Policy_SIGNATURE) {
			continue
		}
		spe := &cb.SignaturePolicyEnvelope{}
		err := proto.Unmarshal(configPolicy.Policy.Value, spe)
		if err != nil {
			appendError(fmt.Sprintf("error unmarshaling policy value to SignaturePolicyEnvelope: %s", err))
			continue
		}

		for i, identity := range spe.Identities {
			var mspID string
			switch identity.PrincipalClassification {
			case mspprotos.MSPPrincipal_ROLE:
				role := &mspprotos.MSPRole{}
				err = proto.Unmarshal(identity.Principal, role)
				if err != nil {
					appendError(fmt.Sprintf("value of identities array at index %d is of type ROLE, but could not be unmarshaled to msp.MSPRole: %s", i, err))
					continue
				}
				mspID = role.MspIdentifier
			case mspprotos.MSPPrincipal_ORGANIZATION_UNIT:
				ou := &mspprotos.OrganizationUnit{}
				err = proto.Unmarshal(identity.Principal, ou)
				if err != nil {
					appendError(fmt.Sprintf("value of identities array at index %d is of type ORGANIZATION_UNIT, but could not be unmarshaled to msp.OrganizationUnit: %s", i, err))
					continue
				}
				mspID = ou.MspIdentifier
			default:
				continue
			}

			_, ok := mspMap[mspID]
			if !ok {
				appendWarning(fmt.Sprintf("identity principal at index %d refers to MSP ID '%s', which is not an MSP in the network", i, mspID))
			}
		}
	}

	for subGroupName, subGroup := range group.Groups {
		subWarnings, subErrors := checkPolicyPrincipals(subGroup, basePath+".groups."+subGroupName, mspMap)
		warnings = append(warnings, subWarnings...)
		errors = append(errors, subErrors...)
	}
	return
}

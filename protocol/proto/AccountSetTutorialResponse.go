// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AccountSetTutorialResponse struct{
	message ProtoMessage
	ResponsePacket

    
}

func (x *AccountSetTutorialResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountSetTutorialResponse) ProtoReflect() Message {
	return x
}

func (x *AccountSetTutorialResponse) GetProtocol() int32 {
	return mx.Protocol_Account_SetTutorial
}

func (x *AccountSetTutorialResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


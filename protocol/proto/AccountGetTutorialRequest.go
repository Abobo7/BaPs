// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AccountGetTutorialRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *AccountGetTutorialRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountGetTutorialRequest) ProtoReflect() Message {
	return x
}

func (x *AccountGetTutorialRequest) GetProtocol() int32 {
	return mx.Protocol_Account_GetTutorial
}

func (x *AccountGetTutorialRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


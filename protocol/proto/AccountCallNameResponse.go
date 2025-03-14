// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AccountCallNameResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountDB *AccountDB
}

func (x *AccountCallNameResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountCallNameResponse) ProtoReflect() Message {
	return x
}

func (x *AccountCallNameResponse) GetProtocol() int32 {
	return mx.Protocol_Account_CallName
}

func (x *AccountCallNameResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


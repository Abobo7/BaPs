// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type FriendSetIdCardResponse struct{
	message ProtoMessage
	ResponsePacket

    
}

func (x *FriendSetIdCardResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *FriendSetIdCardResponse) ProtoReflect() Message {
	return x
}

func (x *FriendSetIdCardResponse) GetProtocol() int32 {
	return mx.Protocol_Friend_SetIdCard
}

func (x *FriendSetIdCardResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


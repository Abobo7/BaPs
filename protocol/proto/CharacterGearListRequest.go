// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CharacterGearListRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *CharacterGearListRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CharacterGearListRequest) ProtoReflect() Message {
	return x
}

func (x *CharacterGearListRequest) GetProtocol() int32 {
	return mx.Protocol_CharacterGear_List
}

func (x *CharacterGearListRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeGiveGiftRequest struct{
	message ProtoMessage
	RequestPacket

    CafeDBId int64
    CharacterUniqueId int64
    ConsumeRequestDB *ConsumeRequestDB
}

func (x *CafeGiveGiftRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeGiveGiftRequest) ProtoReflect() Message {
	return x
}

func (x *CafeGiveGiftRequest) GetProtocol() int32 {
	return mx.Protocol_Cafe_GiveGift
}

func (x *CafeGiveGiftRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


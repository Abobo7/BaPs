// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EventContentBoxGachaShopListRequest struct{
	message ProtoMessage
	RequestPacket

    EventContentId int64
}

func (x *EventContentBoxGachaShopListRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EventContentBoxGachaShopListRequest) ProtoReflect() Message {
	return x
}

func (x *EventContentBoxGachaShopListRequest) GetProtocol() int32 {
	return mx.Protocol_EventContent_BoxGachaShopList
}

func (x *EventContentBoxGachaShopListRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


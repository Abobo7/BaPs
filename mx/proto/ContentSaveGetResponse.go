// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type ContentSaveGetResponse struct{
	message mx.ProtoMessage
	ResponsePacket

    HasValidData bool
    ContentSaveDB *ContentSaveDB
    EventContentChangeDB *EventContentChangeDB
}

func (x *ContentSaveGetResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ContentSaveGetResponse) ProtoReflect() mx.Message {
	return x
}

func (x *ContentSaveGetResponse) GetProtocol() int32 {
	return Protocol_ContentSave_Get
}

func (x *ContentSaveGetResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


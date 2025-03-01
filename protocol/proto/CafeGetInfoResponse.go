// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeGetInfoResponse struct{
	message ProtoMessage
	ResponsePacket

    CafeDB *CafeDB
    CafeDBs []*CafeDB
	FurnitureDBs []*FurnitureDB
}

func (x *CafeGetInfoResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeGetInfoResponse) ProtoReflect() Message {
	return x
}

func (x *CafeGetInfoResponse) GetProtocol() int32 {
	return mx.Protocol_Cafe_Get
}

func (x *CafeGetInfoResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type CafeGetInfoResponse struct{
	message mx.ProtoMessage
	ResponsePacket

    CafeDB *CafeDB
    CafeDBs []*CafeDB
    FurnitureDBs []*FurnitureDB
}

func (x *CafeGetInfoResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeGetInfoResponse) ProtoReflect() mx.Message {
	return x
}

func (x *CafeGetInfoResponse) GetProtocol() int32 {
	return Protocol_Cafe_Get
}

func (x *CafeGetInfoResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


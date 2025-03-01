// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeRemoveAllFurnitureRequest struct{
	message ProtoMessage
	RequestPacket

    CafeDBId int64
}

func (x *CafeRemoveAllFurnitureRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeRemoveAllFurnitureRequest) ProtoReflect() Message {
	return x
}

func (x *CafeRemoveAllFurnitureRequest) GetProtocol() int32 {
	return mx.Protocol_Cafe_RemoveAll
}

func (x *CafeRemoveAllFurnitureRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


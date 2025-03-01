// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeRelocateFurnitureRequest struct{
	message ProtoMessage
	RequestPacket

    CafeDBId int64
    FurnitureDB *FurnitureDB
}

func (x *CafeRelocateFurnitureRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeRelocateFurnitureRequest) ProtoReflect() Message {
	return x
}

func (x *CafeRelocateFurnitureRequest) GetProtocol() int32 {
	return mx.Protocol_Cafe_Relocate
}

func (x *CafeRelocateFurnitureRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


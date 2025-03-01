// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeDeployFurnitureRequest struct{
	message ProtoMessage
	RequestPacket

    CafeDBId int64
    FurnitureDB *FurnitureDB
}

func (x *CafeDeployFurnitureRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeDeployFurnitureRequest) ProtoReflect() Message {
	return x
}

func (x *CafeDeployFurnitureRequest) GetProtocol() int32 {
	return mx.Protocol_Cafe_Deploy
}

func (x *CafeDeployFurnitureRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


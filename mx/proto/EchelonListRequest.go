// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type EchelonListRequest struct{
	message mx.ProtoMessage
	RequestPacket

    
}

func (x *EchelonListRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EchelonListRequest) ProtoReflect() mx.Message {
	return x
}

func (x *EchelonListRequest) GetProtocol() int32 {
	return Protocol_Echelon_List
}

func (x *EchelonListRequest) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


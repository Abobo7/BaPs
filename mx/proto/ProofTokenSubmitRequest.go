// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type ProofTokenSubmitRequest struct{
	message mx.ProtoMessage
	RequestPacket

    
}

func (x *ProofTokenSubmitRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ProofTokenSubmitRequest) ProtoReflect() mx.Message {
	return x
}

func (x *ProofTokenSubmitRequest) GetProtocol() int32 {
	return Protocol_ProofToken_Submit
}

func (x *ProofTokenSubmitRequest) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


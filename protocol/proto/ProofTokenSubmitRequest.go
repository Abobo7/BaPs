// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ProofTokenSubmitRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *ProofTokenSubmitRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ProofTokenSubmitRequest) ProtoReflect() Message {
	return x
}

func (x *ProofTokenSubmitRequest) GetProtocol() int32 {
	return mx.Protocol_ProofToken_Submit
}

func (x *ProofTokenSubmitRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


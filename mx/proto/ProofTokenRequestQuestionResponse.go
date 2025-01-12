// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type ProofTokenRequestQuestionResponse struct{
	message mx.ProtoMessage
	ResponsePacket

	Hint int64
	Question string
}

func (x *ProofTokenRequestQuestionResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ProofTokenRequestQuestionResponse) ProtoReflect() mx.Message {
	return x
}

func (x *ProofTokenRequestQuestionResponse) GetProtocol() int32 {
	return Protocol_ProofToken_RequestQuestion
}

func (x *ProofTokenRequestQuestionResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type AttachmentEmblemListResponse struct{
	message mx.ProtoMessage
	ResponsePacket

    EmblemDBs []*EmblemDB
}

func (x *AttachmentEmblemListResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AttachmentEmblemListResponse) ProtoReflect() mx.Message {
	return x
}

func (x *AttachmentEmblemListResponse) GetProtocol() int32 {
	return Protocol_Attachment_EmblemList
}

func (x *AttachmentEmblemListResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


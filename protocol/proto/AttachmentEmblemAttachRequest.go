// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type AttachmentEmblemAttachRequest struct{
	message ProtoMessage
	RequestPacket

    UniqueId int64
}

func (x *AttachmentEmblemAttachRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AttachmentEmblemAttachRequest) ProtoReflect() Message {
	return x
}

func (x *AttachmentEmblemAttachRequest) GetProtocol() int32 {
	return mx.Protocol_Attachment_EmblemAttach
}

func (x *AttachmentEmblemAttachRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


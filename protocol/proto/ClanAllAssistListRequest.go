// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ClanAllAssistListRequest struct{
	message ProtoMessage
	RequestPacket

    EchelonType EchelonType
    PendingAssistUseInfo []*ClanAssistUseInfo
    IsPractice bool
}

func (x *ClanAllAssistListRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ClanAllAssistListRequest) ProtoReflect() Message {
	return x
}

func (x *ClanAllAssistListRequest) GetProtocol() int32 {
	return mx.Protocol_Clan_AllAssistList
}

func (x *ClanAllAssistListRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


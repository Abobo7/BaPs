// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ClanJoinRequest struct{
	message ProtoMessage
	RequestPacket

    ClanDBId int64
}

func (x *ClanJoinRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ClanJoinRequest) ProtoReflect() Message {
	return x
}

func (x *ClanJoinRequest) GetProtocol() int32 {
	return mx.Protocol_Clan_Join
}

func (x *ClanJoinRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


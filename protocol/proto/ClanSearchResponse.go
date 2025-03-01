// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ClanSearchResponse struct{
	message ProtoMessage
	ResponsePacket

    ClanDBs []*ClanDB
}

func (x *ClanSearchResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ClanSearchResponse) ProtoReflect() Message {
	return x
}

func (x *ClanSearchResponse) GetProtocol() int32 {
	return mx.Protocol_Clan_Search
}

func (x *ClanSearchResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ClanPermitRequest struct{
	message ProtoMessage
	RequestPacket

    ApplicantAccountId int64
    IsPerMit bool
}

func (x *ClanPermitRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ClanPermitRequest) ProtoReflect() Message {
	return x
}

func (x *ClanPermitRequest) GetProtocol() int32 {
	return mx.Protocol_Clan_Permit
}

func (x *ClanPermitRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


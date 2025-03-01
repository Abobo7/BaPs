// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type RaidGetBestTeamResponse struct{
	message ProtoMessage
	ResponsePacket

    RaidTeamSettingDBs []*RaidTeamSettingDB
}

func (x *RaidGetBestTeamResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *RaidGetBestTeamResponse) ProtoReflect() Message {
	return x
}

func (x *RaidGetBestTeamResponse) GetProtocol() int32 {
	return mx.Protocol_Raid_GetBestTeam
}

func (x *RaidGetBestTeamResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


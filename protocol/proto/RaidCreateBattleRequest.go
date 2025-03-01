// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type RaidCreateBattleRequest struct{
	message ProtoMessage
	RequestPacket

    RaidUniqueId int64
    IsPractice bool
    Tags []int32
    IsPublic bool
    Difficulty Difficulty
    AssistUseInfo *ClanAssistUseInfo
}

func (x *RaidCreateBattleRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *RaidCreateBattleRequest) ProtoReflect() Message {
	return x
}

func (x *RaidCreateBattleRequest) GetProtocol() int32 {
	return mx.Protocol_Raid_CreateBattle
}

func (x *RaidCreateBattleRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


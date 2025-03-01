// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type WeekDungeonEnterBattleResponse struct{
	message ProtoMessage
	ResponsePacket

    ParcelResultDB *ParcelResultDB
    Seed int32
    Sequence int32
}

func (x *WeekDungeonEnterBattleResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *WeekDungeonEnterBattleResponse) ProtoReflect() Message {
	return x
}

func (x *WeekDungeonEnterBattleResponse) GetProtocol() int32 {
	return mx.Protocol_WeekDungeon_EnterBattle
}

func (x *WeekDungeonEnterBattleResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


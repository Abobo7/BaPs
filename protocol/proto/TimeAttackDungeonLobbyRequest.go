// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type TimeAttackDungeonLobbyRequest struct{
	message ProtoMessage
	RequestPacket

    
}

func (x *TimeAttackDungeonLobbyRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *TimeAttackDungeonLobbyRequest) ProtoReflect() Message {
	return x
}

func (x *TimeAttackDungeonLobbyRequest) GetProtocol() int32 {
	return mx.Protocol_TimeAttackDungeon_Lobby
}

func (x *TimeAttackDungeonLobbyRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


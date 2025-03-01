// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ScenarioLobbyStudentChangeRequest struct{
	message ProtoMessage
	RequestPacket

    LobbyStudents []int64
    LobbyStudentsBefore []int64
}

func (x *ScenarioLobbyStudentChangeRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ScenarioLobbyStudentChangeRequest) ProtoReflect() Message {
	return x
}

func (x *ScenarioLobbyStudentChangeRequest) GetProtocol() int32 {
	return mx.Protocol_Scenario_LobbyStudentChange
}

func (x *ScenarioLobbyStudentChangeRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ScenarioAccountStudentChangeResponse struct{
	message ProtoMessage
	ResponsePacket

    
}

func (x *ScenarioAccountStudentChangeResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ScenarioAccountStudentChangeResponse) ProtoReflect() Message {
	return x
}

func (x *ScenarioAccountStudentChangeResponse) GetProtocol() int32 {
	return mx.Protocol_Scenario_AccountStudentChange
}

func (x *ScenarioAccountStudentChangeResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


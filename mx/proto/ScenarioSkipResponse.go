// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type ScenarioSkipResponse struct{
	message mx.ProtoMessage
	ResponsePacket

    
}

func (x *ScenarioSkipResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ScenarioSkipResponse) ProtoReflect() mx.Message {
	return x
}

func (x *ScenarioSkipResponse) GetProtocol() int32 {
	return Protocol_Scenario_Skip
}

func (x *ScenarioSkipResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


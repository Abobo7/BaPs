// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type OpenConditionEventListRequest struct{
	message mx.ProtoMessage
	RequestPacket

    ConquestEventIds []int64
    WorldRaidSeasonAndGroupIds map[int64]int64
}

func (x *OpenConditionEventListRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *OpenConditionEventListRequest) ProtoReflect() mx.Message {
	return x
}

func (x *OpenConditionEventListRequest) GetProtocol() int32 {
	return Protocol_OpenCondition_EventList
}

func (x *OpenConditionEventListRequest) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


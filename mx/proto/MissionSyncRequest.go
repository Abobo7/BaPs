// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type MissionSyncRequest struct{
	message mx.ProtoMessage
	RequestPacket

    
}

func (x *MissionSyncRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *MissionSyncRequest) ProtoReflect() mx.Message {
	return x
}

func (x *MissionSyncRequest) GetProtocol() int32 {
	return Protocol_Mission_Sync
}

func (x *MissionSyncRequest) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


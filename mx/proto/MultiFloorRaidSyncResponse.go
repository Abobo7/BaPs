// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type MultiFloorRaidSyncResponse struct{
	message mx.ProtoMessage
	ResponsePacket

    MultiFloorRaidDBs []*MultiFloorRaidDB
}

func (x *MultiFloorRaidSyncResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *MultiFloorRaidSyncResponse) ProtoReflect() mx.Message {
	return x
}

func (x *MultiFloorRaidSyncResponse) GetProtocol() int32 {
	return Protocol_MultiFloorRaid_Sync
}

func (x *MultiFloorRaidSyncResponse) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/mx"
)

type AccountLoginSyncRequest struct{
	message mx.ProtoMessage
	RequestPacket

    SyncProtocols []Protocol
	SkillCutInOption string
}

func (x *AccountLoginSyncRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *AccountLoginSyncRequest) ProtoReflect() mx.Message {
	return x
}

func (x *AccountLoginSyncRequest) GetProtocol() int32 {
	return Protocol_Account_LoginSync
}

func (x *AccountLoginSyncRequest) SetSessionKey(base *mx.BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


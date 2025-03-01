// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CharacterPotentialGrowthResponse struct{
	message ProtoMessage
	ResponsePacket

    CharacterDB *CharacterDB
    ParcelResultDB *ParcelResultDB
}

func (x *CharacterPotentialGrowthResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CharacterPotentialGrowthResponse) ProtoReflect() Message {
	return x
}

func (x *CharacterPotentialGrowthResponse) GetProtocol() int32 {
	return mx.Protocol_Character_PotentialGrowth
}

func (x *CharacterPotentialGrowthResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


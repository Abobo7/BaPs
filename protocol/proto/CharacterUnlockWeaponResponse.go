// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CharacterUnlockWeaponResponse struct{
	message ProtoMessage
	ResponsePacket

    WeaponDB *WeaponDB
}

func (x *CharacterUnlockWeaponResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CharacterUnlockWeaponResponse) ProtoReflect() Message {
	return x
}

func (x *CharacterUnlockWeaponResponse) GetProtocol() int32 {
	return mx.Protocol_Character_UnlockWeapon
}

func (x *CharacterUnlockWeaponResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


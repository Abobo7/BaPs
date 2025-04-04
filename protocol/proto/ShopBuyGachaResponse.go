// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ShopBuyGachaResponse struct{
	message ProtoMessage
	ResponsePacket

    AccountCurrencyDB *AccountCurrencyDB
    ConsumeResultDB *ConsumeResultDB
    ParcelResultDB *ParcelResultDB
}

func (x *ShopBuyGachaResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ShopBuyGachaResponse) ProtoReflect() Message {
	return x
}

func (x *ShopBuyGachaResponse) GetProtocol() int32 {
	return mx.Protocol_Shop_BuyGacha
}

func (x *ShopBuyGachaResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


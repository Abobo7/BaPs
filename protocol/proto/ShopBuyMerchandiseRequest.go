// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ShopBuyMerchandiseRequest struct{
	message ProtoMessage
	RequestPacket

    IsRefreshGoods bool
    ShopUniqueId int64
    GoodsId int64
    PurchaseCount int64
}

func (x *ShopBuyMerchandiseRequest) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ShopBuyMerchandiseRequest) ProtoReflect() Message {
	return x
}

func (x *ShopBuyMerchandiseRequest) GetProtocol() int32 {
	return mx.Protocol_Shop_BuyMerchandise
}

func (x *ShopBuyMerchandiseRequest) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


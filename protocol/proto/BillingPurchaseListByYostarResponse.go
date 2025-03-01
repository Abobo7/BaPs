// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type BillingPurchaseListByYostarResponse struct{
	message ProtoMessage
	ResponsePacket

    CountList []*PurchaseCountDB
    OrderList []*PurchaseOrderDB
    MonthlyProductList []*MonthlyProductPurchaseDB
    BlockedProductDBs []*BlockedProductDB
}

func (x *BillingPurchaseListByYostarResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *BillingPurchaseListByYostarResponse) ProtoReflect() Message {
	return x
}

func (x *BillingPurchaseListByYostarResponse) GetProtocol() int32 {
	return mx.Protocol_Billing_PurchaseListByYostar
}

func (x *BillingPurchaseListByYostarResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


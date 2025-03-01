// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ContentSweepResponse struct{
	message ProtoMessage
	ResponsePacket

    ClearParcels [][]*ParcelInfo
    BonusParcels []*ParcelInfo
    EventContentBonusParcels [][]*ParcelInfo
    ParcelResult *ParcelResultDB
    CampaignStageHistoryDB *CampaignStageHistoryDB
}

func (x *ContentSweepResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ContentSweepResponse) ProtoReflect() Message {
	return x
}

func (x *ContentSweepResponse) GetProtocol() int32 {
	return mx.Protocol_ContentSweep_Request
}

func (x *ContentSweepResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


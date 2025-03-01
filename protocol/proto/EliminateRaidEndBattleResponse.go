// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EliminateRaidEndBattleResponse struct{
	message ProtoMessage
	ResponsePacket

    RankingPoint int64
    BestRankingPoint int64
    ClearTimePoint int64
    HPPercentScorePoint int64
    DefaultClearPoint int64
    ParcelResultDB *ParcelResultDB
}

func (x *EliminateRaidEndBattleResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EliminateRaidEndBattleResponse) ProtoReflect() Message {
	return x
}

func (x *EliminateRaidEndBattleResponse) GetProtocol() int32 {
	return mx.Protocol_EliminateRaid_EndBattle
}

func (x *EliminateRaidEndBattleResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


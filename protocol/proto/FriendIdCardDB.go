// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "github.com/gucooing/BaPs/pkg/mx"
)

type FriendIdCardDB struct{
    Level int32
    FriendCode string
    Comment string
    LastConnectTime mx.MxTime
    RepresentCharacterUniqueId int64
    RepresentCharacterCostumeId int64
    SearchPermission bool
    AutoAcceptFriendRequest bool
    CardBackgroundId int64
    ShowAccountLevel bool
    ShowFriendCode bool
    ShowRaidRanking bool
    ShowArenaRanking bool
    ShowEliminateRaidRanking bool
    ArenaRanking int64
    RaidRanking int64
    RaidTier int32
    EliminateRaidRanking int64
    EliminateRaidTier int32
    EmblemId int64
}


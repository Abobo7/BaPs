// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "time"

    "github.com/gucooing/BaPs/pkg/mx"
)

type TimeAttackDungeonRoomDB struct{
    AccountId int64
    SeasonId int64
    RoomId int64
    CreateDate mx.MxTime
    RewardDate mx.MxTime
    IsPractice bool
    SweepHistoryDates []time.Time
    BattleHistoryDBs []*TimeAttackDungeonBattleHistoryDB
}


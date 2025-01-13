// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "time"
)

type ContentSaveDB struct{
    ContentType ContentType
    AccountServerId int64
    CreateTime time.Time
    StageUniqueId int64
    LastEnterStageEchelonNumber int64
    StageEntranceFee []*ParcelInfo
    EnemyKillCountByUniqueId map[int64]int64
    TacticClearTimeMscSum int64
    AccountLevelWhenCreateDB int64
    BIEchelon string
    BIEchelon1 string
    BIEchelon2 string
    BIEchelon3 string
    BIEchelon4 string
}


// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "time"
)

type AttendanceHistoryDB struct{
    ServerId int64
    AccountServerId int64
    AttendanceBookUniqueId int64
    AttendedDay map[int64]*time.Time
    Expired bool
}


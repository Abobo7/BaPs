// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type NotificationEventContentReddotResponse struct{
	message ProtoMessage
	ResponsePacket

    Reddots map[int64][]NotificationEventReddot
    EventContentUnlockCGDBs map[int64][]*EventContentCollectionDB
}

func (x *NotificationEventContentReddotResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *NotificationEventContentReddotResponse) ProtoReflect() Message {
	return x
}

func (x *NotificationEventContentReddotResponse) GetProtocol() int32 {
	return mx.Protocol_Notification_EventContentReddotCheck
}

func (x *NotificationEventContentReddotResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}


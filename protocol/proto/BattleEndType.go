// Code generated by gucooing DO NOT EDIT.

package proto

type BattleEndType int32

const (
    BattleEndType_None = 0
    BattleEndType_AllNearlyDead = 1
    BattleEndType_TimeOut = 2
    BattleEndType_EscortFailed = 3
    BattleEndType_Clear = 4
)

var (
    BattleEndType_name = map[int32]string{
      0 : "None",
      1 : "AllNearlyDead",
      2 : "TimeOut",
      3 : "EscortFailed",
      4 : "Clear",
   }
    BattleEndType_value = map[string]int32{
      "None" : 0,
      "AllNearlyDead" : 1,
      "TimeOut" : 2,
      "EscortFailed" : 3,
      "Clear" : 4,
   }
)

func (x BattleEndType) String() string {
  return BattleEndType_name[int32(x)]
}


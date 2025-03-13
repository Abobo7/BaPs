// Code generated by gucooing DO NOT EDIT.

package proto

type TimeAttackDungeonType int32

const (
    TimeAttackDungeonType_None = 0
    TimeAttackDungeonType_Defense = 1
    TimeAttackDungeonType_Shooting = 2
    TimeAttackDungeonType_Destruction = 3
    TimeAttackDungeonType_Escort = 4
)

var (
    TimeAttackDungeonType_name = map[int32]string{
      0 : "None",
      1 : "Defense",
      2 : "Shooting",
      3 : "Destruction",
      4 : "Escort",
   }
    TimeAttackDungeonType_value = map[string]int32{
      "None" : 0,
      "Defense" : 1,
      "Shooting" : 2,
      "Destruction" : 3,
      "Escort" : 4,
   }
)

func (x TimeAttackDungeonType) String() string {
  return TimeAttackDungeonType_name[int32(x)]
}

func (x TimeAttackDungeonType) Value() int32 {
    return int32(x)
}

func GetTimeAttackDungeonType(s string) TimeAttackDungeonType {
    return TimeAttackDungeonType(TimeAttackDungeonType_value[s])
}

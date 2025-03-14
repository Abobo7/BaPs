// Code generated by gucooing DO NOT EDIT.

package proto

type EchelonType int32

const (
    EchelonType_None EchelonType = 0
    EchelonType_Adventure EchelonType = 1
    EchelonType_Raid EchelonType = 2
    EchelonType_ArenaAttack EchelonType = 3
    EchelonType_ArenaDefence EchelonType = 4
    EchelonType_WeekDungeonChaserA EchelonType = 5
    EchelonType_Scenario EchelonType = 6
    EchelonType_WeekDungeonBlood EchelonType = 7
    EchelonType_WeekDungeonChaserB EchelonType = 8
    EchelonType_WeekDungeonChaserC EchelonType = 9
    EchelonType_WeekDungeonFindGift EchelonType = 10
    EchelonType_EventContent EchelonType = 11
    EchelonType_SchoolDungeonA EchelonType = 12
    EchelonType_SchoolDungeonB EchelonType = 13
    EchelonType_SchoolDungeonC EchelonType = 14
    EchelonType_TimeAttack EchelonType = 15
    EchelonType_WorldRaid EchelonType = 16
    EchelonType_Conquest EchelonType = 17
    EchelonType_ConquestManage EchelonType = 18
    EchelonType_StoryStrategyStage EchelonType = 19
    EchelonType_EliminateRaid01 EchelonType = 20
    EchelonType_EliminateRaid02 EchelonType = 21
    EchelonType_EliminateRaid03 EchelonType = 22
    EchelonType_Field EchelonType = 23
    EchelonType_MultiFloorRaid EchelonType = 24
    EchelonType_Temp EchelonType = 25
)

var (
    EchelonType_name = map[int32]string{
      0 : "None",
      1 : "Adventure",
      2 : "Raid",
      3 : "ArenaAttack",
      4 : "ArenaDefence",
      5 : "WeekDungeonChaserA",
      6 : "Scenario",
      7 : "WeekDungeonBlood",
      8 : "WeekDungeonChaserB",
      9 : "WeekDungeonChaserC",
      10 : "WeekDungeonFindGift",
      11 : "EventContent",
      12 : "SchoolDungeonA",
      13 : "SchoolDungeonB",
      14 : "SchoolDungeonC",
      15 : "TimeAttack",
      16 : "WorldRaid",
      17 : "Conquest",
      18 : "ConquestManage",
      19 : "StoryStrategyStage",
      20 : "EliminateRaid01",
      21 : "EliminateRaid02",
      22 : "EliminateRaid03",
      23 : "Field",
      24 : "MultiFloorRaid",
      25 : "Temp",
   }
    EchelonType_value = map[string]int32{
      "None" : 0,
      "Adventure" : 1,
      "Raid" : 2,
      "ArenaAttack" : 3,
      "ArenaDefence" : 4,
      "WeekDungeonChaserA" : 5,
      "Scenario" : 6,
      "WeekDungeonBlood" : 7,
      "WeekDungeonChaserB" : 8,
      "WeekDungeonChaserC" : 9,
      "WeekDungeonFindGift" : 10,
      "EventContent" : 11,
      "SchoolDungeonA" : 12,
      "SchoolDungeonB" : 13,
      "SchoolDungeonC" : 14,
      "TimeAttack" : 15,
      "WorldRaid" : 16,
      "Conquest" : 17,
      "ConquestManage" : 18,
      "StoryStrategyStage" : 19,
      "EliminateRaid01" : 20,
      "EliminateRaid02" : 21,
      "EliminateRaid03" : 22,
      "Field" : 23,
      "MultiFloorRaid" : 24,
      "Temp" : 25,
   }
)

func (x EchelonType) String() string {
  return EchelonType_name[int32(x)]
}

func (x EchelonType) Value() int32 {
    return int32(x)
}
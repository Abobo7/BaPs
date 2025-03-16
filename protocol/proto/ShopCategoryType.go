// Code generated by gucooing DO NOT EDIT.

package proto

type ShopCategoryType int32

const (
    ShopCategoryType_General = 0
    ShopCategoryType_SecretStone = 1
    ShopCategoryType_Raid = 2
    ShopCategoryType_Gold = 3
    ShopCategoryType_Ap = 4
    ShopCategoryType_PickupGacha = 5
    ShopCategoryType_NormalGacha = 6
    ShopCategoryType_PointGacha = 7
    ShopCategoryType_EventGacha = 8
    ShopCategoryType_ArenaTicket = 9
    ShopCategoryType_Arena = 10
    ShopCategoryType_TutoGacha = 11
    ShopCategoryType_RecruitSellection = 12
    ShopCategoryType_EventContent_0 = 13
    ShopCategoryType_EventContent_1 = 14
    ShopCategoryType_EventContent_2 = 15
    ShopCategoryType_EventContent_3 = 16
    ShopCategoryType_EventContent_4 = 17
    ShopCategoryType__Obsolete = 18
    ShopCategoryType_LimitedGacha = 19
    ShopCategoryType_MasterCoin = 20
    ShopCategoryType_SecretStoneGrowth = 21
    ShopCategoryType_TicketGacha = 22
    ShopCategoryType_DirectPayGacha_DontUseGlobal = 23
    ShopCategoryType_FesGacha = 24
    ShopCategoryType_TimeAttack = 25
    ShopCategoryType_Chaser = 26
    ShopCategoryType_ChaserTicket = 27
    ShopCategoryType_SchoolDungeonTicket = 28
    ShopCategoryType_AcademyTicket = 29
    ShopCategoryType_Special = 30
    ShopCategoryType_Care = 31
    ShopCategoryType_BeforehandGacha = 32
    ShopCategoryType_EliminateRaid = 33
    ShopCategoryType_GlobalSpecialGacha = 34
)

var (
    ShopCategoryType_name = map[int32]string{
      0 : "General",
      1 : "SecretStone",
      2 : "Raid",
      3 : "Gold",
      4 : "Ap",
      5 : "PickupGacha",
      6 : "NormalGacha",
      7 : "PointGacha",
      8 : "EventGacha",
      9 : "ArenaTicket",
      10 : "Arena",
      11 : "TutoGacha",
      12 : "RecruitSellection",
      13 : "EventContent_0",
      14 : "EventContent_1",
      15 : "EventContent_2",
      16 : "EventContent_3",
      17 : "EventContent_4",
      18 : "_Obsolete",
      19 : "LimitedGacha",
      20 : "MasterCoin",
      21 : "SecretStoneGrowth",
      22 : "TicketGacha",
      23 : "DirectPayGacha_DontUseGlobal",
      24 : "FesGacha",
      25 : "TimeAttack",
      26 : "Chaser",
      27 : "ChaserTicket",
      28 : "SchoolDungeonTicket",
      29 : "AcademyTicket",
      30 : "Special",
      31 : "Care",
      32 : "BeforehandGacha",
      33 : "EliminateRaid",
      34 : "GlobalSpecialGacha",
   }
    ShopCategoryType_value = map[string]int32{
      "General" : 0,
      "SecretStone" : 1,
      "Raid" : 2,
      "Gold" : 3,
      "Ap" : 4,
      "PickupGacha" : 5,
      "NormalGacha" : 6,
      "PointGacha" : 7,
      "EventGacha" : 8,
      "ArenaTicket" : 9,
      "Arena" : 10,
      "TutoGacha" : 11,
      "RecruitSellection" : 12,
      "EventContent_0" : 13,
      "EventContent_1" : 14,
      "EventContent_2" : 15,
      "EventContent_3" : 16,
      "EventContent_4" : 17,
      "_Obsolete" : 18,
      "LimitedGacha" : 19,
      "MasterCoin" : 20,
      "SecretStoneGrowth" : 21,
      "TicketGacha" : 22,
      "DirectPayGacha_DontUseGlobal" : 23,
      "FesGacha" : 24,
      "TimeAttack" : 25,
      "Chaser" : 26,
      "ChaserTicket" : 27,
      "SchoolDungeonTicket" : 28,
      "AcademyTicket" : 29,
      "Special" : 30,
      "Care" : 31,
      "BeforehandGacha" : 32,
      "EliminateRaid" : 33,
      "GlobalSpecialGacha" : 34,
   }
)

func (x ShopCategoryType) String() string {
  return ShopCategoryType_name[int32(x)]
}

func GetShopCategoryType(s string) ShopCategoryType  {
    return ShopCategoryType(ShopCategoryType_value[s])
}
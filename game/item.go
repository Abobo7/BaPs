package game

import (
	"time"

	"github.com/gucooing/BaPs/common/enter"
	sro "github.com/gucooing/BaPs/common/server_only"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func NewItemList(s *enter.Session) map[int64]*sro.ItemInfo {
	list := make(map[int64]*sro.ItemInfo)
	list[2] = &sro.ItemInfo{
		ServerId:   GetServerId(s),
		UniqueId:   2,
		StackCount: 5,
	}
	return list
}

func GetItemBin(s *enter.Session) *sro.ItemBin {
	bin := GetPlayerBin(s)
	if bin == nil {
		return nil
	}
	if bin.ItemBin == nil {
		bin.ItemBin = &sro.ItemBin{}
	}
	return bin.ItemBin
}

func GetItemList(s *enter.Session) map[int64]*sro.ItemInfo {
	bin := GetItemBin(s)
	if bin == nil {
		return nil
	}
	if bin.ItemInfoList == nil {
		bin.ItemInfoList = NewItemList(s)
	}
	return bin.ItemInfoList
}

func GetItemInfo(s *enter.Session, itemId int64) *sro.ItemInfo {
	bin := GetItemList(s)
	if bin == nil {
		return nil
	}
	return bin[itemId]
}

func AddItem(s *enter.Session, id int64, num int32) int64 {
	bin := GetItemBin(s)
	if bin == nil {
		return 0
	}
	if bin.ItemInfoList == nil {
		bin.ItemInfoList = NewItemList(s)
	}
	if info, ok := bin.ItemInfoList[id]; ok {
		info.StackCount += num
		return info.ServerId
	}
	info := &sro.ItemInfo{
		ServerId:   GetServerId(s),
		UniqueId:   id,
		StackCount: num,
	}
	bin.ItemInfoList[id] = info
	return info.ServerId
}

func RemoveItem(s *enter.Session, id int64, num int32) bool {
	bin := GetItemBin(s)
	if bin == nil {
		return false
	}
	if bin.ItemInfoList == nil {
		bin.ItemInfoList = NewItemList(s)
	}
	if info, ok := bin.ItemInfoList[id]; ok {
		if info.StackCount >= num {
			info.StackCount -= num
			return true
		}
	}
	return false
}

var DefaultCurrencyNum = map[int32]int64{
	proto.CurrencyTypes_Gem:                      600,
	proto.CurrencyTypes_GemPaid:                  0,
	proto.CurrencyTypes_GemBonus:                 600,
	proto.CurrencyTypes_Gold:                     10000,
	proto.CurrencyTypes_ActionPoint:              24,
	proto.CurrencyTypes_AcademyTicket:            3,
	proto.CurrencyTypes_ArenaTicket:              5,
	proto.CurrencyTypes_RaidTicket:               3,
	proto.CurrencyTypes_WeekDungeonChaserATicket: 0,
	proto.CurrencyTypes_WeekDungeonChaserBTicket: 0,
	proto.CurrencyTypes_WeekDungeonChaserCTicket: 0,
	proto.CurrencyTypes_SchoolDungeonATicket:     0,
	proto.CurrencyTypes_SchoolDungeonBTicket:     0,
	proto.CurrencyTypes_SchoolDungeonCTicket:     0,
	proto.CurrencyTypes_TimeAttackDungeonTicket:  3,
	proto.CurrencyTypes_MasterCoin:               0,
	proto.CurrencyTypes_WorldRaidTicketA:         40,
	proto.CurrencyTypes_WorldRaidTicketB:         40,
	proto.CurrencyTypes_WorldRaidTicketC:         40,
	proto.CurrencyTypes_ChaserTotalTicket:        6,
	proto.CurrencyTypes_SchoolDungeonTotalTicket: 6,
	proto.CurrencyTypes_EliminateTicketA:         3,
	proto.CurrencyTypes_EliminateTicketB:         3,
	proto.CurrencyTypes_EliminateTicketC:         3,
	proto.CurrencyTypes_EliminateTicketD:         3,
}

func NewCurrencyInfo() map[int32]*sro.CurrencyInfo {
	list := make(map[int32]*sro.CurrencyInfo)
	for k, v := range DefaultCurrencyNum {
		list[k] = &sro.CurrencyInfo{
			CurrencyId:  k,
			CurrencyNum: v,
			UpdateTime:  time.Now().Unix(),
		}
	}
	return list
}

func UpCurrency(s *enter.Session, parcelId int64, num int64) *sro.CurrencyInfo {
	bin := GetCurrencyList(s)
	if bin == nil {
		return nil
	}
	info := GetCurrencyInfo(s, int32(parcelId))
	if info == nil {
		return nil
	}
	if num < 0 && info.CurrencyNum < -(num) {
		return nil
	}
	info.CurrencyNum += num
	info.UpdateTime = time.Now().Unix()

	gemBonus := GetCurrencyInfo(s, proto.CurrencyTypes_GemBonus)
	gem := GetCurrencyInfo(s, proto.CurrencyTypes_Gem)
	if gem != nil || gemBonus != nil {
		gem.CurrencyNum = gemBonus.CurrencyNum
	}

	return info
}

func GetCurrencyList(s *enter.Session) map[int32]*sro.CurrencyInfo {
	bin := GetItemBin(s)
	if bin == nil {
		return nil
	}
	if bin.CurrencyInfoList == nil {
		bin.CurrencyInfoList = NewCurrencyInfo()
	}
	return bin.CurrencyInfoList
}

func GetCurrencyInfo(s *enter.Session, currencyId int32) *sro.CurrencyInfo {
	bin := GetItemBin(s)
	if bin == nil {
		return nil
	}
	if bin.CurrencyInfoList == nil {
		bin.CurrencyInfoList = NewCurrencyInfo()
	}
	if bin.CurrencyInfoList[currencyId] == nil {
		bin.CurrencyInfoList[currencyId] = &sro.CurrencyInfo{
			CurrencyId:  currencyId,
			CurrencyNum: DefaultCurrencyNum[currencyId],
			UpdateTime:  time.Now().Unix(),
		}
	}
	return bin.CurrencyInfoList[currencyId]
}

func GetAccountCurrencyDB(s *enter.Session) *proto.AccountCurrencyDB {
	accountCurrencyDB := &proto.AccountCurrencyDB{
		AccountLevel:           int64(GetAccountLevel(s)),
		AcademyLocationRankSum: 1,
		CurrencyDict:           make(map[proto.CurrencyTypes]int64),
		UpdateTimeDict:         make(map[proto.CurrencyTypes]mx.MxTime),
	}
	for id, db := range GetCurrencyList(s) {
		accountCurrencyDB.CurrencyDict[proto.CurrencyTypes(proto.CurrencyTypes_name[id])] = db.CurrencyNum
		accountCurrencyDB.UpdateTimeDict[proto.CurrencyTypes(proto.CurrencyTypes_name[id])] = mx.Unix(db.UpdateTime, 0)
	}

	return accountCurrencyDB
}

type ParcelResult struct {
	ParcelType proto.ParcelType
	ParcelId   int64
	Amount     int64
}

func GetParcelResultList(typeList []string, idList, numList []int64) []*ParcelResult {
	list := make([]*ParcelResult, 0)
	if len(typeList) == len(idList) &&
		len(idList) == len(numList) {
		for index, rewardType := range typeList {
			list = append(list, &ParcelResult{
				ParcelType: proto.ParcelType(proto.ParcelType_value[rewardType]),
				ParcelId:   idList[index],
				Amount:     numList[index],
			})
		}
	}
	return list
}

func ParcelResultDB(s *enter.Session, parcelResultList []*ParcelResult) *proto.ParcelResultDB {
	info := &proto.ParcelResultDB{
		AccountDB:                       GetAccountDB(s),
		AcademyLocationDBs:              nil,
		CharacterDBs:                    nil,
		WeaponDBs:                       nil,
		CostumeDBs:                      nil,
		TSSCharacterDBs:                 nil,
		EquipmentDBs:                    nil,
		RemovedEquipmentIds:             nil,
		ItemDBs:                         nil,
		RemovedItemIds:                  nil,
		FurnitureDBs:                    nil,
		RemovedFurnitureIds:             nil,
		IdCardBackgroundDBs:             nil,
		StickerDBs:                      nil,
		CharacterNewUniqueIds:           nil,
		SecretStoneCharacterIdAndCounts: nil,
		DisplaySequence:                 make([]*proto.ParcelInfo, 0),
		ParcelForMission:                nil,
		ParcelResultStepInfoList:        nil,
		BaseAccountExp:                  0,
		AdditionalAccountExp:            0,
		GachaResultCharacters:           nil,
	}

	for _, parcelResult := range parcelResultList {
		switch parcelResult.ParcelType {
		case proto.ParcelType_Currency: // 货币
			UpCurrency(s, parcelResult.ParcelId, parcelResult.Amount)
			info.AccountCurrencyDB = GetAccountCurrencyDB(s)
		case proto.ParcelType_MemoryLobby: // 记忆大厅
			UpMemoryLobbyInfo(s, parcelResult.ParcelId)
			info.MemoryLobbyDBs = append(info.MemoryLobbyDBs,
				GetMemoryLobbyDB(s, parcelResult.ParcelId))
		case proto.ParcelType_Emblem: // 称号
			UpEmblemInfoList(s, []int64{parcelResult.ParcelId})
			info.EmblemDBs = append(info.EmblemDBs,
				GetEmblemDB(s, parcelResult.ParcelId))
		default:
			logger.Warn("没有处理的奖励类型 Unknown ParcelType:%s", parcelResult.ParcelType.String())
		}
		info.DisplaySequence = append(info.DisplaySequence, &proto.ParcelInfo{
			Key: &proto.ParcelKeyPair{
				Type: parcelResult.ParcelType,
				Id:   parcelResult.ParcelId,
			},
			Amount: parcelResult.Amount,
			Multiplier: &proto.BasisPoint{
				RawValue: 10000,
			},
			Probability: &proto.BasisPoint{
				RawValue: 10000,
			},
		})
	}

	return info
}

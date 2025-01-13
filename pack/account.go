package pack

import (
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/mx"
	"github.com/gucooing/BaPs/mx/cmd"
	"github.com/gucooing/BaPs/mx/proto"
	"github.com/gucooing/BaPs/pkg/logger"
)

func AccountAuth(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AccountAuthRequest)
	rsp := response.(*proto.AccountAuthResponse)

	rsp.CurrentVersion = req.Version
	rsp.AccountDB = game.GetAccountDB(s)
	rsp.AttendanceBookRewards = game.GetAttendanceBookRewards(s)

	rsp.IssueAlertInfos = make([]*proto.IssueAlertInfoDB, 0)
	rsp.StaticOpenConditions = game.GetStaticOpenConditions(s)
	game.SetLastConnectTime(s)
	s.AccountState = proto.AccountState_Normal
}

func AccountNickname(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AccountNicknameRequest)
	rsp := response.(*proto.AccountNicknameResponse)

	game.SetAccountNickname(s, req.Nickname)
	rsp.AccountDB = game.GetAccountDB(s)
}

func ProofTokenRequestQuestion(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.ProofTokenRequestQuestionResponse)

	rsp.Question = "1"
	rsp.Hint = 1
}

func NetworkTimeSync(s *enter.Session, request, response mx.Message) {}

func AccountLoginSync(s *enter.Session, request, response mx.Message) {
	req := request.(*proto.AccountLoginSyncRequest)
	rsp := response.(*proto.AccountLoginSyncResponse)

	rsp.FriendCode = "BNAGBIES"
	rsp.StaticOpenConditions = game.GetStaticOpenConditions(s)

	for _, cmdId := range req.SyncProtocols {
		syncReq := cmd.Get().GetRequestPacketByCmdId(int32(cmdId))
		if syncReq == nil {
			logger.Error("AccountLoginSync SyncProtocol Req failed:%v", cmdId)
			continue
		}
		syncRsp := cmd.Get().GetResponsePacketByCmdId(int32(cmdId))
		if syncRsp == nil {
			logger.Error("AccountLoginSync SyncProtocol Rsp failed:%v", cmdId)
			continue
		}
		syncRsp.SetSessionKey(rsp.BasePacket)
		switch cmdId {
		case proto.Protocol_Cafe_Get:
			CafeGetInfo(s, syncReq, syncRsp)
			rsp.CafeGetInfoResponse = syncRsp.(*proto.CafeGetInfoResponse)
		case proto.Protocol_Account_CurrencySync:
			AccountCurrencySync(s, syncReq, syncRsp)
			rsp.AccountCurrencySyncResponse = syncRsp.(*proto.AccountCurrencySyncResponse)
		case proto.Protocol_Character_List:
			CharacterList(s, syncReq, syncRsp)
			rsp.CharacterListResponse = syncRsp.(*proto.CharacterListResponse)
		case proto.Protocol_Equipment_List:
			EquipmentList(s, syncReq, syncRsp)
			rsp.EquipmentItemListResponse = syncRsp.(*proto.EquipmentItemListResponse)
		case proto.Protocol_CharacterGear_List:
			CharacterGearList(s, syncReq, syncRsp)
			rsp.CharacterGearListResponse = syncRsp.(*proto.CharacterGearListResponse)
		case proto.Protocol_Echelon_List:
			EchelonList(s, syncReq, syncRsp)
			rsp.EchelonListResponse = syncRsp.(*proto.EchelonListResponse)
		case proto.Protocol_MemoryLobby_List:
			MemoryLobbyList(s, syncReq, syncRsp)
			rsp.MemoryLobbyListResponse = syncRsp.(*proto.MemoryLobbyListResponse)
		case proto.Protocol_Campaign_List:
			CampaignList(s, syncReq, syncRsp)
			rsp.CampaignListResponse = syncRsp.(*proto.CampaignListResponse)
		case proto.Protocol_Arena_Login:
			ArenaLogin(s, syncReq, syncRsp)
			rsp.ArenaLoginResponse = syncRsp.(*proto.ArenaLoginResponse)
		case proto.Protocol_Raid_Login:
			RaidLogin(s, syncReq, syncRsp)
			rsp.RaidLoginResponse = syncRsp.(*proto.RaidLoginResponse)
		case proto.Protocol_EliminateRaid_Login:
			EliminateRaidLogin(s, syncReq, syncRsp)
			rsp.EliminateRaidLoginResponse = syncRsp.(*proto.EliminateRaidLoginResponse)
		case proto.Protocol_Craft_List:
			CraftInfoList(s, syncReq, syncRsp)
			rsp.CraftInfoListResponse = syncRsp.(*proto.CraftInfoListResponse)
		case proto.Protocol_Clan_Login:
			ClanLogin(s, syncReq, syncRsp)
			rsp.ClanLoginResponse = syncRsp.(*proto.ClanLoginResponse)
		case proto.Protocol_MomoTalk_OutLine:
			MomoTalkOutLine(s, syncReq, syncRsp)
			rsp.MomotalkOutlineResponse = syncRsp.(*proto.MomoTalkOutLineResponse)
		case proto.Protocol_Scenario_List:
			ScenarioList(s, syncReq, syncRsp)
			rsp.ScenarioListResponse = syncRsp.(*proto.ScenarioListResponse)
		case proto.Protocol_Shop_GachaRecruitList:
			ShopGachaRecruitList(s, syncReq, syncRsp)
			rsp.ShopGachaRecruitListResponse = syncRsp.(*proto.ShopGachaRecruitListResponse)
		case proto.Protocol_TimeAttackDungeon_Login:
			TimeAttackDungeonLogin(s, syncReq, syncRsp)
			rsp.TimeAttackDungeonLoginResponse = syncRsp.(*proto.TimeAttackDungeonLoginResponse)
		case proto.Protocol_Billing_PurchaseListByYostar:
			BillingPurchaseListByYostar(s, syncReq, syncRsp)
			rsp.BillingPurchaseListByYostarResponse = syncRsp.(*proto.BillingPurchaseListByYostarResponse)
		case proto.Protocol_EventContent_PermanentList:
			EventContentPermanentList(s, syncReq, syncRsp)
			rsp.EventContentPermanentListResponse = syncRsp.(*proto.EventContentPermanentListResponse)
		case proto.Protocol_Attachment_Get:
			AttachmentGet(s, syncReq, syncRsp)
			rsp.AttachmentGetResponse = syncRsp.(*proto.AttachmentGetResponse)
		case proto.Protocol_Attachment_EmblemList:
			AttachmentEmblemList(s, syncReq, syncRsp)
			rsp.AttachmentEmblemListResponse = syncRsp.(*proto.AttachmentEmblemListResponse)
		case proto.Protocol_Sticker_Login:
			StickerLogin(s, syncReq, syncRsp)
			rsp.StickerListResponse = syncRsp.(*proto.StickerLoginResponse)
		case proto.Protocol_MultiFloorRaid_Sync:
			MultiFloorRaidSync(s, syncReq, syncRsp)
			rsp.MultiFloorRaidSyncResponse = syncRsp.(*proto.MultiFloorRaidSyncResponse)
		case proto.Protocol_ContentSweep_MultiSweepPresetList:
			ContentSweepMultiSweepPresetList(s, syncReq, syncRsp)
			rsp.ContentSweepMultiSweepPresetListResponse = syncRsp.(*proto.ContentSweepMultiSweepPresetListResponse)
		default:
			logger.Error("AccountLoginSync 没有处理的数据:%s", cmdId.String())
		}
	}
}

func AccountCurrencySync(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.AccountCurrencySyncResponse)

	rsp.AccountCurrencyDB = game.GetAccountCurrencyDB(s)
	rsp.ExpiredCurrency = make(map[proto.CurrencyTypes]int64)
}

func EchelonList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.EchelonListResponse)

	rsp.EchelonDBs = make([]*proto.EchelonDB, 0)

	echelonDB := &proto.EchelonDB{
		AccountServerId:               s.AccountServerId,
		EchelonType:                   proto.EchelonType_Adventure,
		EchelonNumber:                 1,
		ExtensionType:                 proto.EchelonExtensionType_Base,
		LeaderServerId:                game.GetServerId(),
		MainSlotServerIds:             []int64{game.GetServerId(), 0, 0, 0},
		SupportSlotServerIds:          []int64{game.GetServerId(), 0},
		TSSInteractionServerId:        0,
		UsingFlag:                     0,
		SkillCardMulliganCharacterIds: make([]int64, 0),
		CombatStyleIndex:              []int{0, 0, 0, 0, 0, 0},
	}
	rsp.EchelonDBs = append(rsp.EchelonDBs, echelonDB)
}

func ContentSaveGet(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.ContentSaveGetResponse)

	rsp.ServerNotification = proto.ServerNotificationFlag_HasUnreadMail
}

func ProofTokenSubmit(s *enter.Session, request, response mx.Message) {

}

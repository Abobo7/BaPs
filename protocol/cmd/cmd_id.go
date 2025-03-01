package cmd

import (
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(mx.Protocol_Queuing_GetTicket, func() any { return new(proto.QueuingGetTicketRequest) }, true)
	c.regMsg(mx.Protocol_Queuing_GetTicket, func() any { return new(proto.QueuingGetTicketResponse) }, false)
	c.regMsg(mx.Protocol_Account_CheckYostar, func() any { return new(proto.AccountCheckYostarRequest) }, true)
	c.regMsg(mx.Protocol_Account_CheckYostar, func() any { return new(proto.AccountCheckYostarResponse) }, false)
	c.regMsg(mx.Protocol_Account_Auth, func() any { return new(proto.AccountAuthRequest) }, true)
	c.regMsg(mx.Protocol_Account_Auth, func() any { return new(proto.AccountAuthResponse) }, false)
	c.regMsg(mx.Protocol_Account_Nickname, func() any { return new(proto.AccountNicknameRequest) }, true)
	c.regMsg(mx.Protocol_Account_Nickname, func() any { return new(proto.AccountNicknameResponse) }, false)
	c.regMsg(mx.Protocol_ProofToken_RequestQuestion, func() any { return new(proto.ProofTokenRequestQuestionRequest) }, true)
	c.regMsg(mx.Protocol_ProofToken_RequestQuestion, func() any { return new(proto.ProofTokenRequestQuestionResponse) }, false)
	c.regMsg(mx.Protocol_NetworkTime_Sync, func() any { return new(proto.NetworkTimeSyncRequest) }, true)
	c.regMsg(mx.Protocol_NetworkTime_Sync, func() any { return new(proto.NetworkTimeSyncResponse) }, false)
	c.regMsg(mx.Protocol_Academy_GetInfo, func() any { return new(proto.AcademyGetInfoRequest) }, true)
	c.regMsg(mx.Protocol_Academy_GetInfo, func() any { return new(proto.AcademyGetInfoResponse) }, false)
	c.regMsg(mx.Protocol_Account_LoginSync, func() any { return new(proto.AccountLoginSyncRequest) }, true)
	c.regMsg(mx.Protocol_Account_LoginSync, func() any { return new(proto.AccountLoginSyncResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Get, func() any { return new(proto.CafeGetInfoRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Get, func() any { return new(proto.CafeGetInfoResponse) }, false)
	c.regMsg(mx.Protocol_Account_CurrencySync, func() any { return new(proto.AccountCurrencySyncRequest) }, true)
	c.regMsg(mx.Protocol_Account_CurrencySync, func() any { return new(proto.AccountCurrencySyncResponse) }, false)
	c.regMsg(mx.Protocol_Character_List, func() any { return new(proto.CharacterListRequest) }, true)
	c.regMsg(mx.Protocol_Character_List, func() any { return new(proto.CharacterListResponse) }, false)
	c.regMsg(mx.Protocol_Item_List, func() any { return new(proto.ItemListRequest) }, true)
	c.regMsg(mx.Protocol_Item_List, func() any { return new(proto.ItemListResponse) }, false)
	c.regMsg(mx.Protocol_ContentSave_Get, func() any { return new(proto.ContentSaveGetRequest) }, true)
	c.regMsg(mx.Protocol_ContentSave_Get, func() any { return new(proto.ContentSaveGetResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaGet, func() any { return new(proto.ShopBeforehandGachaGetRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaGet, func() any { return new(proto.ShopBeforehandGachaGetResponse) }, false)
	c.regMsg(mx.Protocol_ProofToken_Submit, func() any { return new(proto.ProofTokenSubmitRequest) }, true)
	c.regMsg(mx.Protocol_ProofToken_Submit, func() any { return new(proto.ProofTokenSubmitResponse) }, false)
	c.regMsg(mx.Protocol_Equipment_List, func() any { return new(proto.EquipmentItemListRequest) }, true)
	c.regMsg(mx.Protocol_Equipment_List, func() any { return new(proto.EquipmentItemListResponse) }, false)
	c.regMsg(mx.Protocol_CharacterGear_List, func() any { return new(proto.CharacterGearListRequest) }, true)
	c.regMsg(mx.Protocol_CharacterGear_List, func() any { return new(proto.CharacterGearListResponse) }, false)
	c.regMsg(mx.Protocol_Echelon_List, func() any { return new(proto.EchelonListRequest) }, true)
	c.regMsg(mx.Protocol_Echelon_List, func() any { return new(proto.EchelonListResponse) }, false)
	c.regMsg(mx.Protocol_MemoryLobby_List, func() any { return new(proto.MemoryLobbyListRequest) }, true)
	c.regMsg(mx.Protocol_MemoryLobby_List, func() any { return new(proto.MemoryLobbyListResponse) }, false)
	c.regMsg(mx.Protocol_Campaign_List, func() any { return new(proto.CampaignListRequest) }, true)
	c.regMsg(mx.Protocol_Campaign_List, func() any { return new(proto.CampaignListResponse) }, false)
	c.regMsg(mx.Protocol_Arena_Login, func() any { return new(proto.ArenaLoginRequest) }, true)
	c.regMsg(mx.Protocol_Arena_Login, func() any { return new(proto.ArenaLoginResponse) }, false)
	c.regMsg(mx.Protocol_Raid_Login, func() any { return new(proto.RaidLoginRequest) }, true)
	c.regMsg(mx.Protocol_Raid_Login, func() any { return new(proto.RaidLoginResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_Login, func() any { return new(proto.EliminateRaidLoginRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_Login, func() any { return new(proto.EliminateRaidLoginResponse) }, false)
	c.regMsg(mx.Protocol_Craft_List, func() any { return new(proto.CraftInfoListRequest) }, true)
	c.regMsg(mx.Protocol_Craft_List, func() any { return new(proto.CraftInfoListResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Login, func() any { return new(proto.ClanLoginRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Login, func() any { return new(proto.ClanLoginResponse) }, false)
	c.regMsg(mx.Protocol_MomoTalk_OutLine, func() any { return new(proto.MomoTalkOutLineRequest) }, true)
	c.regMsg(mx.Protocol_MomoTalk_OutLine, func() any { return new(proto.MomoTalkOutLineResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_List, func() any { return new(proto.ScenarioListRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_List, func() any { return new(proto.ScenarioListResponse) }, false)
	c.regMsg(mx.Protocol_Shop_GachaRecruitList, func() any { return new(proto.ShopGachaRecruitListRequest) }, true)
	c.regMsg(mx.Protocol_Shop_GachaRecruitList, func() any { return new(proto.ShopGachaRecruitListResponse) }, false)
	c.regMsg(mx.Protocol_TimeAttackDungeon_Login, func() any { return new(proto.TimeAttackDungeonLoginRequest) }, true)
	c.regMsg(mx.Protocol_TimeAttackDungeon_Login, func() any { return new(proto.TimeAttackDungeonLoginResponse) }, false)
	c.regMsg(mx.Protocol_Billing_PurchaseListByYostar, func() any { return new(proto.BillingPurchaseListByYostarRequest) }, true)
	c.regMsg(mx.Protocol_Billing_PurchaseListByYostar, func() any { return new(proto.BillingPurchaseListByYostarResponse) }, false)
	c.regMsg(mx.Protocol_EventContent_PermanentList, func() any { return new(proto.EventContentPermanentListRequest) }, true)
	c.regMsg(mx.Protocol_EventContent_PermanentList, func() any { return new(proto.EventContentPermanentListResponse) }, false)
	c.regMsg(mx.Protocol_Attachment_Get, func() any { return new(proto.AttachmentGetRequest) }, true)
	c.regMsg(mx.Protocol_Attachment_Get, func() any { return new(proto.AttachmentGetResponse) }, false)
	c.regMsg(mx.Protocol_Attachment_EmblemList, func() any { return new(proto.AttachmentEmblemListRequest) }, true)
	c.regMsg(mx.Protocol_Attachment_EmblemList, func() any { return new(proto.AttachmentEmblemListResponse) }, false)
	c.regMsg(mx.Protocol_Sticker_Login, func() any { return new(proto.StickerLoginRequest) }, true)
	c.regMsg(mx.Protocol_Sticker_Login, func() any { return new(proto.StickerLoginResponse) }, false)
	c.regMsg(mx.Protocol_MultiFloorRaid_Sync, func() any { return new(proto.MultiFloorRaidSyncRequest) }, true)
	c.regMsg(mx.Protocol_MultiFloorRaid_Sync, func() any { return new(proto.MultiFloorRaidSyncResponse) }, false)
	c.regMsg(mx.Protocol_ContentSweep_MultiSweepPresetList, func() any { return new(proto.ContentSweepMultiSweepPresetListRequest) }, true)
	c.regMsg(mx.Protocol_ContentSweep_MultiSweepPresetList, func() any { return new(proto.ContentSweepMultiSweepPresetListResponse) }, false)
	c.regMsg(mx.Protocol_Account_GetTutorial, func() any { return new(proto.AccountGetTutorialRequest) }, true)
	c.regMsg(mx.Protocol_Account_GetTutorial, func() any { return new(proto.AccountGetTutorialResponse) }, false)
	c.regMsg(mx.Protocol_Mission_List, func() any { return new(proto.MissionListRequest) }, true)
	c.regMsg(mx.Protocol_Mission_List, func() any { return new(proto.MissionListResponse) }, false)
	c.regMsg(mx.Protocol_Mission_GuideMissionSeasonList, func() any { return new(proto.GuideMissionSeasonListRequest) }, true)
	c.regMsg(mx.Protocol_Mission_GuideMissionSeasonList, func() any { return new(proto.GuideMissionSeasonListResponse) }, false)
	c.regMsg(mx.Protocol_Toast_List, func() any { return new(proto.ToastListRequest) }, true)
	c.regMsg(mx.Protocol_Toast_List, func() any { return new(proto.ToastListResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_Skip, func() any { return new(proto.ScenarioSkipRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_Skip, func() any { return new(proto.ScenarioSkipResponse) }, false)
	c.regMsg(mx.Protocol_Account_SetTutorial, func() any { return new(proto.AccountSetTutorialRequest) }, true)
	c.regMsg(mx.Protocol_Account_SetTutorial, func() any { return new(proto.AccountSetTutorialResponse) }, false)
	c.regMsg(mx.Protocol_Event_RewardIncrease, func() any { return new(proto.EventRewardIncreaseRequest) }, true)
	c.regMsg(mx.Protocol_Event_RewardIncrease, func() any { return new(proto.EventRewardIncreaseResponse) }, false)
	c.regMsg(mx.Protocol_Mission_Sync, func() any { return new(proto.MissionSyncRequest) }, true)
	c.regMsg(mx.Protocol_Mission_Sync, func() any { return new(proto.MissionSyncResponse) }, false)
	c.regMsg(mx.Protocol_OpenCondition_EventList, func() any { return new(proto.OpenConditionEventListRequest) }, true)
	c.regMsg(mx.Protocol_OpenCondition_EventList, func() any { return new(proto.OpenConditionEventListResponse) }, false)
	c.regMsg(mx.Protocol_Notification_EventContentReddotCheck, func() any { return new(proto.NotificationEventContentReddotRequest) }, true)
	c.regMsg(mx.Protocol_Notification_EventContentReddotCheck, func() any { return new(proto.NotificationEventContentReddotResponse) }, false)
	c.regMsg(mx.Protocol_Mail_Check, func() any { return new(proto.MailCheckRequest) }, true)
	c.regMsg(mx.Protocol_Mail_Check, func() any { return new(proto.MailCheckResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Check, func() any { return new(proto.ClanCheckRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Check, func() any { return new(proto.ClanCheckResponse) }, false)
	c.regMsg(mx.Protocol_Friend_Check, func() any { return new(proto.FriendCheckRequest) }, true)
	c.regMsg(mx.Protocol_Friend_Check, func() any { return new(proto.FriendCheckResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaRun, func() any { return new(proto.ShopBeforehandGachaRunRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaRun, func() any { return new(proto.ShopBeforehandGachaRunResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaSave, func() any { return new(proto.ShopBeforehandGachaSaveRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaSave, func() any { return new(proto.ShopBeforehandGachaSaveResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaPick, func() any { return new(proto.ShopBeforehandGachaPickRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BeforehandGachaPick, func() any { return new(proto.ShopBeforehandGachaPickResponse) }, false)
	c.regMsg(mx.Protocol_ContentLog_UIOpenStatistics, func() any { return new(proto.ContentLogUIOpenStatisticsRequest) }, true)
	c.regMsg(mx.Protocol_ContentLog_UIOpenStatistics, func() any { return new(proto.ContentLogUIOpenStatisticsResponse) }, false)
	c.regMsg(mx.Protocol_Shop_List, func() any { return new(proto.ShopListRequest) }, true)
	c.regMsg(mx.Protocol_Shop_List, func() any { return new(proto.ShopListResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_GroupHistoryUpdate, func() any { return new(proto.ScenarioGroupHistoryUpdateRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_GroupHistoryUpdate, func() any { return new(proto.ScenarioGroupHistoryUpdateResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_Clear, func() any { return new(proto.ScenarioClearRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_Clear, func() any { return new(proto.ScenarioClearResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_Select, func() any { return new(proto.ScenarioSelectRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_Select, func() any { return new(proto.ScenarioSelectResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BuyGacha, func() any { return new(proto.ShopBuyGachaRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BuyGacha, func() any { return new(proto.ShopBuyGachaResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BuyGacha2, func() any { return new(proto.ShopBuyGacha2Request) }, true)
	c.regMsg(mx.Protocol_Shop_BuyGacha2, func() any { return new(proto.ShopBuyGacha2Response) }, false)
	c.regMsg(mx.Protocol_Shop_BuyGacha3, func() any { return new(proto.ShopBuyGacha3Request) }, true)
	c.regMsg(mx.Protocol_Shop_BuyGacha3, func() any { return new(proto.ShopBuyGacha3Response) }, false)
	c.regMsg(mx.Protocol_Echelon_Save, func() any { return new(proto.EchelonSaveRequest) }, true)
	c.regMsg(mx.Protocol_Echelon_Save, func() any { return new(proto.EchelonSaveResponse) }, false)
	c.regMsg(mx.Protocol_Echelon_PresetList, func() any { return new(proto.EchelonPresetListRequest) }, true)
	c.regMsg(mx.Protocol_Echelon_PresetList, func() any { return new(proto.EchelonPresetListResponse) }, false)
	c.regMsg(mx.Protocol_Echelon_PresetSave, func() any { return new(proto.EchelonPresetSaveRequest) }, true)
	c.regMsg(mx.Protocol_Echelon_PresetSave, func() any { return new(proto.EchelonPresetSaveResponse) }, false)
	c.regMsg(mx.Protocol_Echelon_PresetGroupRename, func() any { return new(proto.EchelonPresetGroupRenameRequest) }, true)
	c.regMsg(mx.Protocol_Echelon_PresetGroupRename, func() any { return new(proto.EchelonPresetGroupRenameResponse) }, false)
	c.regMsg(mx.Protocol_Attachment_EmblemAcquire, func() any { return new(proto.AttachmentEmblemAcquireRequest) }, true)
	c.regMsg(mx.Protocol_Attachment_EmblemAcquire, func() any { return new(proto.AttachmentEmblemAcquireResponse) }, false)
	c.regMsg(mx.Protocol_Attachment_EmblemAttach, func() any { return new(proto.AttachmentEmblemAttachRequest) }, true)
	c.regMsg(mx.Protocol_Attachment_EmblemAttach, func() any { return new(proto.AttachmentEmblemAttachResponse) }, false)
	c.regMsg(mx.Protocol_Account_SetRepresentCharacterAndComment, func() any { return new(proto.AccountSetRepresentCharacterAndCommentRequest) }, true)
	c.regMsg(mx.Protocol_Account_SetRepresentCharacterAndComment, func() any { return new(proto.AccountSetRepresentCharacterAndCommentResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_LobbyStudentChange, func() any { return new(proto.ScenarioLobbyStudentChangeRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_LobbyStudentChange, func() any { return new(proto.ScenarioLobbyStudentChangeResponse) }, false)
	c.regMsg(mx.Protocol_Scenario_AccountStudentChange, func() any { return new(proto.ScenarioAccountStudentChangeRequest) }, true)
	c.regMsg(mx.Protocol_Scenario_AccountStudentChange, func() any { return new(proto.ScenarioAccountStudentChangeResponse) }, false)
	c.regMsg(mx.Protocol_MomoTalk_MessageList, func() any { return new(proto.MomoTalkMessageListRequest) }, true)
	c.regMsg(mx.Protocol_MomoTalk_MessageList, func() any { return new(proto.MomoTalkMessageListResponse) }, false)
	c.regMsg(mx.Protocol_MomoTalk_Read, func() any { return new(proto.MomoTalkReadRequest) }, true)
	c.regMsg(mx.Protocol_MomoTalk_Read, func() any { return new(proto.MomoTalkReadResponse) }, false)
	c.regMsg(mx.Protocol_MomoTalk_FavorSchedule, func() any { return new(proto.MomoTalkFavorScheduleRequest) }, true)
	c.regMsg(mx.Protocol_MomoTalk_FavorSchedule, func() any { return new(proto.MomoTalkFavorScheduleResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BuyMerchandise, func() any { return new(proto.ShopBuyMerchandiseRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BuyMerchandise, func() any { return new(proto.ShopBuyMerchandiseResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BuyEligma, func() any { return new(proto.ShopBuyEligmaRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BuyEligma, func() any { return new(proto.ShopBuyEligmaResponse) }, false)
	c.regMsg(mx.Protocol_Shop_BuyRefreshMerchandise, func() any { return new(proto.ShopBuyRefreshMerchandiseRequest) }, true)
	c.regMsg(mx.Protocol_Shop_BuyRefreshMerchandise, func() any { return new(proto.ShopBuyRefreshMerchandiseResponse) }, false)
	c.regMsg(mx.Protocol_Mail_List, func() any { return new(proto.MailListRequest) }, true)
	c.regMsg(mx.Protocol_Mail_List, func() any { return new(proto.MailListResponse) }, false)
	c.regMsg(mx.Protocol_Mail_Receive, func() any { return new(proto.MailReceiveRequest) }, true)
	c.regMsg(mx.Protocol_Mail_Receive, func() any { return new(proto.MailReceiveResponse) }, false)
	c.regMsg(mx.Protocol_Character_Transcendence, func() any { return new(proto.CharacterTranscendenceRequest) }, true)
	c.regMsg(mx.Protocol_Character_Transcendence, func() any { return new(proto.CharacterTranscendenceResponse) }, false)
	c.regMsg(mx.Protocol_Character_UnlockWeapon, func() any { return new(proto.CharacterUnlockWeaponRequest) }, true)
	c.regMsg(mx.Protocol_Character_UnlockWeapon, func() any { return new(proto.CharacterUnlockWeaponResponse) }, false)
	c.regMsg(mx.Protocol_Character_WeaponTranscendence, func() any { return new(proto.CharacterWeaponTranscendenceRequest) }, true)
	c.regMsg(mx.Protocol_Character_WeaponTranscendence, func() any { return new(proto.CharacterWeaponTranscendenceResponse) }, false)
	c.regMsg(mx.Protocol_Character_SetFavorites, func() any { return new(proto.CharacterSetFavoritesRequest) }, true)
	c.regMsg(mx.Protocol_Character_SetFavorites, func() any { return new(proto.CharacterSetFavoritesResponse) }, false)
	c.regMsg(mx.Protocol_Character_UpdateSkillLevel, func() any { return new(proto.CharacterSkillLevelUpdateRequest) }, true)
	c.regMsg(mx.Protocol_Character_UpdateSkillLevel, func() any { return new(proto.CharacterSkillLevelUpdateResponse) }, false)
	c.regMsg(mx.Protocol_Character_BatchSkillLevelUpdate, func() any { return new(proto.CharacterBatchSkillLevelUpdateRequest) }, true)
	c.regMsg(mx.Protocol_Character_BatchSkillLevelUpdate, func() any { return new(proto.CharacterBatchSkillLevelUpdateResponse) }, false)
	c.regMsg(mx.Protocol_Campaign_EnterMainStage, func() any { return new(proto.CampaignEnterMainStageRequest) }, true)
	c.regMsg(mx.Protocol_Campaign_EnterMainStage, func() any { return new(proto.CampaignEnterMainStageResponse) }, false)
	c.regMsg(mx.Protocol_Campaign_ChapterClearReward, func() any { return new(proto.CampaignChapterClearRewardRequest) }, true)
	c.regMsg(mx.Protocol_Campaign_ChapterClearReward, func() any { return new(proto.CampaignChapterClearRewardResponse) }, false)
	c.regMsg(mx.Protocol_WeekDungeon_List, func() any { return new(proto.WeekDungeonListRequest) }, true)
	c.regMsg(mx.Protocol_WeekDungeon_List, func() any { return new(proto.WeekDungeonListResponse) }, false)
	c.regMsg(mx.Protocol_WeekDungeon_EnterBattle, func() any { return new(proto.WeekDungeonEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_WeekDungeon_EnterBattle, func() any { return new(proto.WeekDungeonEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_WeekDungeon_BattleResult, func() any { return new(proto.WeekDungeonBattleResultRequest) }, true)
	c.regMsg(mx.Protocol_WeekDungeon_BattleResult, func() any { return new(proto.WeekDungeonBattleResultResponse) }, false)
	c.regMsg(mx.Protocol_WeekDungeon_Retreat, func() any { return new(proto.WeekDungeonRetreatRequest) }, true)
	c.regMsg(mx.Protocol_WeekDungeon_Retreat, func() any { return new(proto.WeekDungeonRetreatResponse) }, false)
	c.regMsg(mx.Protocol_SchoolDungeon_List, func() any { return new(proto.SchoolDungeonListRequest) }, true)
	c.regMsg(mx.Protocol_SchoolDungeon_List, func() any { return new(proto.SchoolDungeonListResponse) }, false)
	c.regMsg(mx.Protocol_SchoolDungeon_EnterBattle, func() any { return new(proto.SchoolDungeonEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_SchoolDungeon_EnterBattle, func() any { return new(proto.SchoolDungeonEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_SchoolDungeon_BattleResult, func() any { return new(proto.SchoolDungeonBattleResultRequest) }, true)
	c.regMsg(mx.Protocol_SchoolDungeon_BattleResult, func() any { return new(proto.SchoolDungeonBattleResultResponse) }, false)
	c.regMsg(mx.Protocol_SchoolDungeon_Retreat, func() any { return new(proto.SchoolDungeonRetreatRequest) }, true)
	c.regMsg(mx.Protocol_SchoolDungeon_Retreat, func() any { return new(proto.SchoolDungeonRetreatResponse) }, false)
	c.regMsg(mx.Protocol_Equipment_Equip, func() any { return new(proto.EquipmentItemEquipRequest) }, true)
	c.regMsg(mx.Protocol_Equipment_Equip, func() any { return new(proto.EquipmentItemEquipResponse) }, false)
	c.regMsg(mx.Protocol_Equipment_LevelUp, func() any { return new(proto.EquipmentItemLevelUpRequest) }, true)
	c.regMsg(mx.Protocol_Equipment_LevelUp, func() any { return new(proto.EquipmentItemLevelUpResponse) }, false)
	c.regMsg(mx.Protocol_Equipment_TierUp, func() any { return new(proto.EquipmentItemTierUpRequest) }, true)
	c.regMsg(mx.Protocol_Equipment_TierUp, func() any { return new(proto.EquipmentItemTierUpResponse) }, false)
	c.regMsg(mx.Protocol_Equipment_BatchGrowth, func() any { return new(proto.EquipmentBatchGrowthRequest) }, true)
	c.regMsg(mx.Protocol_Equipment_BatchGrowth, func() any { return new(proto.EquipmentBatchGrowthResponse) }, false)
	c.regMsg(mx.Protocol_Character_WeaponExpGrowth, func() any { return new(proto.CharacterWeaponExpGrowthRequest) }, true)
	c.regMsg(mx.Protocol_Character_WeaponExpGrowth, func() any { return new(proto.CharacterWeaponExpGrowthResponse) }, false)
	c.regMsg(mx.Protocol_Raid_Lobby, func() any { return new(proto.RaidLobbyRequest) }, true)
	c.regMsg(mx.Protocol_Raid_Lobby, func() any { return new(proto.RaidLobbyResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Ack, func() any { return new(proto.CafeAckRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Ack, func() any { return new(proto.CafeAckResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Open, func() any { return new(proto.CafeOpenRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Open, func() any { return new(proto.CafeOpenResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Remove, func() any { return new(proto.CafeRemoveFurnitureRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Remove, func() any { return new(proto.CafeRemoveFurnitureResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_RemoveAll, func() any { return new(proto.CafeRemoveAllFurnitureRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_RemoveAll, func() any { return new(proto.CafeRemoveAllFurnitureResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Deploy, func() any { return new(proto.CafeDeployFurnitureRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Deploy, func() any { return new(proto.CafeDeployFurnitureResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Interact, func() any { return new(proto.CafeInteractWithCharacterRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Interact, func() any { return new(proto.CafeInteractWithCharacterResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_SummonCharacter, func() any { return new(proto.CafeSummonCharacterRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_SummonCharacter, func() any { return new(proto.CafeSummonCharacterResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Relocate, func() any { return new(proto.CafeRelocateFurnitureRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Relocate, func() any { return new(proto.CafeRelocateFurnitureResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_RankUp, func() any { return new(proto.CafeRankUpRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_RankUp, func() any { return new(proto.CafeRankUpResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_ReceiveCurrency, func() any { return new(proto.CafeReceiveCurrencyRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_ReceiveCurrency, func() any { return new(proto.CafeReceiveCurrencyResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_ListPreset, func() any { return new(proto.CafeListPresetRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_ListPreset, func() any { return new(proto.CafeListPresetResponse) }, false)
	c.regMsg(mx.Protocol_Friend_List, func() any { return new(proto.FriendListRequest) }, true)
	c.regMsg(mx.Protocol_Friend_List, func() any { return new(proto.FriendListResponse) }, false)
	c.regMsg(mx.Protocol_Friend_GetIdCard, func() any { return new(proto.FriendGetIdCardRequest) }, true)
	c.regMsg(mx.Protocol_Friend_GetIdCard, func() any { return new(proto.FriendGetIdCardResponse) }, false)
	c.regMsg(mx.Protocol_Friend_SetIdCard, func() any { return new(proto.FriendSetIdCardRequest) }, true)
	c.regMsg(mx.Protocol_Friend_SetIdCard, func() any { return new(proto.FriendSetIdCardResponse) }, false)
	c.regMsg(mx.Protocol_Friend_Search, func() any { return new(proto.FriendSearchRequest) }, true)
	c.regMsg(mx.Protocol_Friend_Search, func() any { return new(proto.FriendSearchResponse) }, false)
	c.regMsg(mx.Protocol_Friend_GetFriendDetailedInfo, func() any { return new(proto.FriendGetFriendDetailedInfoRequest) }, true)
	c.regMsg(mx.Protocol_Friend_GetFriendDetailedInfo, func() any { return new(proto.FriendGetFriendDetailedInfoResponse) }, false)
	c.regMsg(mx.Protocol_Friend_SendFriendRequest, func() any { return new(proto.FriendSendFriendRequestRequest) }, true)
	c.regMsg(mx.Protocol_Friend_SendFriendRequest, func() any { return new(proto.FriendSendFriendRequestResponse) }, false)
	c.regMsg(mx.Protocol_Friend_AcceptFriendRequest, func() any { return new(proto.FriendAcceptFriendRequestRequest) }, true)
	c.regMsg(mx.Protocol_Friend_AcceptFriendRequest, func() any { return new(proto.FriendAcceptFriendRequestResponse) }, false)
	c.regMsg(mx.Protocol_Friend_DeclineFriendRequest, func() any { return new(proto.FriendDeclineFriendRequestRequest) }, true)
	c.regMsg(mx.Protocol_Friend_DeclineFriendRequest, func() any { return new(proto.FriendDeclineFriendRequestResponse) }, false)
	c.regMsg(mx.Protocol_Friend_Remove, func() any { return new(proto.FriendRemoveRequest) }, true)
	c.regMsg(mx.Protocol_Friend_Remove, func() any { return new(proto.FriendRemoveResponse) }, false)
	c.regMsg(mx.Protocol_ContentSweep_Request, func() any { return new(proto.ContentSweepRequest) }, true)
	c.regMsg(mx.Protocol_ContentSweep_Request, func() any { return new(proto.ContentSweepResponse) }, false)
	c.regMsg(mx.Protocol_Character_ExpGrowth, func() any { return new(proto.CharacterExpGrowthRequest) }, true)
	c.regMsg(mx.Protocol_Character_ExpGrowth, func() any { return new(proto.CharacterExpGrowthResponse) }, false)
	c.regMsg(mx.Protocol_CharacterGear_Unlock, func() any { return new(proto.CharacterGearUnlockRequest) }, true)
	c.regMsg(mx.Protocol_CharacterGear_Unlock, func() any { return new(proto.CharacterGearUnlockResponse) }, false)
	c.regMsg(mx.Protocol_Character_PotentialGrowth, func() any { return new(proto.CharacterPotentialGrowthRequest) }, true)
	c.regMsg(mx.Protocol_Character_PotentialGrowth, func() any { return new(proto.CharacterPotentialGrowthResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_Travel, func() any { return new(proto.CafeTravelRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_Travel, func() any { return new(proto.CafeTravelResponse) }, false)
	c.regMsg(mx.Protocol_Cafe_GiveGift, func() any { return new(proto.CafeGiveGiftRequest) }, true)
	c.regMsg(mx.Protocol_Cafe_GiveGift, func() any { return new(proto.CafeGiveGiftResponse) }, false)
	c.regMsg(mx.Protocol_Academy_AttendSchedule, func() any { return new(proto.AcademyAttendScheduleRequest) }, true)
	c.regMsg(mx.Protocol_Academy_AttendSchedule, func() any { return new(proto.AcademyAttendScheduleResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Lobby, func() any { return new(proto.ClanLobbyRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Lobby, func() any { return new(proto.ClanLobbyResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Search, func() any { return new(proto.ClanSearchRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Search, func() any { return new(proto.ClanSearchResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Create, func() any { return new(proto.ClanCreateRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Create, func() any { return new(proto.ClanCreateResponse) }, false)
	c.regMsg(mx.Protocol_Clan_MemberList, func() any { return new(proto.ClanMemberListRequest) }, true)
	c.regMsg(mx.Protocol_Clan_MemberList, func() any { return new(proto.ClanMemberListResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Join, func() any { return new(proto.ClanJoinRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Join, func() any { return new(proto.ClanJoinResponse) }, false)
	c.regMsg(mx.Protocol_Clan_AutoJoin, func() any { return new(proto.ClanAutoJoinRequest) }, true)
	c.regMsg(mx.Protocol_Clan_AutoJoin, func() any { return new(proto.ClanAutoJoinResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Setting, func() any { return new(proto.ClanSettingRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Setting, func() any { return new(proto.ClanSettingResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Applicant, func() any { return new(proto.ClanApplicantRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Applicant, func() any { return new(proto.ClanApplicantResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Member, func() any { return new(proto.ClanMemberRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Member, func() any { return new(proto.ClanMemberResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Quit, func() any { return new(proto.ClanQuitRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Quit, func() any { return new(proto.ClanQuitResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Kick, func() any { return new(proto.ClanKickRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Kick, func() any { return new(proto.ClanKickResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Confer, func() any { return new(proto.ClanConferRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Confer, func() any { return new(proto.ClanConferResponse) }, false)
	c.regMsg(mx.Protocol_Clan_Permit, func() any { return new(proto.ClanPermitRequest) }, true)
	c.regMsg(mx.Protocol_Clan_Permit, func() any { return new(proto.ClanPermitResponse) }, false)
	c.regMsg(mx.Protocol_Clan_MyAssistList, func() any { return new(proto.ClanMyAssistListRequest) }, true)
	c.regMsg(mx.Protocol_Clan_MyAssistList, func() any { return new(proto.ClanMyAssistListResponse) }, false)
	c.regMsg(mx.Protocol_Clan_SetAssist, func() any { return new(proto.ClanSetAssistRequest) }, true)
	c.regMsg(mx.Protocol_Clan_SetAssist, func() any { return new(proto.ClanSetAssistResponse) }, false)
	c.regMsg(mx.Protocol_Clan_AllAssistList, func() any { return new(proto.ClanAllAssistListRequest) }, true)
	c.regMsg(mx.Protocol_Clan_AllAssistList, func() any { return new(proto.ClanAllAssistListResponse) }, false)
	c.regMsg(mx.Protocol_Raid_OpponentList, func() any { return new(proto.RaidOpponentListRequest) }, true)
	c.regMsg(mx.Protocol_Raid_OpponentList, func() any { return new(proto.RaidOpponentListResponse) }, false)
	c.regMsg(mx.Protocol_Raid_CreateBattle, func() any { return new(proto.RaidCreateBattleRequest) }, true)
	c.regMsg(mx.Protocol_Raid_CreateBattle, func() any { return new(proto.RaidCreateBattleResponse) }, false)
	c.regMsg(mx.Protocol_Raid_GetBestTeam, func() any { return new(proto.RaidGetBestTeamRequest) }, true)
	c.regMsg(mx.Protocol_Raid_GetBestTeam, func() any { return new(proto.RaidGetBestTeamResponse) }, false)
	c.regMsg(mx.Protocol_Raid_EndBattle, func() any { return new(proto.RaidEndBattleRequest) }, true)
	c.regMsg(mx.Protocol_Raid_EndBattle, func() any { return new(proto.RaidEndBattleResponse) }, false)
	c.regMsg(mx.Protocol_Raid_EnterBattle, func() any { return new(proto.RaidEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_Raid_EnterBattle, func() any { return new(proto.RaidEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_Raid_GiveUp, func() any { return new(proto.RaidGiveUpRequest) }, true)
	c.regMsg(mx.Protocol_Raid_GiveUp, func() any { return new(proto.RaidGiveUpResponse) }, false)
	c.regMsg(mx.Protocol_Raid_RankingReward, func() any { return new(proto.RaidRankingRewardRequest) }, true)
	c.regMsg(mx.Protocol_Raid_RankingReward, func() any { return new(proto.RaidRankingRewardResponse) }, false)
	c.regMsg(mx.Protocol_Raid_SeasonReward, func() any { return new(proto.RaidSeasonRewardRequest) }, true)
	c.regMsg(mx.Protocol_Raid_SeasonReward, func() any { return new(proto.RaidSeasonRewardResponse) }, false)
	c.regMsg(mx.Protocol_Attendance_Reward, func() any { return new(proto.AttendanceRewardRequest) }, true)
	c.regMsg(mx.Protocol_Attendance_Reward, func() any { return new(proto.AttendanceRewardResponse) }, false)
	c.regMsg(mx.Protocol_Mission_Reward, func() any { return new(proto.MissionRewardRequest) }, true)
	c.regMsg(mx.Protocol_Mission_Reward, func() any { return new(proto.MissionRewardResponse) }, false)
	c.regMsg(mx.Protocol_MultiFloorRaid_EnterBattle, func() any { return new(proto.MultiFloorRaidEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_MultiFloorRaid_EnterBattle, func() any { return new(proto.MultiFloorRaidEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_MultiFloorRaid_EndBattle, func() any { return new(proto.MultiFloorRaidEndBattleRequest) }, true)
	c.regMsg(mx.Protocol_MultiFloorRaid_EndBattle, func() any { return new(proto.MultiFloorRaidEndBattleResponse) }, false)
	c.regMsg(mx.Protocol_MultiFloorRaid_ReceiveReward, func() any { return new(proto.MultiFloorRaidReceiveRewardRequest) }, true)
	c.regMsg(mx.Protocol_MultiFloorRaid_ReceiveReward, func() any { return new(proto.MultiFloorRaidReceiveRewardResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_Lobby, func() any { return new(proto.EliminateRaidLobbyRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_Lobby, func() any { return new(proto.EliminateRaidLobbyResponse) }, false)
	c.regMsg(mx.Protocol_TimeAttackDungeon_Lobby, func() any { return new(proto.TimeAttackDungeonLobbyRequest) }, true)
	c.regMsg(mx.Protocol_TimeAttackDungeon_Lobby, func() any { return new(proto.TimeAttackDungeonLobbyResponse) }, false)
	c.regMsg(mx.Protocol_TimeAttackDungeon_CreateBattle, func() any { return new(proto.TimeAttackDungeonCreateBattleRequest) }, true)
	c.regMsg(mx.Protocol_TimeAttackDungeon_CreateBattle, func() any { return new(proto.TimeAttackDungeonCreateBattleResponse) }, false)
	c.regMsg(mx.Protocol_TimeAttackDungeon_EnterBattle, func() any { return new(proto.TimeAttackDungeonEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_TimeAttackDungeon_EnterBattle, func() any { return new(proto.TimeAttackDungeonEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_TimeAttackDungeon_EndBattle, func() any { return new(proto.TimeAttackDungeonEndBattleRequest) }, true)
	c.regMsg(mx.Protocol_TimeAttackDungeon_EndBattle, func() any { return new(proto.TimeAttackDungeonEndBattleResponse) }, false)
	c.regMsg(mx.Protocol_Arena_EnterLobby, func() any { return new(proto.ArenaEnterLobbyRequest) }, true)
	c.regMsg(mx.Protocol_Arena_EnterLobby, func() any { return new(proto.ArenaEnterLobbyResponse) }, false)
	c.regMsg(mx.Protocol_Arena_OpponentList, func() any { return new(proto.ArenaOpponentListRequest) }, true)
	c.regMsg(mx.Protocol_Arena_OpponentList, func() any { return new(proto.ArenaOpponentListResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_CreateBattle, func() any { return new(proto.EliminateRaidCreateBattleRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_CreateBattle, func() any { return new(proto.EliminateRaidCreateBattleResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_EndBattle, func() any { return new(proto.EliminateRaidEndBattleRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_EndBattle, func() any { return new(proto.EliminateRaidEndBattleResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_EnterBattle, func() any { return new(proto.EliminateRaidEnterBattleRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_EnterBattle, func() any { return new(proto.EliminateRaidEnterBattleResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_OpponentList, func() any { return new(proto.EliminateRaidOpponentListRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_OpponentList, func() any { return new(proto.EliminateRaidOpponentListResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_GetBestTeam, func() any { return new(proto.EliminateRaidGetBestTeamRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_GetBestTeam, func() any { return new(proto.EliminateRaidGetBestTeamResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_GiveUp, func() any { return new(proto.EliminateRaidGiveUpRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_GiveUp, func() any { return new(proto.EliminateRaidGiveUpResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_SeasonReward, func() any { return new(proto.EliminateRaidSeasonRewardRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_SeasonReward, func() any { return new(proto.EliminateRaidSeasonRewardResponse) }, false)
	c.regMsg(mx.Protocol_EliminateRaid_RankingReward, func() any { return new(proto.EliminateRaidRankingRewardRequest) }, true)
	c.regMsg(mx.Protocol_EliminateRaid_RankingReward, func() any { return new(proto.EliminateRaidRankingRewardResponse) }, false)
}

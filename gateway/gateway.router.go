package gateway

import (
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gucooing/BaPs/common/check"
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/pack"
	"github.com/gucooing/BaPs/pkg/logger"
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/cmd"
	"github.com/gucooing/BaPs/protocol/proto"
)

type handlerFunc func(s *enter.Session, request, response proto.Message)

var funcRouteMap = map[int32]handlerFunc{
	mx.Protocol_Craft_List:                           pack.CraftInfoList,                       // 获取制造信息
	mx.Protocol_Billing_PurchaseListByYostar:         pack.BillingPurchaseListByYostar,         // Yostar采购清单??
	mx.Protocol_EventContent_PermanentList:           pack.EventContentPermanentList,           // 获取永久时间列表?
	mx.Protocol_Sticker_Login:                        pack.StickerLogin,                        // 登录获取贴纸信息??
	mx.Protocol_ContentSweep_MultiSweepPresetList:    pack.ContentSweepMultiSweepPresetList,    // ????
	mx.Protocol_ContentSave_Get:                      pack.ContentSaveGet,                      // ???
	mx.Protocol_Event_RewardIncrease:                 pack.EventRewardIncrease,                 // ???
	mx.Protocol_Notification_EventContentReddotCheck: pack.NotificationEventContentReddotCheck, // 红点检查?
	mx.Protocol_ContentLog_UIOpenStatistics:          pack.ContentLogUIOpenStatistics,          // 历史ui打开
	// 基础
	mx.Protocol_Account_CheckYostar:        AccountCheckYostar,             // 验证EnterTicket
	mx.Protocol_Account_Auth:               pack.AccountAuth,               // 账号验证
	mx.Protocol_ProofToken_RequestQuestion: pack.ProofTokenRequestQuestion, // 验证登录token
	mx.Protocol_Toast_List:                 pack.ToastList,                 // 公告提示信息
	mx.Protocol_ProofToken_Submit:          pack.ProofTokenSubmit,          // 密钥更新
	mx.Protocol_Account_LoginSync:          pack.AccountLoginSync,          // 同步账号信息
	mx.Protocol_NetworkTime_Sync:           pack.NetworkTimeSync,           // 同步时间
	mx.Protocol_OpenCondition_EventList:    pack.OpenConditionEventList,    // 获取开放事件
	mx.Protocol_ContentSweep_Request:       pack.ContentSweepRequest,       // 战斗扫荡
	// 玩家档案
	mx.Protocol_Account_Nickname:                        pack.AccountNickname,                        // 设置/修改 昵称
	mx.Protocol_Account_CallName:                        pack.AccountCallName,                        // 设定语音称呼
	mx.Protocol_Account_SetRepresentCharacterAndComment: pack.AccountSetRepresentCharacterAndComment, // 修改 签名/值日生
	mx.Protocol_Scenario_AccountStudentChange:           pack.ScenarioAccountStudentChange,           // 更换值日生 二次确认
	mx.Protocol_Scenario_LobbyStudentChange:             pack.ScenarioLobbyStudentChange,             // 修改大厅值日生
	mx.Protocol_Attachment_Get:                          pack.AttachmentGet,                          // 获取玩家称号
	mx.Protocol_Attachment_EmblemList:                   pack.AttachmentEmblemList,                   // 获取解锁的玩家称号
	mx.Protocol_Attachment_EmblemAcquire:                pack.AttachmentEmblemAcquire,                // 客户端解锁称号
	mx.Protocol_Attachment_EmblemAttach:                 pack.AttachmentEmblemAttach,                 // 装备称号
	// 登录奖励
	mx.Protocol_Attendance_Reward: pack.AttendanceReward, // 领取奖励
	// MomoTalk
	mx.Protocol_MemoryLobby_List:       pack.MemoryLobbyList,       // 获取记忆大厅列表
	mx.Protocol_MomoTalk_OutLine:       pack.MomoTalkOutLine,       // 获取MomoTalk信息
	mx.Protocol_MomoTalk_MessageList:   pack.MomoTalkMessageList,   // 获取单个角色的MomoTalk信息
	mx.Protocol_MomoTalk_Read:          pack.MomoTalkRead,          // MomoTalk对话选择
	mx.Protocol_MomoTalk_FavorSchedule: pack.MomoTalkFavorSchedule, // 完成MomoTalk剧情
	// 邮箱
	mx.Protocol_Mail_Check:   pack.MailCheck,   // 邮件检查
	mx.Protocol_Mail_List:    pack.MailList,    // 获取邮件列表
	mx.Protocol_Mail_Receive: pack.MailReceive, // 领取邮件
	// 好友
	mx.Protocol_Friend_Check:                 pack.FriendCheck,                 // 好友检查
	mx.Protocol_Friend_List:                  pack.FriendList,                  // 获取好友详情
	mx.Protocol_Friend_GetIdCard:             pack.FriendGetIdCard,             // 获取账号板
	mx.Protocol_Friend_SetIdCard:             pack.FriendSetIdCard,             // 设置账号板
	mx.Protocol_Friend_Search:                pack.FriendSearch,                // 获取附近的人
	mx.Protocol_Friend_GetFriendDetailedInfo: pack.FriendGetFriendDetailedInfo, // 获取玩家详细信息
	mx.Protocol_Friend_SendFriendRequest:     pack.FriendSendFriendRequest,     // 发送好友申请
	mx.Protocol_Friend_AcceptFriendRequest:   pack.FriendAcceptFriendRequest,   // 同意好友申请
	mx.Protocol_Friend_DeclineFriendRequest:  pack.FriendDeclineFriendRequest,  // 拒绝好友申请
	mx.Protocol_Friend_Remove:                pack.FriendRemove,                // 删除好友
	// 背包
	mx.Protocol_Account_CurrencySync:          pack.AccountCurrencySync,          // 同步账号货币
	mx.Protocol_Item_List:                     pack.ItemList,                     // 获取背包物品
	mx.Protocol_Equipment_List:                pack.EquipmentList,                // 获取装备信息
	mx.Protocol_Equipment_LevelUp:             pack.EquipmentLevelUp,             // 装备升级
	mx.Protocol_Equipment_TierUp:              pack.EquipmentTierUp,              // 装备进阶
	mx.Protocol_Equipment_BatchGrowth:         pack.EquipmentBatchGrowth,         // 装备一键升级
	mx.Protocol_Character_WeaponTranscendence: pack.CharacterWeaponTranscendence, // 角色武器升星
	mx.Protocol_Character_WeaponExpGrowth:     pack.CharacterWeaponExpGrowth,     // 角色武器升级
	// 角色
	mx.Protocol_CharacterGear_List:              pack.CharacterGearList,              // 获取角色爱用品
	mx.Protocol_Character_List:                  pack.CharacterList,                  // 获取角色列表
	mx.Protocol_Character_SetFavorites:          pack.CharacterSetFavorites,          // 标记角色
	mx.Protocol_Character_UpdateSkillLevel:      pack.CharacterUpdateSkillLevel,      // 角色技能升级
	mx.Protocol_Character_BatchSkillLevelUpdate: pack.CharacterBatchSkillLevelUpdate, // 角色技能批量升级
	mx.Protocol_Character_Transcendence:         pack.CharacterTranscendence,         // 角色升星
	mx.Protocol_Character_UnlockWeapon:          pack.CharacterUnlockWeapon,          // 角色解锁武器
	mx.Protocol_Equipment_Equip:                 pack.EquipmentEquip,                 // 装备角色装备
	mx.Protocol_Character_ExpGrowth:             pack.CharacterExpGrowth,             // 角色升级
	mx.Protocol_CharacterGear_Unlock:            pack.CharacterGearUnlock,            // 角色解锁爱用品
	mx.Protocol_Character_PotentialGrowth:       pack.CharacterPotentialGrowth,       // 角色能力解放
	// 队伍
	mx.Protocol_Echelon_List:       pack.EchelonList,       // 获取队伍信息
	mx.Protocol_Echelon_Save:       pack.EchelonSave,       // 保存/更新队伍
	mx.Protocol_Echelon_PresetList: pack.EchelonPresetList, // 获取预设队伍
	// 剧情/教程
	mx.Protocol_Scenario_List:               pack.ScenarioList,               // 获取场景剧情信息
	mx.Protocol_Scenario_GroupHistoryUpdate: pack.ScenarioGroupHistoryUpdate, // 完成场景剧情信息
	mx.Protocol_Scenario_Clear:              pack.ScenarioClear,              // 完成剧情
	mx.Protocol_Scenario_Select:             pack.ScenarioSelect,             // 剧情选择
	mx.Protocol_Account_GetTutorial:         pack.AccountGetTutorial,         // 获取教程
	mx.Protocol_Scenario_Skip:               pack.ScenarioSkip,               // 剧情跳过
	mx.Protocol_Account_SetTutorial:         pack.AccountSetTutorial,         // 设置完成教程
	// 任务/成就
	mx.Protocol_Mission_List:                   pack.MissionList,            // 获取 任务/成就 信息
	mx.Protocol_Mission_Sync:                   pack.MissionSync,            // 同步任务/成就
	mx.Protocol_Mission_GuideMissionSeasonList: pack.GuideMissionSeasonList, // 获取成就信息
	mx.Protocol_Mission_Reward:                 pack.MissionReward,          // 领取 任务/成就 奖励
	// 活动
	mx.Protocol_EventContent_AdventureList:              pack.EventContentAdventureList,              // 获取活动关卡信息
	mx.Protocol_EventContent_BoxGachaShopList:           pack.EventContentBoxGachaShopList,           // 获取活动盲盒商店列表
	mx.Protocol_EventContent_ScenarioGroupHistoryUpdate: pack.EventContentScenarioGroupHistoryUpdate, // 活动剧情完成请求
	// 咖啡馆
	mx.Protocol_Cafe_Get:             pack.CafeGetInfo,         // 获取咖啡馆信息
	mx.Protocol_Cafe_Ack:             pack.CafeAck,             // 确认咖啡馆
	mx.Protocol_Cafe_Open:            pack.CafeOpen,            // 客户端主动解锁咖啡馆
	mx.Protocol_Cafe_Remove:          pack.CafeRemove,          // 收纳部分家具
	mx.Protocol_Cafe_RemoveAll:       pack.CafeRemoveAll,       // 收纳全部家具
	mx.Protocol_Cafe_Deploy:          pack.CafeDeploy,          // 摆放家具
	mx.Protocol_Cafe_Relocate:        pack.CafeRelocate,        // 移动家具
	mx.Protocol_Cafe_Interact:        pack.CafeInteract,        // 摸摸头
	mx.Protocol_Cafe_SummonCharacter: pack.CafeSummonCharacter, // 邀请角色
	mx.Protocol_Cafe_RankUp:          pack.CafeRankUp,          // 升级咖啡馆
	mx.Protocol_Cafe_ReceiveCurrency: pack.CafeReceiveCurrency, // 咖啡馆领取产物
	mx.Protocol_Cafe_ListPreset:      pack.CafeListPreset,      // 获取蓝图列表
	mx.Protocol_Cafe_Travel:          pack.CafeTravel,          // 访问好友咖啡馆
	mx.Protocol_Cafe_GiveGift:        pack.CafeGiveGift,        // 礼物赠送
	// 课程表
	mx.Protocol_Academy_GetInfo:        pack.AcademyGetInfo,        // 获取课程表信息
	mx.Protocol_Academy_AttendSchedule: pack.AcademyAttendSchedule, // 上课
	// 社团
	mx.Protocol_Clan_Login:         pack.ClanLogin,         // 登录获取社团信息
	mx.Protocol_Clan_Check:         pack.ClanCheck,         // 社团检查
	mx.Protocol_Clan_Lobby:         pack.ClanLobby,         // 获取社团大厅信息
	mx.Protocol_Clan_Search:        pack.ClanSearch,        // 搜索社团
	mx.Protocol_Clan_Create:        pack.ClanCreate,        // 创建社团
	mx.Protocol_Clan_MemberList:    pack.ClanMemberList,    // 查看社团详情
	mx.Protocol_Clan_Join:          pack.ClanJoin,          // 申请加入社团
	mx.Protocol_Clan_AutoJoin:      pack.ClanAutoJoin,      // 自动加入社团
	mx.Protocol_Clan_Setting:       pack.ClanSetting,       // 更新社团设置
	mx.Protocol_Clan_Applicant:     pack.ClanApplicant,     // 查询社团申请
	mx.Protocol_Clan_Member:        pack.ClanMember,        // 查询成员详情
	mx.Protocol_Clan_Quit:          pack.ClanQuit,          // 退出社团
	mx.Protocol_Clan_Kick:          pack.ClanKick,          // 管理员删除成员
	mx.Protocol_Clan_Confer:        pack.ClanConfer,        // 授予职位
	mx.Protocol_Clan_Permit:        pack.ClanPermit,        // 同意加入社团
	mx.Protocol_Clan_MyAssistList:  pack.ClanMyAssistList,  // 获取自己的援助角色
	mx.Protocol_Clan_SetAssist:     pack.ClanSetAssist,     // 设置援助角色
	mx.Protocol_Clan_AllAssistList: pack.ClanAllAssistList, // 获取全部援助角色
	// 商店
	mx.Protocol_Shop_List:                  pack.ShopList,                  // 获取商店信息
	mx.Protocol_Shop_BuyRefreshMerchandise: pack.ShopBuyRefreshMerchandise, // 每日刷新商店购买
	mx.Protocol_Shop_BuyEligma:             pack.ShopBuyEligma,             // 角色碎片购买
	// 卡池
	mx.Protocol_Shop_GachaRecruitList:    pack.ShopGachaRecruitList,    // 获取卡池历史数据
	mx.Protocol_Shop_BeforehandGachaGet:  pack.ShopBeforehandGachaGet,  // 获取新手免费十连信息
	mx.Protocol_Shop_BeforehandGachaRun:  pack.ShopBeforehandGachaRun,  // 新手免费十连抽卡请求
	mx.Protocol_Shop_BeforehandGachaSave: pack.ShopBeforehandGachaSave, // 缓存新手免费十连结果
	mx.Protocol_Shop_BeforehandGachaPick: pack.ShopBeforehandGachaPick, // 确定新手卡池免费十连结果
	mx.Protocol_Shop_BuyGacha3:           pack.ShopBuyGacha3,           // 卡池3抽卡请求
	// 任务
	mx.Protocol_Campaign_List:                        pack.CampaignList,                        // 获取任务信息
	mx.Protocol_Campaign_EnterMainStage:              pack.CampaignEnterMainStage,              // 进入任务
	mx.Protocol_Campaign_ChapterClearReward:          pack.CampaignChapterClearReward,          // 领取总关卡奖励
	mx.Protocol_Campaign_EnterMainStageStrategySkip:  pack.CampaignEnterMainStageStrategySkip,  // 简易攻略
	mx.Protocol_Campaign_MainStageStrategySkipResult: pack.CampaignMainStageStrategySkipResult, // 简易攻略结算
	// 悬赏通缉/特别依赖
	mx.Protocol_WeekDungeon_List:         pack.WeekDungeonList,         // 获取 悬赏通缉/特别依赖 通关信息
	mx.Protocol_WeekDungeon_EnterBattle:  pack.WeekDungeonEnterBattle,  // 开始战斗
	mx.Protocol_WeekDungeon_BattleResult: pack.WeekDungeonBattleResult, // 战斗结算
	// 学院交流会
	mx.Protocol_SchoolDungeon_List:         pack.SchoolDungeonList,         // 获取学院交流会通关信息
	mx.Protocol_SchoolDungeon_EnterBattle:  pack.SchoolDungeonEnterBattle,  // 开始战斗
	mx.Protocol_SchoolDungeon_BattleResult: pack.SchoolDungeonBattleResult, // 战斗结算
	// 综合战术考试
	mx.Protocol_TimeAttackDungeon_Login:        pack.TimeAttackDungeonLogin,        // 登录获取综合战术考试信息
	mx.Protocol_TimeAttackDungeon_Lobby:        pack.TimeAttackDungeonLobby,        // 获取综合战术考试详情
	mx.Protocol_TimeAttackDungeon_CreateBattle: pack.TimeAttackDungeonCreateBattle, // 综合战术考试开始战斗
	mx.Protocol_TimeAttackDungeon_EnterBattle:  pack.TimeAttackDungeonEnterBattle,  // 战术综合考试开始
	mx.Protocol_TimeAttackDungeon_EndBattle:    pack.TimeAttackDungeonEndBattle,    // 战术综合考试战斗结算
	mx.Protocol_TimeAttackDungeon_GiveUp:       pack.TimeAttackDungeonGiveUp,       // 主动结算
	mx.Protocol_TimeAttackDungeon_Sweep:        pack.TimeAttackDungeonSweep,        // 扫荡
	// 竞技场
	mx.Protocol_Arena_Login:        pack.ArenaLogin,        // 登录获取竞技场信息
	mx.Protocol_Arena_EnterLobby:   pack.ArenaEnterLobby,   // 获取竞技场详情
	mx.Protocol_Arena_OpponentList: pack.ArenaOpponentList, // 获取竞技场战斗列表
	// 总力战
	mx.Protocol_Raid_Login:         pack.RaidLogin,         // 登录获取总力战开放信息
	mx.Protocol_Raid_Lobby:         pack.RaidLobby,         // 获取总力战详情
	mx.Protocol_Raid_OpponentList:  pack.RaidOpponentList,  // 获取总力战排行榜
	mx.Protocol_Raid_GetBestTeam:   pack.RaidGetBestTeam,   // 查询玩家总力战参加队伍
	mx.Protocol_Raid_CreateBattle:  pack.RaidCreateBattle,  // 开始总力战
	mx.Protocol_Raid_EndBattle:     pack.RaidEndBattle,     // 战斗结算
	mx.Protocol_Raid_EnterBattle:   pack.RaidEnterBattle,   // 再次进入战斗
	mx.Protocol_Raid_GiveUp:        pack.RaidGiveUp,        // 主动结束总力战
	mx.Protocol_Raid_SeasonReward:  pack.RaidSeasonReward,  // 领取总分奖励
	mx.Protocol_Raid_RankingReward: pack.RaidRankingReward, // 领取排名奖励
	// 大决战
	mx.Protocol_EliminateRaid_Login:         pack.EliminateRaidLogin,         // 登录获取大决战信息
	mx.Protocol_EliminateRaid_Lobby:         pack.EliminateRaidLobby,         // 获取大决战详情
	mx.Protocol_EliminateRaid_OpponentList:  pack.EliminateRaidOpponentList,  // 获取大决战排行榜
	mx.Protocol_EliminateRaid_GetBestTeam:   pack.EliminateRaidGetBestTeam,   // 查询玩家大决战参加队伍
	mx.Protocol_EliminateRaid_CreateBattle:  pack.EliminateRaidCreateBattle,  // 大决战开始战斗
	mx.Protocol_EliminateRaid_EndBattle:     pack.EliminateRaidEndBattle,     // 大决战战斗结算
	mx.Protocol_EliminateRaid_EnterBattle:   pack.EliminateRaidEnterBattle,   // 再次进入战斗
	mx.Protocol_EliminateRaid_GiveUp:        pack.EliminateRaidGiveUp,        // 主动结算
	mx.Protocol_EliminateRaid_SeasonReward:  pack.EliminateRaidSeasonReward,  // 领取总分奖励
	mx.Protocol_EliminateRaid_RankingReward: pack.EliminateRaidRankingReward, // 领取排名奖励
	// 制约解除决战
	mx.Protocol_MultiFloorRaid_Sync:          pack.MultiFloorRaidSync,          // 制约解除决战信息同步
	mx.Protocol_MultiFloorRaid_EnterBattle:   pack.MultiFloorRaidEnterBattle,   // 制约解除决战开始战斗
	mx.Protocol_MultiFloorRaid_EndBattle:     pack.MultiFloorRaidEndBattle,     // 制约解除决战结束战斗
	mx.Protocol_MultiFloorRaid_ReceiveReward: pack.MultiFloorRaidReceiveReward, // 制约解除决战领取奖励
}

func (g *Gateway) registerMessage(c *gin.Context, request proto.Message, base *proto.BasePacket) {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			errBestHTTP(c, 15022)
			logger.Error("@LogTag(player_panic)cmdId:%s json:%s\nerr:%s\nstack:%s", cmd.Get().GetCmdNameByCmdId(request.GetProtocol()),
				request.String(), err, logger.Stack())
			return
		}
	}()
	handler, ok := funcRouteMap[request.GetProtocol()]
	if !ok {
		errBestHTTP(c, 15022)
		logPlayerMsg(NoRoute, request)
		return
	}
	response := cmd.Get().GetResponsePacketByCmdId(request.GetProtocol())
	if response == nil {
		errBestHTTP(c, 15022)
		logger.Debug("response unknown cmd id: %v\n", cmd.Get().GetCmdNameByCmdId(request.GetProtocol()))
		return
	}
	sessionKey := base.GetSessionKey()
	var s *enter.Session
	if sessionKey == nil &&
		request.GetProtocol() != mx.Protocol_Account_CheckYostar {
		errTokenBestHTTP(c) // TODO 异常请求
		logger.Debug("get request sessionKey nil")
		return
	} else if request.GetProtocol() != mx.Protocol_Account_CheckYostar {
		s = enter.GetSessionBySessionKey(sessionKey)
		if s == nil {
			errTokenBestHTTP(c) // TODO 异常请求
			logger.Debug("get session nil,SessionKey:%s", sessionKey.String())
			return
		}
		s.EndTime = time.Now().Add(time.Duration(enter.MaxCachePlayerTime) * time.Minute)
	}
	response.SetSessionKey(base) //  任何情况下都不要更改handler执行和SetSessionKey的顺序
	if s != nil {                // 唯一线程操作锁
		s.GoroutinesSync.Lock()
		defer s.GoroutinesSync.Unlock()
	}

	// 计时并执行函数
	atomic.AddInt64(&check.TPS, 1)
	time1 := time.Now()
	handler(s, request, response)
	atomic.AddInt64(&check.RT, int64(time.Now().Sub(time1)/time.Nanosecond))

	// 函数执行完毕
	base.ServerTimeTicks = game.GetServerTime()
	base.MissionProgressDBs = game.GetMissionProgressDBs(s)
	base.ServerNotification = game.GetServerNotification(s)
	logPlayerMsg(Client, request)
	logPlayerMsg(Server, response)
	if base.ErrorCode != 0 {
		errBestHTTP(c, base.ErrorCode)
		return
	}
	g.send(c, response)
	return
}

const (
	Client  = 1
	Server  = 2
	NoRoute = 3
)

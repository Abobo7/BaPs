package pack

import (
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func AccountCurrencySync(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.AccountCurrencySyncResponse)

	rsp.AccountCurrencyDB = game.GetAccountCurrencyDB(s)
	rsp.ExpiredCurrency = make(map[proto.CurrencyTypes]int64)
}

func ItemList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.ItemListResponse)

	rsp.ExpiryItemDBs = make([]*proto.ItemDB, 0)
	rsp.ItemDBs = make([]*proto.ItemDB, 0)

	for _, conf := range game.GetItemList(s) {
		rsp.ItemDBs = append(rsp.ItemDBs, &proto.ItemDB{
			Type:       proto.ParcelType_Item,
			ServerId:   conf.ServerId,
			UniqueId:   conf.UniqueId,
			StackCount: conf.StackCount,
		})
	}
}

func EquipmentList(s *enter.Session, request, response mx.Message) {
	rsp := response.(*proto.EquipmentItemListResponse)

	rsp.EquipmentDBs = make([]*proto.EquipmentDB, 0)
}

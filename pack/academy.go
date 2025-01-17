package pack

import (
	"github.com/gucooing/BaPs/common/enter"
	"github.com/gucooing/BaPs/game"
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/proto"
)

func AcademyGetInfo(s *enter.Session, request, response mx.Message) {
	// req := request.(*proto.AcademyGetInfoRequest)
	rsp := response.(*proto.AcademyGetInfoResponse)

	rsp.AcademyDB = game.GetAcademyDB(s)
	rsp.AcademyLocationDBs = game.GetAcademyLocationDBs(s)
}

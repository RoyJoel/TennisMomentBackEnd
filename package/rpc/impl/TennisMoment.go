package impl

import (
	"context"
	"encoding/json"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/proto"
)

type PlayerInfoRPCImpl struct {
	dao *impl.PlayerDaoImpl
}

func NewPlayerControllerImpl() *PlayerInfoRPCImpl {
	return &PlayerInfoRPCImpl{dao: impl.NewPlayerDaoImpl()}
}
func (impl *PlayerInfoRPCImpl) AddNumByKey(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	loginName := request.GetLoginName()
	name := request.GetName()
	icon := request.GetIcon()
	sex := request.GetSex()
	age := request.GetAge()
	yearsPlayed := request.GetYearsPlayed()
	height := request.GetHeight()
	width := request.GetWidth()
	grip := request.GetGrip()
	backhand := request.GetBackhand()
	impl.dao.UpdatePlayerById(ctx, model.Player{
		id,
		loginName,
		name,
		icon,
		sex,
		age,
		yearsPlayed,
		height,
		width,
		grip,
		backhand,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) FindPlayerByKey(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	key := request.GetInfoKey()
// 	Player := impl.dao.GetPlayerByUid(ctx, key)
// 	info, _ := json.Marshal(Player)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
// }

func (impl *PlayerInfoRPCImpl) SavePlayerInfo(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	loginName := request.GetLoginName()
	name := request.GetName()
	icon := request.GetIcon()
	sex := request.GetSex()
	age := request.GetAge()
	yearsPlayed := request.GetYearsPlayed()
	height := request.GetHeight()
	width := request.GetWidth()
	grip := request.GetGrip()
	backhand := request.GetBackhand()

	impl.dao.CreatePlayer(ctx, model.Player{
		id,
		loginName,
		name,
		icon,
		sex,
		age,
		yearsPlayed,
		height,
		width,
		grip,
		backhand,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) DeleteById(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	id := request.GetId()
// 	impl.dao.DeletePlayerById(ctx, id)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
// }

func (impl *PlayerInfoRPCImpl) FindAll(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	Players := impl.dao.GetAll(ctx, 0, 0)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *PlayerInfoRPCImpl) GetPlayerInfoByUid(ctx context.Context, req *proto.PlayerInfoRequest) (resp *proto.PlayerInfoResponse, err error) {
	id := req.GetId()
	Player := impl.dao.GetPlayerInfoByUid(ctx, id)
	info, _ := json.Marshal(Player)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
}

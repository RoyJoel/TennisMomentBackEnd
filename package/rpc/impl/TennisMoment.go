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
func (impl *PlayerInfoRPCImpl) AddPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	YearsPlayed := request.GetYearsPlayed()
	Height := request.GetHeight()
	Width := request.GetWidth()
	Grip := request.GetGrip()
	Backhand := request.GetBackhand()
	Points := request.GetPoints()
	IsAdult := request.GetIsAdult()
	CareerStatsId := request.GetCareerStatsId()
	Friends := request.GetFriends()
	Relationship := request.GetRelationship()
	PlayerStats := request.GetPlayerStats()

	impl.dao.AddPlayer(ctx, model.Player{
		LoginName,
		Name,
		Icon,
		Sex,
		Age,
		YearsPlayed,
		Height,
		Width,
		Grip,
		Backhand,
		Points,
		IsAdult,
		CareerStatsId,
		Friends,
		PlayerStats,
		Relationship,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) FindPlayerByKey(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	key := request.GetInfoKey()
// 	Player := impl.dao.GetPlayerByUid(ctx, key)
// 	info, _ := json.Marshal(Player)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
// }

func (impl *PlayerInfoRPCImpl) UpdatePlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	YearsPlayed := request.GetYearsPlayed()
	Height := request.GetHeight()
	Width := request.GetWidth()
	Grip := request.GetGrip()
	Backhand := request.GetBackhand()
	Points := request.GetPoints()
	IsAdult := request.GetIsAdult()
	CareerStatsId := request.GetCareerStatsId()
	Friends := request.GetFriends()
	Relationship := request.GetRelationship()
	PlayerStats := request.GetPlayerStats()
	impl.dao.UpdatePlayer(ctx, model.Player{
		LoginName,
		Name,
		Icon,
		Sex,
		Age,
		YearsPlayed,
		Height,
		Width,
		Grip,
		Backhand,
		Points,
		IsAdult,
		CareerStatsId,
		Friends,
		PlayerStats,
		Relationship,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) DeleteById(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	id := request.GetId()
// 	impl.dao.DeletePlayerById(ctx, id)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
// }

func (impl *PlayerInfoRPCImpl) SearchPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	loginName := request.GetLoginName()
	Players := impl.dao.GetPlayerInfo(ctx, loginName)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil

}

func (impl *PlayerInfoRPCImpl) GetPlayerInfo(ctx context.Context, req *proto.PlayerInfoRequest) (resp *proto.PlayerInfoResponse, err error) {
	loginName := req.GetLoginName()
	Players := impl.dao.GetPlayerInfo(ctx, loginName)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil

}

func (impl *PlayerInfoRPCImpl) AddFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	loginName := request.GetLoginName()
	Players := impl.dao.GetPlayerInfo(ctx, loginName)

	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil

}
func (impl *PlayerInfoRPCImpl) DeleteFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	loginName := request.GetLoginName()
	Players := impl.dao.GetPlayerInfo(ctx, loginName)

	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil

}

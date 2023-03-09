package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/RoyJoel/TennisMomentBackEnd/package/cache"
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
)

func TestCreate(t *testing.T) {
	daoImpl := impl.NewPlayerDaoImpl()
	info := model.Player{LoginName: "Jason Zhang", Name: "Jason Zhang", Icon: "make([]byte, 16)", Sex: "man", Age: 21, YearsPlayed: 2, Height: 170.5, Width: 88.3, Grip: "western", Backhand: "two-handed backhand"}

	daoImpl.AddPlayer(context.Background(), info)
	fmt.Println(config.DB)
}

func IntToBytes(i int) {
	panic("unimplemented")
}

func TestFindById(t *testing.T) {
	daoImpl := impl.NewPlayerDaoImpl()
	key := daoImpl.GetPlayerInfo(context.Background(), "Jason Zhang")
	fmt.Println(key)
}

// func TestFindAll(t *testing.T) {
// 	daoImpl := impl.NewPlayerDaoImpl()
// 	key := daoImpl.FindAllPlayers(context.Background(), 1, 1)
// 	fmt.Println(key)
// }

func TestUpdate(t *testing.T) {
	daoImpl := impl.NewPlayerDaoImpl()
	daoImpl.UpdatePlayer(context.Background(), model.Player{LoginName: "Jason Zhang", Name: "Jason Zhang", Icon: "make([]byte, 16)", Sex: "man", Age: 21, YearsPlayed: 2, Height: 170.5, Width: 88.3, Grip: "western", Backhand: "two-handed backhand"})
}

// func TestDelete(t *testing.T) {
// 	impl.NewPlayerDaoImpl().DeletePlayerInfoById(context.Background(), 1)
// }

func TestCache(t *testing.T) {
	cacheDAOImpl := cache.NewPlayerCacheDAOImpl()
	cacheDAOImpl.SetPlayer(context.Background(), "1", model.Player{LoginName: "Jason Zhang", Name: "Jason Zhang", Icon: "make([]byte, 16)", Sex: "man", Age: 21, YearsPlayed: 2, Height: 170.5, Width: 88.3, Grip: "western", Backhand: "two-handed backhand"}, time.Second*1100)
	cacheDAOImpl.GetPlayer(context.Background(), "1")
}

// func TestFindByPage(t *testing.T) {
// 	allPlayerInfo := impl.NewPlayerDaoImpl().FindAllPlayers(context.Background(), 2, 10)
// 	fmt.Println(allPlayerInfo)
// }

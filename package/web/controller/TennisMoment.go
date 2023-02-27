package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type StatsInfoControllerImpl struct {
	dao *impl.StatsDaoImpl
}

type StatsInfoController interface {
	// AddNumByKey(c *gin.Context)
	// FindNumByKey(c *gin.Context)
	SaveStatsInfo(c *gin.Context)
	// DeleteById(c *gin.Context)
	FindAll(c *gin.Context)
	FindStatsById(c *gin.Context)
	Update(context *gin.Context)
}

func NewStatsInfoControllerImpl() *StatsInfoControllerImpl {
	return &StatsInfoControllerImpl{dao: impl.NewStatsDaoImpl()}
}

// func (impl StatsInfoControllerImpl) AddNumByKey(c *gin.Context) {
// 	key := c.Param("key")
// 	statsInfo := impl.dao.GetStatsInfoByKey(c, key)
// 	statsInfo.InfoNum = statsInfo.InfoNum + 1
// 	isOk := impl.dao.UpdateStatsInfoByKey(c, statsInfo)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
// }

// func (impl StatsInfoControllerImpl) FindNumByKey(c *gin.Context) {
// 	key := c.Param("key")
// 	statsInfo := impl.dao.GetStatsInfoByKey(c, key)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": statsInfo})
// }

func (impl StatsInfoControllerImpl) SaveStatsInfo(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	info := model.Stats{}
	json.Unmarshal(bytes, &info)
	if err != nil {
		panic(err)
	}
	isOk := impl.dao.CreateStatsInfo(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

// func (impl StatsInfoControllerImpl) DeleteById(c *gin.Context) {
// 	id := c.Param("id")
// 	i, _ := strconv.Atoi(id)
// 	isOk := impl.dao.DeleteStatsInfoById(c, int64(i))
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
// }

func (impl StatsInfoControllerImpl) FindAll(c *gin.Context) {
	a := model.Stats{Aces: 1, DoubleFaults: 1, FirstServePoints: 1, FirstServePointsIn: 1, FirstServePointsWon: 1, SecondServePoints: 1, SecondServePointsWon: 1, BreakPointsFaced: 1, BreakPointsSaved: 1, ServeGamesPlayed: 1, ServeGamesWon: 1, TotalServePointsWon: 1, FirstServeReturnPoints: 1, FirstServeReturnAces: 1, FirstServeReturnPointsWon: 1, SecondServeReturnPoints: 1, SecondServeReturnAces: 1, SecondServeReturnPointsWon: 1, BreakPointsOpportunities: 1, BreakPointsConverted: 1, ReturnGamesPlayed: 1, ReturnGamesWon: 1, TotalReturnPointsWon: 1, TotalPointsWon: 1, NetPoints: 1, UnforcedErrors: 1, ForehandWinners: 1, BackhandWinners: 1, PlayerId: 1, Player: model.Player{Id: 123, LoginName: "Jason Zhang", Name: "Jason Zhang", Icon: make([]byte, 16), Sex: "man", Age: 21, YearsPlayed: 2, Height: 170.5, Width: 88.3, Grip: "western", Backhand: "two-handed backhand"}}
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := a
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}

	fmt.Println(player)
	im := impl.dao.CreateStatsInfo(c, player)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": im})

	page := c.Query("page")
	limit := c.Query("limit")
	statsInfos := impl.dao.FindAllStats(c, cast.ToInt(page), cast.ToInt(limit))

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(statsInfos), "data": statsInfos})
}

func (impl StatsInfoControllerImpl) FindStatsById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	statsInfo := impl.dao.GetStatsInfoByUid(c, int64(i))
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": statsInfo})
}

func (impl StatsInfoControllerImpl) Update(c *gin.Context) {
	body := c.Request.Body
	jsonBytes, err := ioutil.ReadAll(body)
	d := json.NewDecoder(bytes.NewReader(jsonBytes))
	d.UseNumber()
	m := make(map[string]string)
	d.Decode(&m)
	if err != nil {
		panic(err)
	}
	info := model.Stats{

		Aces:                       cast.ToInt64(m["aces"]),
		DoubleFaults:               cast.ToInt64(m["doubleFaults"]),
		FirstServePoints:           cast.ToInt64(m["firstServePoints"]),
		FirstServePointsIn:         cast.ToInt64(m["firstServePointsIn"]),
		FirstServePointsWon:        cast.ToInt64(m["firstServePointsWon"]),
		SecondServePoints:          cast.ToInt64(m["secondServePoints"]),
		SecondServePointsWon:       cast.ToInt64(m["secondServePointsWon"]),
		BreakPointsFaced:           cast.ToInt64(m["breakPointsFaced"]),
		BreakPointsSaved:           cast.ToInt64(m["breakPointsSaved"]),
		ServeGamesPlayed:           cast.ToInt64(m["serveGamesPlayed"]),
		ServeGamesWon:              cast.ToInt64(m["serveGamesWon"]),
		TotalServePointsWon:        cast.ToInt64(m["totalServePointsWon"]),
		FirstServeReturnPoints:     cast.ToInt64(m["firstServeReturnPoints"]),
		FirstServeReturnAces:       cast.ToInt64(m["firstServeReturnAces"]),
		FirstServeReturnPointsWon:  cast.ToInt64(m["firstServeReturnPointsWon"]),
		SecondServeReturnPoints:    cast.ToInt64(m["secondServeReturnPoints"]),
		SecondServeReturnAces:      cast.ToInt64(m["secondServeReturnAces"]),
		SecondServeReturnPointsWon: cast.ToInt64(m["secondServeReturnPointsWon"]),
		BreakPointsOpportunities:   cast.ToInt64(m["breakPointsOpportunities"]),
		BreakPointsConverted:       cast.ToInt64(m["breakPointsConverted"]),
		ReturnGamesPlayed:          cast.ToInt64(m["returnGamesPlayed"]),
		ReturnGamesWon:             cast.ToInt64(m["returnGamesWon"]),
		TotalReturnPointsWon:       cast.ToInt64(m["totalReturnPointsWon"]),
		TotalPointsWon:             cast.ToInt64(m["totalPointsWon"]),
		NetPoints:                  cast.ToInt64(m["netPoints"]),
		UnforcedErrors:             cast.ToInt64(m["unforcedErrors"]),
		ForehandWinners:            cast.ToInt64(m["forehandWinners"]),
		BackhandWinners:            cast.ToInt64(m["backhandWinners"]),
	}
	isOk := impl.dao.UpdateStatsInfoById(c, info)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
}

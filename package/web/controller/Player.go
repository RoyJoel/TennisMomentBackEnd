package controller

import (
	"encoding/json"
	"net/http"

	// "fmt"
	"io/ioutil"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/cast"
	// "github.com/RoyJoel/TennisMomentBackEnd/package/utils"
	// "github.com/RoyJoel/TennisMomentBackEnd/package/web/auth"
)

type PlayerControllerImpl struct {
	dao *impl.PlayerDaoImpl
}

type PlayerController interface {
	AddPlayer(c *gin.Context)
	SearchPlayer(c *gin.Context)
	UpdatePlayer(c *gin.Context)
	GetPlayerInfo(c *gin.Context)
	AddFriend(c *gin.Context)
	DeleteFriend(c *gin.Context)
}

func NewPlayerControllerImpl() *PlayerControllerImpl {
	return &PlayerControllerImpl{dao: impl.NewPlayerDaoImpl()}
}

func (impl PlayerControllerImpl) AddPlayer(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	res := impl.dao.AddPlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl PlayerControllerImpl) SearchPlayer(c *gin.Context) {

	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetPlayerInfo(c, req.LoginName)
	if player.LoginName != "" {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
	}
}

func (impl PlayerControllerImpl) UpdatePlayer(c *gin.Context) {

	// LoginName := c.PostForm("loginName")
	// name := c.PostForm("name")
	// icon := []byte(c.PostForm("icon"))
	// sex := c.PostForm("sex")
	// age, _ := strconv.ParseInt(c.PostForm("age"), 10, 64)
	// yearsPlayed, _ := strconv.ParseInt(c.PostForm("yearsPlayed"), 10, 64)
	// height, _ := strconv.ParseFloat(c.PostForm("height"), 32)
	// width, _ := strconv.ParseFloat(c.PostForm("width"), 32)
	// grip := c.PostForm("grip")
	// backhand := c.PostForm("backhand")
	// careerStatsId, _ := strconv.Atoi(c.PostForm("careerStatsId"))

	// Requestbody
	// data, _ := json.Marshal()
	// fmt.Println(nameArray)
	// var byteSlice []byte
	// err := json.Unmarshal(data, &byteSlice)
	// if err != nil {
	// 	panic(err)
	// }

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}

	player = impl.dao.UpdatePlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl PlayerControllerImpl) GetPlayerInfo(c *gin.Context) {
	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetPlayerInfo(c, req.LoginName)
	if player.LoginName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get player info"})
	} else {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
	}
	//密码通过
	// if playerByLoginName.Pwd == utils.GetMd5Str(player.Pwd) {
	// setToken := auth.SetToken(c, utils.GetTokenStr(), player)
	// } else {
	// 	if playerByLoginName.Id == 0 {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
	// 	} else {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
	// 	}
	// }
}

func (impl PlayerControllerImpl) AddFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}
	result := impl.dao.AddFriend(c, relationship)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
	//密码通过
	// if playerByLoginName.Pwd == utils.GetMd5Str(player.Pwd) {
	// setToken := auth.SetToken(c, utils.GetTokenStr(), player)
	// } else {
	// 	if playerByLoginName.Id == 0 {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
	// 	} else {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
	// 	}
	// }
}

func (impl PlayerControllerImpl) DeleteFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}
	result := impl.dao.DeleteFriend(c, relationship)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

func (impl PlayerControllerImpl) GetAllFriends(c *gin.Context) {
	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetAllFriends(c, req.LoginName)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl PlayerControllerImpl) SearchFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}

	res := impl.dao.SearchFriend(c, relationship)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

// func (impl PlayerControllerImpl) addPlayer(c *gin.Context) {
// 	body := c.Request.Body
// 	bytes, err := ioutil.ReadAll(body)
// 	player := model.Player{}
// 	json.Unmarshal(bytes, &player)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res := impl.dao.addPlayer(c, player)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
// }

// func (impl PlayerControllerImpl) FindAllPlayers(c *gin.Context) {
// 	page := c.Query("page")
// 	limit := c.Query("limit")
// 	players := impl.dao.FindAllPlayers(c, cast.ToInt(page), cast.ToInt(limit))
// 	fmt.Println(players)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(players), "data": players})
// }

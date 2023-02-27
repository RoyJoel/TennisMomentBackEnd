package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
	"github.com/RoyJoel/TennisMomentBackEnd/package/web/auth"
	"github.com/gin-gonic/gin"
)

type PlayerControllerImpl struct {
	dao *impl.PlayerDaoImpl
}

type PlayerController interface {
	CreatePlayer(c *gin.Context)
	FindPlayerByLoginNameAndPwd(c *gin.Context)
	Register(c *gin.Context)
}

func NewPlayerController() *PlayerControllerImpl {
	return &PlayerControllerImpl{dao: impl.NewPlayerDaoImpl()}
}

func (impl PlayerControllerImpl) CreatePlayer(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}
	res := impl.dao.CreatePlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl PlayerControllerImpl) FindPlayerByLoginNameAndPwd(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}
	// playerByLoginName := impl.dao.GetPlayerByLoginName(c, player.LoginName)
	//密码通过
	// if playerByLoginName.Pwd == utils.GetMd5Str(player.Pwd) {
	setToken := auth.SetToken(c, utils.GetTokenStr(), player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": setToken, "count": 0, "data": utils.GetTokenStr()})
	// } else {
	// 	if playerByLoginName.Id == 0 {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
	// 	} else {
	// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
	// 	}
	// }
}

func (impl PlayerControllerImpl) Register(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	res := impl.dao.CreatePlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

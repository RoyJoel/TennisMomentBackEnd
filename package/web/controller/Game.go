package controller

import (
	"encoding/json"

	// "fmt"
	"io/ioutil"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
	"github.com/gin-gonic/gin"
	// "github.com/spf13/cast"
	// "github.com/RoyJoel/TennisMomentBackEnd/package/utils"
	// "github.com/RoyJoel/TennisMomentBackEnd/package/web/auth"
)

type GameControllerImpl struct {
	dao *impl.GameDaoImpl
}

type GameController interface {
	AddGame(c *gin.Context)
	SearchGame(c *gin.Context)
	UpdateGame(c *gin.Context)
}

func NewGameControllerImpl() *GameControllerImpl {
	return &GameControllerImpl{dao: impl.NewGameDaoImpl()}
}

func (impl GameControllerImpl) AddGame(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	Game := model.Game{}
	json.Unmarshal(bytes, &Game)
	if err != nil {
		panic(err)
	}
	// Game.Role = 1
	res := impl.dao.AddGame(c, &Game)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl GameControllerImpl) SearchAllGames(c *gin.Context) {

	type SearchGameRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetAllGames(c, req.LoginName)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl GameControllerImpl) UpdateGame(c *gin.Context) {
	type SearchGameRequest struct {
		Date                float64         `json:"date"`
		Place               string          `json:"place"`
		Surface             string          `json:"surface"`
		IsPlayer1Serving    bool            `json:"isPlayer1Serving"`
		IsCompleted         bool            `json:"isCompleted"`
		Player1LoginName    string          `json:"player1LoginName"`
		Player1StatsId      int             `json: "player1StatsId"`
		Player2LoginName    string          `json:"player2LoginName"`
		IsPlayer1FirstServe bool            `json:"isPlayer1FirstServe"`
		IsPlayer2FirstServe bool            `json:"isPlayer2FirstServe"`
		Result              utils.IntMatrix `json: "result"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	game := impl.dao.UpdateGame(c, req.Date, req.Place, req.Surface, req.IsCompleted, req.Player1LoginName, req.Player1StatsId, req.Player2LoginName, req.Result)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": game})
}

func (impl GameControllerImpl) SearchGame(c *gin.Context) {
	type SearchGameRequest struct {
		Player1 string `json:"player1LoginName"`
		Player2 string `json:"player2LoginName"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	game := impl.dao.SearchGame(c, req.Player1, req.Player2)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": game})
}

// func (impl GameControllerImpl) GetGameInfo(c *gin.Context) {
// 	body := c.Request.Body
// 	bytes, err := ioutil.ReadAll(body)
// 	GameLoginName := model.Game{}.LoginName
// 	json.Unmarshal(bytes, &GameLoginName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	Game := impl.dao.GetGameInfo(c, GameLoginName)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": Game})
//密码通过
// if GameByLoginName.Pwd == utils.GetMd5Str(Game.Pwd) {
// setToken := auth.SetToken(c, utils.GetTokenStr(), Game)
// } else {
// 	if GameByLoginName.Id == 0 {
// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
// 	} else {
// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
// 	}
// }
// }

// func (impl GameControllerImpl) AddFriend(c *gin.Context) {
// 	body := c.Request.Body
// 	bytes, err := ioutil.ReadAll(body)
// 	GameLoginName := model.Game{}.LoginName
// 	json.Unmarshal(bytes, &GameLoginName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	Game := impl.dao.GetGameInfo(c, GameLoginName)

// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": Game})
//密码通过
// if GameByLoginName.Pwd == utils.GetMd5Str(Game.Pwd) {
// setToken := auth.SetToken(c, utils.GetTokenStr(), Game)
// } else {
// 	if GameByLoginName.Id == 0 {
// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "账号不存在", "count": 0, "data": "-1"})
// 	} else {
// 		c.JSON(200, map[string]interface{}{"code": 0, "msg": "密码错误", "count": 0, "data": "-1"})
// 	}
// }
// }

// func (impl GameControllerImpl) DeleteFriend(c *gin.Context) {

// }

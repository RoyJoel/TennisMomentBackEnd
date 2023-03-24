package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/gin-gonic/gin"
)

type TennisMomentControllerImpl struct {
	dao *impl.TennisMomentDaoImpl
}

func NewTennisMomentControllerImpl() *TennisMomentControllerImpl {
	return &TennisMomentControllerImpl{dao: impl.NewTennisMomentDaoImpl()}
}

func (impl TennisMomentControllerImpl) AddGame(c *gin.Context) {
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

func (impl TennisMomentControllerImpl) SearchRecentGames(c *gin.Context) {

	type SearchGameRequest struct {
		Id    int64 `json:"id"`
		Limit int   `json:"limit"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	Games := impl.dao.GetRecentGames(c, req.Id, req.Limit)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": Games})
}

func (impl TennisMomentControllerImpl) SearchAllGames(c *gin.Context) {

	type SearchGameRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetAllGames(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl TennisMomentControllerImpl) UpdateGameAndStats(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	Game := model.GameResponse{}
	json.Unmarshal(bytes, &Game)
	if err != nil {
		panic(err)
	}

	game := impl.dao.UpdateGameAndStats(c, Game)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": game})
}

func (impl TennisMomentControllerImpl) GetHistoryGames(c *gin.Context) {
	type SearchGameRequest struct {
		Player1Id int64 `json:"player1Id"`
		Player2Id int64 `json:"player2Id"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	game := impl.dao.GetHistoryGames(c, req.Player1Id, req.Player2Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": game})
}

func (impl TennisMomentControllerImpl) AddPlayer(c *gin.Context) {
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

func (impl TennisMomentControllerImpl) SignUp(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	user := model.User{}
	json.Unmarshal(bytes, &user)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	res := impl.dao.SignUp(c, user)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl TennisMomentControllerImpl) SignIn(c *gin.Context) {

	type SignInRequest struct {
		Id int64 `json:"id"`
	}

	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	user := impl.dao.SignIn(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": user})

}

func (impl TennisMomentControllerImpl) SearchPlayer(c *gin.Context) {

	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetPlayerInfo(c, req.Id)
	if player.LoginName != "" {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
	}
}

func (impl TennisMomentControllerImpl) UpdatePlayer(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.PlayerResponse{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}

	player = impl.dao.UpdatePlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl TennisMomentControllerImpl) GetPlayerInfo(c *gin.Context) {
	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetPlayerInfo(c, req.Id)
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

func (impl TennisMomentControllerImpl) AddFriend(c *gin.Context) {
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

func (impl TennisMomentControllerImpl) DeleteFriend(c *gin.Context) {
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

func (impl TennisMomentControllerImpl) GetAllFriends(c *gin.Context) {
	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetAllFriends(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl TennisMomentControllerImpl) SearchFriend(c *gin.Context) {
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

func (impl TennisMomentControllerImpl) SearchStats(c *gin.Context) {
	type SearchStatsRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchStatsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	stats := impl.dao.GetStatsInfo(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": stats})
}

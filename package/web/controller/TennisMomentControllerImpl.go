package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/middleware"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/RoyJoel/TennisMomentBackEnd/package/utils"
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
		Id          int64 `json:"id"`
		Limit       int   `json:"limit"`
		IsCompleted bool  `json:"isCompleted"`
	}

	var req SearchGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	Games := impl.dao.GetRecentGames(c, req.Id, req.Limit, req.IsCompleted)

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
	res, _ := impl.dao.AddPlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl TennisMomentControllerImpl) SignUp(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	User := model.User{}
	json.Unmarshal(bytes, &User)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	user, res := impl.dao.SignUp(c, User)
	type addResponse struct {
		User model.User `json:"user"`
		Res  bool       `json:"res"`
	}
	result := addResponse{User: user, Res: res}
	fmt.Println(result)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

func (impl TennisMomentControllerImpl) UpdateUser(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	User := model.User{}
	json.Unmarshal(bytes, &User)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	user := impl.dao.UpdateUser(c, User)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": user})
}

func (impl TennisMomentControllerImpl) Auth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claims, error := middleware.ParseToken(auth)
	type authResponse struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}
	res := authResponse{LoginName: claims.LoginName, Password: claims.Password}

	if error != nil {
		c.JSON(401, map[string]interface{}{"code": 0, "msg": error, "count": 0, "data": nil})
	} else {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
	}
}

func (impl TennisMomentControllerImpl) SignIn(c *gin.Context) {

	type SignInRequest struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}

	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	user, error := impl.dao.SignIn(c, req.LoginName, req.Password)
	if error != nil {
		c.JSON(401, map[string]interface{}{"code": 0, "msg": error, "count": 0, "data": nil})
	}
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": user})

}

func (impl TennisMomentControllerImpl) ResetPassword(c *gin.Context) {

	type resetRequest struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}

	var req resetRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.ResetPassword(c, req.LoginName, req.Password)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})

}

func (impl TennisMomentControllerImpl) SearchPlayer(c *gin.Context) {

	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.SearchPlayer(c, req.LoginName)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})

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
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player, _ := impl.dao.GetPlayerInfoByLoginName(c, req.LoginName)
	if player.LoginName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get player info"})
	} else {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
	}
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

func (impl TennisMomentControllerImpl) AddSchedule(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	schedule := model.Schedule{}
	json.Unmarshal(bytes, &schedule)
	if err != nil {
		panic(err)
	}
	result := impl.dao.AddSchedule(c, schedule)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
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

func (impl TennisMomentControllerImpl) GetClubInfos(c *gin.Context) {
	type SearchPlayerRequest struct {
		Ids utils.IntMatrix `json:"ids"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}
	res := make([]model.ClubResponse, 0)

	for _, id := range req.Ids {
		club := impl.dao.GetClubInfo(c, id)
		res = append(res, club)
	}
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

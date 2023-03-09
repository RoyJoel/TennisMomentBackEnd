package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/RoyJoel/TennisMomentBackEnd/package/dao/impl"
	"github.com/RoyJoel/TennisMomentBackEnd/package/model"
	"github.com/gin-gonic/gin"
)

type StatsControllerImpl struct {
	dao *impl.StatsDaoImpl
}

type StatsController interface {
	// AddNumByKey(c *gin.Context)
	// FindNumByKey(c *gin.Context)
	SaveStatsInfo(c *gin.Context)
	// DeleteById(c *gin.Context)
	FindAll(c *gin.Context)
	FindStatsById(c *gin.Context)
	Update(context *gin.Context)
}

func NewStatsControllerImpl() *StatsControllerImpl {
	return &StatsControllerImpl{dao: impl.NewStatsDaoImpl()}
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

// func (impl StatsControllerImpl) GetStatsInfoByPlayerLoginName(c *gin.Context) {
// 	type SearchPlayerRequest struct {
// 		LoginName string `json:"loginName"`
// 	}

// 	var req SearchPlayerRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		// 处理错误
// 	}

// 	player := impl.dao.GetStatsInfoByPlayerLoginName(c, req.LoginName)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
// }

// func (impl StatsInfoControllerImpl) DeleteById(c *gin.Context) {
// 	id := c.Param("id")
// 	i, _ := strconv.Atoi(id)
// 	isOk := impl.dao.DeleteStatsInfoById(c, int64(i))
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": isOk})
// }

// func (impl StatsControllerImpl) CreateStatsByPlayerLoginName(c *gin.Context) {
// 	type SearchStatsRequest struct {
// 		LoginName string `json:"loginName"`
// 	}

// 	var req SearchStatsRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		// 处理错误
// 	}

// 	player := impl.dao.CreateStatsByPlayerLoginName(c, req.LoginName)
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
// }

func (impl StatsControllerImpl) SearchStats(c *gin.Context) {
	type SearchStatsRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchStatsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	stats := impl.dao.SearchStatsInfo(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": stats})
}

// func (impl StatsControllerImpl) FindStatsById(c *gin.Context) {
// 	id := c.Param("id")
// 	i, _ := strconv.Atoi(id)
// 	statsInfo := impl.dao.GetStatsInfoByUid(c, int64(i))
// 	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": statsInfo})
// }

func (impl StatsControllerImpl) UpdateStats(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	Stats := model.Stats{}
	json.Unmarshal(bytes, &Stats)
	if err != nil {
		panic(err)
	}
	result := impl.dao.UpdateStatsInfoById(c, Stats)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

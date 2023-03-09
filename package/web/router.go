package web

import (
	"github.com/RoyJoel/TennisMomentBackEnd/package/config"
	"github.com/RoyJoel/TennisMomentBackEnd/package/web/controller"
	"github.com/RoyJoel/TennisMomentBackEnd/package/web/interceptor"

	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	//增加拦截器
	r.Use(interceptor.HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())
	//路由组
	playerInfo := r.Group("/player")
	{
		playerInfo.POST("/update", controller.NewPlayerControllerImpl().UpdatePlayer)
		playerInfo.POST("/getInfo", controller.NewPlayerControllerImpl().GetPlayerInfo)
		playerInfo.POST("/add", controller.NewPlayerControllerImpl().AddPlayer)
		playerInfo.POST("/search", controller.NewPlayerControllerImpl().SearchPlayer)
	}

	friendInfo := r.Group("/friend")
	{
		friendInfo.POST("/add", controller.NewPlayerControllerImpl().AddFriend)
		friendInfo.POST("/search", controller.NewPlayerControllerImpl().SearchFriend)
		friendInfo.POST("/delete", controller.NewPlayerControllerImpl().DeleteFriend)
		friendInfo.POST("/getAll", controller.NewPlayerControllerImpl().GetAllFriends)
	}

	gameInfo := r.Group("/game")
	{
		gameInfo.POST("/update", controller.NewGameControllerImpl().UpdateGame)
		gameInfo.POST("/add", controller.NewGameControllerImpl().AddGame)
		gameInfo.POST("/searchAll", controller.NewGameControllerImpl().SearchAllGames)
		gameInfo.POST("/search", controller.NewGameControllerImpl().SearchGame)
	}

	statsInfo := r.Group("/stats")
	{
		statsInfo.POST("/search", controller.NewStatsControllerImpl().SearchStats)
		statsInfo.POST("/update", controller.NewStatsControllerImpl().UpdateStats)
	}

	r.Run(":8080")
}

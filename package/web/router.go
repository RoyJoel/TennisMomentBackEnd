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

	authInfo := r.Group("/")
	{
		authInfo.GET("auth", controller.NewTennisMomentControllerImpl().Auth)
	}

	//路由组
	userInfo := r.Group("/user")
	{
		userInfo.POST("/signIn", controller.NewTennisMomentControllerImpl().SignIn)
		userInfo.POST("/signUp", controller.NewTennisMomentControllerImpl().SignUp)
		userInfo.POST("/resetPassword", controller.NewTennisMomentControllerImpl().ResetPassword)
		userInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateUser)
	}

	scheduleInfo := r.Group("/schedule")
	{
		scheduleInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddSchedule)
	}
	playerInfo := r.Group("/player")
	{
		playerInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdatePlayer)
		playerInfo.POST("/getInfo", controller.NewTennisMomentControllerImpl().GetPlayerInfo)
		playerInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddPlayer)
		playerInfo.POST("/search", controller.NewTennisMomentControllerImpl().SearchPlayer)
	}

	friendInfo := r.Group("/friend")
	{
		friendInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddFriend)
		friendInfo.POST("/search", controller.NewTennisMomentControllerImpl().SearchFriend)
		friendInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteFriend)
		friendInfo.POST("/getAll", controller.NewTennisMomentControllerImpl().GetAllFriends)
	}

	gameInfo := r.Group("/game")
	{
		gameInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateGameAndStats)
		gameInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddGame)
		gameInfo.POST("/searchAll", controller.NewTennisMomentControllerImpl().SearchAllGames)
		gameInfo.POST("/h2h", controller.NewTennisMomentControllerImpl().GetHistoryGames)
		gameInfo.POST("/searchRecent", controller.NewTennisMomentControllerImpl().SearchRecentGames)
	}

	statsInfo := r.Group("/stats")
	{
		statsInfo.POST("/search", controller.NewTennisMomentControllerImpl().SearchStats)
	}

	clubInfo := r.Group("/club")
	{
		clubInfo.POST("/getInfos", controller.NewTennisMomentControllerImpl().GetClubInfos)
	}

	r.Run(":8080")
}

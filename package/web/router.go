package web

import (
	"fmt"

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
	stats := r.Group("/Stats")
	{
		// stats.POST("/add/:key", controller.NewStatsInfoControllerImpl().AddStatsByKey)
		// stats.GET("/findByKey/:key", controller.NewStatsInfoControllerImpl().FindStatsByKey)
		stats.GET("/findById/:id", controller.NewStatsInfoControllerImpl().FindStatsById)
		stats.POST("/saveInfo", controller.NewStatsInfoControllerImpl().SaveStatsInfo)
		// stats.POST("/deleteInfo/:id", controller.NewStatsInfoControllerImpl().DeleteById)
		stats.GET("/getAll", controller.NewStatsInfoControllerImpl().FindAll)
		fmt.Println("fdsafsafasdff")
		stats.POST("/update", controller.NewStatsInfoControllerImpl().Update)
	}

	// roleInfo := r.Group("/role")
	// {
	// 	roleInfo.POST("/save", controller.NewRoleControllerImpl().CreateRole)
	// 	roleInfo.GET("/all", controller.NewRoleControllerImpl().GetAll)
	// }

	playerInfo := r.Group("/player")
	{
		playerInfo.POST("/save", controller.NewPlayerController().CreatePlayer)
		playerInfo.POST("/login", controller.NewPlayerController().FindPlayerByLoginNameAndPwd)
		playerInfo.POST("/register", controller.NewPlayerController().Register)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/add", controller.AddPolicy)
	}

	r.Run("0.0.0.0:" + config.PORT)
}

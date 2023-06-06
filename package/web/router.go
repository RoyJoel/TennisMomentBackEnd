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

	orderInfo := r.Group("/order")
	{
		orderInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddOrder)
		orderInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteOrder)
		// orderInfo.POST("/search", controller.NewTennisMomentControllerImpl().GetOrderInfos)
		orderInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateOrder)
		orderInfo.POST("/getInfos", controller.NewTennisMomentControllerImpl().GetOrderInfosByUserId)
		orderInfo.GET("/getAll", controller.NewTennisMomentControllerImpl().GetAllOrders)
	}

	addressInfo := r.Group("/address")
	{
		addressInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddAddress)
		addressInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteAddress)
		addressInfo.POST("/getInfos", controller.NewTennisMomentControllerImpl().GetAddressInfos)
		addressInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateAddress)
	}

	cartInfo := r.Group("/cart")
	{
		cartInfo.POST("/getInfo", controller.NewTennisMomentControllerImpl().GetCartInfo)
		cartInfo.POST("/addBill", controller.NewTennisMomentControllerImpl().AddBillToCart)
		cartInfo.POST("/deleteBill", controller.NewTennisMomentControllerImpl().DeleteBillInCart)
		cartInfo.POST("/assign", controller.NewTennisMomentControllerImpl().AssignCartForUser)
	}

	commodityInfo := r.Group("/commodity")
	{
		commodityInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddCommodity)
		commodityInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteCommodity)
		// commodityInfo.POST("/search", controller.NewTennisMomentControllerImpl().getCommodityInfo)
		commodityInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateCommodity)
		commodityInfo.GET("/getAll", controller.NewTennisMomentControllerImpl().GetAllCommodities)
	}

	optionInfo := r.Group("/option")
	{
		optionInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddOption)
		optionInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateOption)
		optionInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteOption)

	}

	// billInfo := r.Group("/bill")
	// {
	// billInfo.POST("/add", controller.NewTennisMomentControllerImpl().AddBill)
	// billInfo.POST("/delete", controller.NewTennisMomentControllerImpl().DeleteBill)
	// billInfo.POST("/search", controller.NewTennisMomentControllerImpl().getBillInfo)
	// billInfo.POST("/update", controller.NewTennisMomentControllerImpl().UpdateBill)
	// }

	r.Run(":8080")
}

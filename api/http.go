package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HttpInitEngine() {
	// 用户接口
	r := gin.Default()
	r.Use(cors())

	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.POST("/login", Login)
		user.DELETE("/state/delete", QuitLogin)
		user.GET("/state", State)
		user.GET("/balance", Balance)
		user.PUT("/recharge", Recharge)
		user.GET("/message", GetMessage)
		user.POST("/message/change", ChangeMessage)
		user.GET("/addr", GetAddr)
		user.PUT("/addr/add", AddAddr)
		user.DELETE("/addr/delete", DeleteAddr)
		user.GET("/trolley", GetTrolley)
		user.POST("/trolley/buy", BuyTrolleyGood)
		user.DELETE("/trolley/delete", DeleteTrolleyGood)
		user.GET("/buy/addr", GetAddr)
		user.PUT("/buy/addr/add", AddAddr)
		user.POST("/buy/addr/choose", ChooseAddr)
		user.POST("/buy/sure", SureBuy)
		user.POST("/buy/return", ReturnBuy)
	}

	good := r.Group("/good")
	{
		good.GET("/", Good)
		good.POST("/find", FindGood)
		good.POST("/classify", GoodClassify)
		good.POST("/message", GoodMessage)
		good.GET("/comment", GetGoodComment)
		good.POST("/comment/publish", PublishComment)
		good.DELETE("/res/delete", DeleteRes)
		good.POST("/comment/respond", PublishRespond)
		good.DELETE("/comment/delete", DeleteComment)
		good.POST("/trolley/add", AddTrolley)
		good.POST("/collect/add", CollectGood)
		good.DELETE("/collect/delete", DeleteCollect)
		good.GET("/collect", GetCollect)
	}

	shop := r.Group("/shop")
	{
		shop.GET("/", Shop)
		shop.POST("/classify", ClassifyGood)
		shop.PUT("/good/add", AddGood)
		shop.DELETE("/good/delete", DeleteGood)
		shop.GET("/balance", GetShopBalance)
		shop.POST("/register", RegisterShop)
	}

	order := r.Group("/order")
	{
		order.GET("/sort", OrderSort)
		order.DELETE("/delete", DeleteOrder)
		order.POST("/arrive/sure", SureArrive)
		order.POST("/comment", CommentOrder)
	}

	err := r.Run(":8086")
	if err != nil {
		fmt.Println(err)
		return
	}
}

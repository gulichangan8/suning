package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"suning/respond"
	"suning/service"
)

func Shop(c *gin.Context) {
	shop := c.PostForm("shop_name")
	s, G := service.GetShopMessage(shop)
	respond.GetShopTrue(c, s, G)
}

func ClassifyGood(c *gin.Context) {
	classify := c.PostForm("classify")
	shop := c.PostForm("shop_name")
	G := service.Classify(classify, shop)
	respond.GoodTrue(c, G)
}

func AddGood(c *gin.Context) {
	shop := c.PostForm("shop_name")
	good := c.PostForm("good_name")
	image := c.PostForm("image")
	r := c.PostForm("range")
	color := c.PostForm("color")
	serve := c.PostForm("serve")
	message := c.PostForm("message")
	sellMessage := c.PostForm("sell_message")
	classify := c.PostForm("classify")
	money, _ := strconv.Atoi(c.PostForm("money"))
	service.AddGood(shop, good, money, image, r, color, serve, message, sellMessage, classify)
	respond.CollectGoodTure(c)
}

func DeleteGood(c *gin.Context) {
	shop := c.PostForm("shop_name")
	good := c.PostForm("good_name")
	service.DeleteGood(shop, good)
	respond.DeleteTrolleyGoodTrue(c)
}

func GetShopBalance(c *gin.Context) {
	shop := c.PostForm("shop_name")
	m := service.GetShopBalance(shop)
	respond.BuyTrolleyGoodTrue(c, m)
}

func RegisterShop(c *gin.Context) {
	shop := c.PostForm("shop_name")
	announcement := c.PostForm("announcement")
	good := c.PostForm("good_name")
	image := c.PostForm("image")
	r := c.PostForm("range")
	color := c.PostForm("color")
	serve := c.PostForm("serve")
	message := c.PostForm("message")
	sellMessage := c.PostForm("sell_message")
	classify := c.PostForm("classify")
	money, _ := strconv.Atoi(c.PostForm("money"))
	service.RegisterShop(shop, announcement, good, money, image, r, color, serve, message, sellMessage, classify)
	respond.RegisterTrue(c)
}

package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"suning/respond"
	"suning/service"
)

func Good(c *gin.Context) {
	G := service.GetGood()
	respond.GoodTrue(c, G)
}

func FindGood(c *gin.Context) {
	message := c.PostForm("message")
	sort := c.PostForm("sort")
	G := service.GetWantGood(message, sort)
	respond.GoodTrue(c, G)
}

func GoodClassify(c *gin.Context) {
	classify := c.PostForm("classify")
	G := service.ClassifyGood(classify)
	respond.GoodTrue(c, G)
}

func GoodMessage(c *gin.Context) {
	goodName := c.PostForm("good_name")
	g := service.GetGoodMessage(goodName)
	respond.GoodMessageTrue(c, g)
}

func AddTrolley(c *gin.Context) {
	username := c.PostForm("username")
	goodName := c.PostForm("good_name")
	money, _ := strconv.Atoi(c.PostForm("money"))
	sellNumber, _ := strconv.Atoi(c.PostForm("sell_number"))
	service.AddTrolley(username, goodName, money, sellNumber)
	respond.AddTrolleyTrue(c)
}

func CollectGood(c *gin.Context) {
	username := c.PostForm("username")
	goodName := c.PostForm("good_name")
	money, _ := strconv.Atoi(c.PostForm("money"))
	sellNumber, _ := strconv.Atoi(c.PostForm("sell_number"))
	service.CollectGood(username, goodName, money, sellNumber)
	respond.CollectGoodTure(c)
}

func DeleteCollect(c *gin.Context) {
	username := c.PostForm("username")
	good := c.PostForm("good_name")
	goods := strings.Split(good, " ")
	service.DeleteCollectGood(username, goods)
	respond.DeleteTrolleyGoodTrue(c)
}

func GetCollect(c *gin.Context) {
	username := c.PostForm("username")
	C := service.GetCollect(username)
	respond.GetCollectTrue(c, C)
}

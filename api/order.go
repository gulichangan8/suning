package api

import (
	"github.com/gin-gonic/gin"
	"suning/respond"
	"suning/service"
)

func OrderSort(c *gin.Context) {
	username := c.PostForm("username")
	sort := c.PostForm("sort")
	O := service.OrderSort(username, sort)
	respond.OrderSortTrue(c, O)
}

func DeleteOrder(c *gin.Context) {
	username := c.PostForm("username")
	time := c.PostForm("time")
	service.DeleteOrder(username, time)
	respond.DeleteTrolleyGoodTrue(c)
}

func SureArrive(c *gin.Context) {
	username := c.PostForm("username")
	time := c.PostForm("time")
	service.SureArrive(username, time)
	respond.SureArriveTrue(c)
}

func CommentOrder(c *gin.Context) {
	username := c.PostForm("username")
	time := c.PostForm("time")
	comment := c.PostForm("comment")
	service.CommentOrder(username, time, comment)
	respond.CommentOrderTrue(c)
}

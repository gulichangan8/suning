package api

import (
	"github.com/gin-gonic/gin"
	"suning/respond"
	"suning/service"
)

// GetGoodComment 获得相应商品下的全部评论
func GetGoodComment(c *gin.Context) {
	goodName := c.PostForm("good_name")
	C := service.GetGoodComment(goodName)
	respond.GetGoodCommentTrue(c, C)
}

// PublishComment 发表初始评论
func PublishComment(c *gin.Context) {
	comment := c.PostForm("comment")
	writer := c.PostForm("username")
	good := c.PostForm("good_name")
	service.PublishComment(comment, writer, good)
	respond.PublishCommentTrue(c)
}

// DeleteRes 删除回复
func DeleteRes(c *gin.Context) {
	username := c.PostForm("username")
	comment := c.PostForm("comment")
	writerTo := c.PostForm("write_to")
	r := c.PostForm("respond")
	good := c.PostForm("good_name")
	resTime := c.PostForm("respond_time")
	service.DeleteRes(username, comment, writerTo, r, good, resTime)
	respond.DeleteResTrue(c)
}

// PublishRespond 回复接口
func PublishRespond(c *gin.Context) {
	username := c.PostForm("username")
	comment := c.PostForm("comment")
	writerTo := c.PostForm("write_to")
	r := c.PostForm("respond")
	time := c.PostForm("comment_time")
	good := c.PostForm("good_name")
	service.PublishRespond(username, comment, writerTo, r, time, good)
	respond.PublishRespondTrue(c)
}

// DeleteComment 删除初始评论
func DeleteComment(c *gin.Context) {
	username := c.PostForm("username")
	comment := c.PostForm("comment")
	time := c.PostForm("comment_time")
	good := c.PostForm("good_name")
	service.DeleteComment(username, comment, time, good)
	respond.DeleteComment(c)
}

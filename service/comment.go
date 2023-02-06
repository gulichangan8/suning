package service

import (
	"suning/dao"
	"suning/model"
	"time"
)

// GetGoodComment 获得该商品下的所有评论
func GetGoodComment(goodName string) model.Comments {
	C := dao.TakeOutGoodComment(goodName)
	return C
}

// PublishComment 发表评论
func PublishComment(comment, writer, good string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	dao.BringComments(comment, writer, good, t)
}

// DeleteRes 删除评论
func DeleteRes(username, comment, writerTo, respond, good, resTime string) {
	dao.DeleteRes(username, comment, writerTo, respond, good, resTime)
}

func PublishRespond(username, comment, writerTo, respond, T, good string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	dao.PublishRespond(username, comment, writerTo, respond, t, T, good)
}

func DeleteComment(username, comment, time, good string) {
	dao.DeleteComment(username, comment, time, good)
}

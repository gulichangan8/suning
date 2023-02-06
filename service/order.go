package service

import (
	"suning/dao"
	"suning/model"
)

func OrderSort(username, sort string) model.Orders {
	O := dao.OrderSort(username, sort)
	return O
}

func DeleteOrder(username, time string) {
	dao.DeleteOrder(username, time)
}

func SureArrive(username, time string) {
	dao.SureArrive(username, time)
}

func CommentOrder(username, time, comment string) {
	dao.CommentOrder(username, time, comment)
}

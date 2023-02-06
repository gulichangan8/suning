package service

import (
	"suning/dao"
	"suning/model"
)

func GetShopMessage(shop string) (model.Shop, model.Goods) {
	s, G := dao.GetShopMessage(shop)
	return s, G
}

func Classify(classify, shop string) model.Goods {
	G := dao.GetClassify(classify, shop)
	return G
}

func AddGood(shop, good string, money int, image, r, color, serve, message, sellMessage, classify string) {
	dao.AddGood(shop, good, money, image, r, color, serve, message, sellMessage, classify)
}

func DeleteGood(shop, good string) {
	dao.DeleteGood(shop, good)
}

func GetShopBalance(shop string) int {
	m := dao.GetShopBalance(shop)
	return m
}

func RegisterShop(shop, announcement, good string, money int, image, r, color, serve, message, sellMessage, classify string) {
	dao.RegisterShop(shop, announcement, good, money, image, r, color, serve, message, sellMessage, classify)
}

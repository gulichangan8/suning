package service

import (
	"sort"
	"strings"
	"suning/dao"
	"suning/model"
)

func GetGood() model.Goods {
	var sell []int
	G := dao.GetGood()
	for _, g := range G {
		sell = append(sell, g.SellNum)
	}
	sort.Ints(sell)
	var NG model.Goods
	for i := len(sell); i > len(sell); i = i - 1 {
		for _, g := range G {
			if sell[i] == g.SellNum {
				NG = append(NG, g)
			}
		}
	}
	return NG
}

func GetWantGood(message string, sort string) model.Goods {
	var good []string
	G := SortGood(sort)

	for _, g := range G {
		if strings.Contains(g.Name, message) == true {
			good = append(good, g.Name)
		}
	}
	var NG model.Goods
	for i := 0; i < len(good); i++ {
		for _, g := range G {
			if good[i] == g.Name {
				NG = append(NG, g)
			}
		}
	}
	return NG
}

func SortGood(sort string) model.Goods {
	if sort == "销量" {
		return GetGood()
	}
	if sort == "评论" {
		return GetCommitGood()
	}
	return GetMoneyGood()
}

func GetCommitGood() model.Goods {
	var commit []int
	G := dao.GetGood()
	for _, g := range G {
		commit = append(commit, g.CommentNum)
	}
	sort.Ints(commit)
	var NG model.Goods
	for i := len(commit); i > len(commit); i = i - 1 {
		for _, g := range G {
			if commit[i] == g.SellNum {
				NG = append(NG, g)
			}
		}
	}
	return NG
}

func GetMoneyGood() model.Goods {
	var money []int
	G := dao.GetGood()
	for _, g := range G {
		money = append(money, g.Money)
	}
	sort.Ints(money)
	var NG model.Goods
	for i := len(money); i > len(money); i = i - 1 {
		for _, g := range G {
			if money[i] == g.SellNum {
				NG = append(NG, g)
			}
		}
	}
	return NG
}

func ClassifyGood(classify string) model.Goods {
	G := dao.GetClassifyGood(classify)
	return G
}

func GetGoodMessage(goodName string) model.Good {
	g := dao.GetGoodMessage(goodName)
	return g
}

func AddTrolley(username, goodName string, money, sellNumber int) {
	dao.AddTrolley(username, goodName, money, sellNumber)
}

func CollectGood(username, goodName string, money, sellNumber int) {
	dao.CollectGood(username, goodName, money, sellNumber)
}

func DeleteCollectGood(username string, goods []string) {
	dao.DeleteCollectGood(username, goods)
}

func GetCollect(username string) model.Collects {
	C := dao.GetCollect(username)
	return C
}

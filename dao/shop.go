package dao

import (
	"fmt"
	"suning/model"
)

func GetShopMessage(shop string) (model.Shop, model.Goods) {
	var s model.Shop
	var g model.Good
	var G model.Goods
	var good string
	var goods []string
	rows, err := dB.Query("select good,announcement from shop where name=?", shop)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&good, &s.Announcement)
		if err != nil {
			fmt.Println(err)
		}
		goods = append(goods, good)
	}
	rows, err = dB.Query("select name,money,sell_number from good where name=?", shop)
	for rows.Next() {
		err = rows.Scan(&g.Name, &g.Money, &g.SellNum)
		if err != nil {
			fmt.Println(err)
		}
		G = append(G, g)
	}
	return s, G
}

func GetClassify(classify, shop string) model.Goods {
	var g model.Good
	var G model.Goods
	rows, err := dB.Query("select name,sell_number,money from good where classify=? && shop=?", classify, shop)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&g.Name, &g.Money, &g.SellNum)
		if err != nil {
			fmt.Println(err)
		}
		G = append(G, g)
	}
	return G
}

func AddGood(shop, good string, money int, image, r, color, serve, message, sellMessage, classify string) {
	_, err := dB.Exec("insert into shop (name, good) value (?,?)",
		shop, good)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("insert into good (money, name, shop, `range`, color, serve, message, sell_message, classify, image) value (?,?,?,?,?,?,?,?,?,?)",
		money, good, shop, r, color, serve, message, sellMessage, classify, image)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteGood(shop, good string) {
	_, err := dB.Exec("delete from shop where name=? && good=?", shop, good)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("delete from good where name=? && shop=?", good, shop)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetShopBalance(shop string) int {
	var m int
	row := dB.QueryRow("select shop_money from shop where name=?", shop)
	err := row.Scan(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func RegisterShop(shop, announcement, good string, money int, image, r, color, serve, message, sellMessage, classify string) {
	_, err := dB.Exec("insert into shop (name, good, announcement, shop_money) value (?,?,?,?)",
		shop, good, announcement, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("insert into good (money, name, shop, `range`, color, serve, message, sell_message, classify, image) value (?,?,?,?,?,?,?,?,?,?)",
		money, good, shop, r, color, serve, message, sellMessage, classify, image)
	if err != nil {
		fmt.Println(err)
		return
	}
}

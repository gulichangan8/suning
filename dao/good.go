package dao

import (
	"fmt"
	"suning/model"
)

func GetGood() model.Goods {
	var g model.Good
	var G model.Goods
	rows, err := dB.Query("select name,sell_number,money from good")
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

func GetClassifyGood(classify string) model.Goods {
	var g model.Good
	var G model.Goods
	rows, err := dB.Query("select name,sell_number,money from good where classify=?", classify)
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

func GetGoodMessage(goodName string) model.Good {
	var g model.Good
	row := dB.QueryRow("select * from good where name=?", goodName)
	err := row.Scan(&g.Money, &g.Name, &g.Shop, &g.SellNum, &g.CommentNum, &g.Range, &g.Color, &g.Serve,
		&g.Message, &g.SellMessage, &g.Classify)
	if err != nil {
		fmt.Println(err)
	}
	return g
}

func AddTrolley(username, goodName string, money, sellNumber int) {
	_, err := dB.Exec("insert into trolley (username,good,money,sell_number) value (?,?,?,?)",
		username, goodName, money, sellNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CollectGood(username, goodName string, money, sellNumber int) {
	_, err := dB.Exec("insert into collect (username,good,money,sell_number) value (?,?,?,?)",
		username, goodName, money, sellNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteCollectGood(username string, goods []string) {
	for _, good := range goods {
		_, err := dB.Exec("delete from collect where username=? && good=?", username, good)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func GetCollect(username string) model.Collects {
	var c model.Collect
	var C model.Collects
	rows, err := dB.Query("select good,money,sell_number from trolley where username=?", username)
	if err != nil {
		fmt.Println(err)
		return C
	}
	for rows.Next() {
		err = rows.Scan(&c.Good, &c.Money, &c.SellNumber)
		if err != nil {
			fmt.Println(err)
		}
		C = append(C, c)
	}
	return C
}

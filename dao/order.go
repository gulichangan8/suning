package dao

import (
	"fmt"
	"suning/model"
	"time"
)

func OrderSort(username, sort string) model.Orders {
	var o model.Order
	var O model.Orders
	rows, err := dB.Query("select good,money,addr,time from `order` where username=? && sort=?",
		username, sort)
	if err != nil {
		fmt.Println(err)
		return O
	}
	for rows.Next() {
		err = rows.Scan(&o.Good, &o.Money, &o.Addr, &o.Time)
		if err != nil {
			fmt.Println(err)
		}
		O = append(O, o)
	}
	return O
}

func DeleteOrder(username, time string) {
	_, err := dB.Exec("delete from `order` where username=? && time=?", username, time)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SureArrive(username, Time string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	_, err := dB.Exec("update  `order` set sort=?,time=? where username=? && time=? && sort=?",
		"待评价", t, username, Time, "待收货")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CommentOrder(username, Time, comment string) {
	var good string
	row := dB.QueryRow("select good from `order` where username=? && time=?", username, Time)
	err := row.Scan(&good)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	_, err = dB.Exec("insert into comment (parent_id, good, comment, writer, comment_time) value (?,?,?,?,?)",
		0, good, comment, username, t)
	if err != nil {
		fmt.Println(err)
		return
	}
}

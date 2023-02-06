package dao

import (
	"fmt"
	"suning/model"
	"time"
)

// TakeUserDate 将user表中的数据取出
func TakeUserDate() model.Users {
	var u model.User
	var U model.Users
	rows, err := dB.Query("select username,password from user")
	if err != nil {
		fmt.Println(err)
		return U
	}
	for rows.Next() {
		err = rows.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Println(err)
		}
		U = append(U, u)
	}
	return U
}

// BringUserDate 将数据存入user表
func BringUserDate(u model.User) {
	_, err := dB.Exec("insert into user (username,password) value (?,?)",
		u.Username, u.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// BringJwt 保存jwt
func BringJwt(username string, jwt string) {
	_, err := dB.Exec("insert into login (username,token) value (?,?)",
		username, jwt)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// DeleteToken 将退出登录的用户的token删除
func DeleteToken(username string) {
	_, err := dB.Exec("update login set token='' where username=?", username)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CheckState(username string) string {
	row := dB.QueryRow("select token from login where username=?", username)
	var j string
	err := row.Scan(&j)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return j
}

func GetBalance(username string) int {
	var m int
	row := dB.QueryRow("select money from user where username=?", username)
	err := row.Scan(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func Recharge(username string, money int) {
	var m int
	var p string
	row := dB.QueryRow("select money,password from user where username=?", username)
	err := row.Scan(&m, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	m = m + money
	_, err = dB.Exec("insert into user  (money,username,password) value ?",
		m, username, p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetMessage(username string) model.User {
	var u model.User
	row := dB.QueryRow("select money,age,sex,true_name,career from user where username=?", username)
	err := row.Scan(&u.Money, &u.Age, &u.Sex, &u.TrueName, &u.Career)
	if err != nil {
		fmt.Println(err)
	}
	return u
}

func ChangeMessage(username string, age int, sex, trueName, career string) {
	_, err := dB.Exec("update user set age=?,sex=?,true_name=?,career=? where username=?",
		username, age, sex, trueName, career)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetAddr(username string) model.Address {
	var a model.Addr
	var A model.Address
	rows, err := dB.Query("select addr from addr where username=?", username)
	if err != nil {
		fmt.Println(err)
		return A
	}
	for rows.Next() {
		err = rows.Scan(&a.Addr)
		if err != nil {
			fmt.Println(err)
		}
		A = append(A, a)
	}
	return A
}

func AddAddr(username, addr string) {
	_, err := dB.Exec("insert into addr (username, addr)  value (?,?)",
		username, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DeleteAddr(username, addr string) {
	_, err := dB.Exec("delete from addr where username=? && addr=?", username, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetTrolley(username string) model.Trolleys {
	var t model.Trolley
	var T model.Trolleys
	rows, err := dB.Query("select good,money,sell_number from trolley where username=?", username)
	if err != nil {
		fmt.Println(err)
		return T
	}
	for rows.Next() {
		err = rows.Scan(&t.Good, &t.Money, &t.SellNumber)
		if err != nil {
			fmt.Println(err)
		}
		T = append(T, t)
	}
	return T
}

func BuyTrolleyGood(username string, goods []string) int {
	var m int
	var M int
	for _, good := range goods {
		row := dB.QueryRow("select money from trolley where username=? && good=?", username, good)
		err := row.Scan(&m)
		if err != nil {
			fmt.Println(err)
		}
		M = M + m
	}
	return M
}

func DeleteTrolleyGood(username string, goods []string) {
	for _, good := range goods {
		_, err := dB.Exec("delete from trolley where username=? && good=?", username, good)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func SureBuy(username, t, good, addr string, money int) {
	var m int
	row := dB.QueryRow("select money from user where username=? ", username)
	err := row.Scan(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	m = m - money
	_, err = dB.Exec("update user set money=? where username=?",
		m, username)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("update `order` set time=?,sort=? where username=? && good=? && addr=?",
		t, "待收货", username, good, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	row = dB.QueryRow("select sell_number from good where name=? ", username)
	err = row.Scan(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("update good set sell_number=? where name=?",
		m+1, good)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("update collect set sell_number=? where good=?",
		m+1, good)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = dB.Exec("update trolley set sell_number=? where good=?",
		m+1, good)
	if err != nil {
		fmt.Println(err)
		return
	}
	var shop string
	row = dB.QueryRow("select shop_money,name from shop where good=? ", good)
	err = row.Scan(&m, &shop)
	if err != nil {
		fmt.Println(err)
		return
	}
	m = m + money
	_, err = dB.Exec("update shop set shop_money=? where name=?",
		m, shop)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReturnBuy(username, addr, good string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	_, err := dB.Exec("update `order` set time=?,sort=? where username=? && good=? && addr=?",
		t, "待付款", username, good, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ChooseAddr(username, addr, good string, money int) {
	_, err := dB.Exec("insert into `order` (username,good, money, addr,sort) value (?,?,?,?,?)",
		username, good, money, addr, "")
	if err != nil {
		fmt.Println(err)
		return
	}
}

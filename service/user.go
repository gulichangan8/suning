package service

import (
	"suning/dao"
	"suning/model"
	"suning/tool"
	"time"
)

// CheckUsername 检查用户名格式
func CheckUsername(username string) string {
	ok1 := tool.IsShortName(username)
	ok2 := tool.IsNameErr(username)
	ok3 := tool.IsRepeatUsername(username)
	if ok1 && ok2 {
		if ok3 {
			return "用户名正确"
		} else {
			return "用户名重复"
		}
	} else {
		return "提示：1.用户名不能含有\\字符\n2.用户名不能为空"
	}
}

// CheckPassword 检查密码格式
func CheckPassword(password string) string {
	ok1 := tool.IsPasswordErr(password)
	ok2 := tool.IsShortPassword(password)
	ok3 := tool.IsLongPassword(password)
	if ok1 && ok2 && ok3 {
		return "密码正确"
	} else {
		return "提示：1.密码不能含有\\字符\n2.密码不能小于六个字符\n3.密码不能超过12个字符"
	}
}

// BringUserDate 将用户数据存入数据库
func BringUserDate(username string, P string) {
	U := tool.CreateUser(username, P)
	dao.BringUserDate(U)
}

// DeleteToken 将退出登录的用户的token删除
func DeleteToken(username string) {
	dao.DeleteToken(username)
}

// CheckState 检查用户登录状态
func CheckState(username string, jwt string) bool {
	j := dao.CheckState(username)
	if j == jwt {
		return true
	} else {
		return false
	}
}

func Balance(username string) int {
	m := dao.GetBalance(username)
	return m
}

func Recharge(username string, money int) {
	dao.Recharge(username, money)
}

func GetMessage(username string) model.User {
	u := dao.GetMessage(username)
	return u
}

func ChangeMessage(username string, age int, sex, trueName, career string) {
	dao.ChangeMessage(username, age, sex, trueName, career)
}

func GetAddr(username string) model.Address {
	A := dao.GetAddr(username)
	return A
}

func AddAddr(username, addr string) {
	dao.AddAddr(username, addr)
}

func DeleteAddr(username, addr string) {
	dao.DeleteAddr(username, addr)
}

func GetTrolley(username string) model.Trolleys {
	T := dao.GetTrolley(username)
	return T
}

func BuyTrolleyGood(username string, goods []string) int {
	m := dao.BuyTrolleyGood(username, goods)
	return m
}

func DeleteTrolleyGood(username string, goods []string) {
	dao.DeleteTrolleyGood(username, goods)
}

func SureBuy(username, good, addr string, money int) {
	t := time.Now().Format("2006-01-02 15:04:05")
	dao.SureBuy(username, t, good, addr, money)
}

func ReturnBuy(username, addr, good string) {
	dao.ReturnBuy(username, addr, good)
}

func ChooseAddr(username, addr, good string, money int) {
	dao.ChooseAddr(username, addr, good, money)
}

func CheckLogin(username string, password string) bool {
	U := dao.TakeUserDate()
	ok := false
	for _, date := range U {
		if date.Username == username && date.Password == password {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

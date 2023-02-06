package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"suning/dao"
	"suning/model"
	"suning/respond"
	"suning/service"
	"suning/tool"
	"time"
)

// Register 注册接口
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	r1 := service.CheckUsername(username)
	r2 := service.CheckPassword(password)
	P := tool.Hmac("a", password)
	if r1 == "用户名正确" && r2 == "密码正确" {
		respond.RegisterTrue(c)
		service.BringUserDate(username, P)
	} else {
		respond.RegisterErr(c, r1, r2)
	}
}

// Login 登录接口
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	P := tool.Hmac("a", password)
	var h model.Header
	var p model.Payload
	key := "gulichangan"
	ok := service.CheckLogin(username, P)
	if ok {
		h.Alg = "HS256"
		h.Typ = "JWT"
		p.Aud = username
		p.Iat = time.Now().Format("2006-01-02-15:04:05") + "," + strconv.FormatInt(time.Now().Unix(), 10)
		p.Exp = time.Unix(time.Now().Unix()+24*60*60*1000, 0).Format("2006-01-02-15:04:05")
		p.Sub = "Login"
		date := base64.URLEncoding.EncodeToString([]byte(h.Alg+" "+h.Typ)) + "." +
			base64.URLEncoding.EncodeToString([]byte(p.Aud+" "+p.Iat+" "+p.Exp+" "+p.Sub))
		signature := HmacSha(key, date)
		j := date + "." + signature
		dao.BringJwt(username, j)
		c.JSON(200, gin.H{
			"status": 20001,
			"res":    "登陆成功",
			"JWT":    j,
		})
	} else {
		jwt := c.Request.Header.Get("Authorization")
		if jwt == "" {
			println("sdfghj")
			respond.LoginErr(c)
		} else {
			jwts := strings.Split(jwt, ".")
			H := jwts[0]
			fmt.Println(H)
			header, _ := base64.URLEncoding.DecodeString(H)
			headers := strings.Split(string(header), " ")
			h.Alg = headers[0]
			h.Typ = headers[1]
			Pd := jwts[1]
			fmt.Println(Pd)
			pd, _ := base64.URLEncoding.DecodeString(Pd)
			payloads := strings.Split(string(pd), " ")
			signature := jwts[2]
			p.Aud = payloads[0]
			p.Iat = payloads[1]
			p.Exp = payloads[2]
			p.Sub = payloads[3]
			date := base64.URLEncoding.EncodeToString([]byte(h.Alg+" "+h.Typ)) + "." +
				base64.URLEncoding.EncodeToString([]byte(p.Aud+" "+p.Iat+" "+p.Exp+" "+p.Sub))
			fmt.Println(date)
			exps := strings.Split(p.Iat, ",")
			e, _ := strconv.Atoi(exps[1])
			if e+24*60*60*1000 <= int(time.Now().Unix()) {
				c.JSON(404, gin.H{
					"status": 20002,
					"res":    "登陆失败",
					"JWT":    "已过期",
				})
				c.Abort()
			}
			ok = VerifyHmac(signature, date, key)
			if !ok {
				c.JSON(404, gin.H{
					"status": 20002,
					"res":    "登陆失败",
					"JWT":    "验证失败",
				})
				c.Abort()
			}
		}
	}
}

// QuitLogin 退出登录接口
func QuitLogin(c *gin.Context) {
	username := c.PostForm("username")
	service.DeleteToken(username)
	respond.QuitLoginTrue(c)
}

// State 查询登录状态接口
func State(c *gin.Context) {
	username := c.PostForm("username")
	jwt := c.Request.Header.Get("Authorization")
	ok := service.CheckState(username, jwt)
	if ok {
		respond.StateTure(c)
	} else {
		respond.StateErr(c)
	}
}

func Balance(c *gin.Context) {
	username := c.PostForm("username")
	m := service.Balance(username)
	respond.BalanceTrue(c, m)
}

func Recharge(c *gin.Context) {
	username := c.PostForm("username")
	money, _ := strconv.Atoi(c.PostForm("money"))
	service.Recharge(username, money)
	respond.RechargeTrue(c)
}

func GetMessage(c *gin.Context) {
	username := c.PostForm("username")
	u := service.GetMessage(username)
	respond.GetMessageTrue(c, u)
}

func ChangeMessage(c *gin.Context) {
	username := c.PostForm("username")
	age, _ := strconv.Atoi(c.PostForm("age"))
	sex := c.PostForm("sex")
	trueName := c.PostForm("true_name")
	career := c.PostForm("career")
	service.ChangeMessage(username, age, sex, trueName, career)
	respond.ChangeMessageTrue(c)
}

func GetAddr(c *gin.Context) {
	username := c.PostForm("username")
	A := service.GetAddr(username)
	respond.GetAddrTrue(c, A)
}

func AddAddr(c *gin.Context) {
	username := c.PostForm("username")
	addr := c.PostForm("addr")
	service.AddAddr(username, addr)
	respond.AddAddrTrue(c)
}

func DeleteAddr(c *gin.Context) {
	username := c.PostForm("username")
	addr := c.PostForm("addr")
	service.DeleteAddr(username, addr)
	respond.DeleteAddrTrue(c)
}

func GetTrolley(c *gin.Context) {
	username := c.PostForm("username")
	T := service.GetTrolley(username)
	respond.GetTrolleyTrue(c, T)
}

func BuyTrolleyGood(c *gin.Context) {
	username := c.PostForm("username")
	good := c.PostForm("good_name")
	goods := strings.Split(good, " ")
	m := service.BuyTrolleyGood(username, goods)
	respond.BuyTrolleyGoodTrue(c, m)
}

func DeleteTrolleyGood(c *gin.Context) {
	username := c.PostForm("username")
	good := c.PostForm("good_name")
	goods := strings.Split(good, " ")
	service.DeleteTrolleyGood(username, goods)
	respond.DeleteTrolleyGoodTrue(c)
}

func ChooseAddr(c *gin.Context) {
	username := c.PostForm("username")
	addr := c.PostForm("addr")
	good := c.PostForm("good_name")
	money, _ := strconv.Atoi(c.PostForm("money"))
	service.ChooseAddr(username, addr, good, money)
	respond.ChooseAddrTrue(c)
}

func SureBuy(c *gin.Context) {
	username := c.PostForm("username")
	money, _ := strconv.Atoi(c.PostForm("money"))
	addr := c.PostForm("addr")
	good := c.PostForm("good_name")
	service.SureBuy(username, good, addr, money)
	respond.SureBuyTrue(c)
}

func ReturnBuy(c *gin.Context) {
	username := c.PostForm("username")
	addr := c.PostForm("addr")
	good := c.PostForm("good_name")
	service.ReturnBuy(username, addr, good)
	respond.ReturnBuyTrue(c)
}

func HmacSha(key, date string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(date))
	return hex.EncodeToString(hash.Sum(nil))
}

func VerifyHmac(signature string, date string, key string) bool {
	s := []byte(signature)
	now := HmacSha(key, date)
	return hmac.Equal(s, []byte(now))
}

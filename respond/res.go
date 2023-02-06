package respond

import (
	"github.com/gin-gonic/gin"
	"suning/model"
)

/*
	状态码表示:
	20001:成功
	20002:失败
*/

func RegisterTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "注册成功",
	})
}

func RegisterErr(c *gin.Context, r1, r2 string) {
	c.JSON(200, gin.H{
		"status":       20002,
		"username_res": r1,
		"pwd_res":      r2,
	})
}

func LoginErr(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20002,
		"res":    "登陆失败",
	})
}

func QuitLoginTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "退出登陆成功",
	})
}

func StateTure(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "已登录",
	})
}

func StateErr(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20002,
		"res":    "未登录",
	})
}

func GoodTrue(c *gin.Context, G model.Goods) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   G,
	})
}

func GoodMessageTrue(c *gin.Context, g model.Good) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   g,
	})
}

func GetGoodCommentTrue(c *gin.Context, C model.Comments) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   C,
	})
}

func PublishCommentTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "发表成功",
	})
}

func DeleteResTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "删除成功",
	})
}

func PublishRespondTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "发表成功",
	})
}

func DeleteComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "删除成功",
	})
}

func GetShopTrue(c *gin.Context, s model.Shop, G model.Goods) {
	c.JSON(200, gin.H{
		"status": 20001,
		"shop":   s,
		"good":   G,
	})
}

func BalanceTrue(c *gin.Context, m int) {
	c.JSON(200, gin.H{
		"status": 20001,
		"money":  m,
	})
}

func RechargeTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "充值成功",
	})
}

func GetMessageTrue(c *gin.Context, u model.User) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   u,
	})
}

func ChangeMessageTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "修改成功",
	})
}

func GetAddrTrue(c *gin.Context, A model.Address) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   A,
	})
}

func AddAddrTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "添加成功",
	})
}

func DeleteAddrTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "删除成功",
	})
}

func GetTrolleyTrue(c *gin.Context, T model.Trolleys) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   T,
	})
}

func BuyTrolleyGoodTrue(c *gin.Context, m int) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   m,
	})
}

func DeleteTrolleyGoodTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "删除成功",
	})
}

func AddTrolleyTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "添加成功",
	})
}

func CollectGoodTure(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "添加成功",
	})
}

func GetCollectTrue(c *gin.Context, C model.Collects) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   C,
	})
}

func ChooseAddrTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "选择成功",
	})
}

func SureBuyTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "支付成功",
	})
}

func ReturnBuyTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
	})
}

func OrderSortTrue(c *gin.Context, O model.Orders) {
	c.JSON(200, gin.H{
		"status": 20001,
		"data":   O,
	})
}

func SureArriveTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "确认成功",
	})
}

func CommentOrderTrue(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 20001,
		"res":    "评论已发表",
	})
}

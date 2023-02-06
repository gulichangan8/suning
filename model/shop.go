package model

type Shop struct {
	Name         string `json:"name"`
	Good         string `json:"good"`
	Announcement string `json:"announcement"`
	ShopMoney    int    `json:"shop_money"`
}

type Shops []Shop

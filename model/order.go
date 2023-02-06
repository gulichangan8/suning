package model

type Order struct {
	Good     string `json:"good"`
	Money    int    `json:"money"`
	Sort     string `json:"sort"`
	Time     string `json:"time"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
}

type Orders []Order

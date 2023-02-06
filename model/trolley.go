package model

type Trolley struct {
	Username   string `json:"username"`
	Good       string `json:"good"`
	Money      int    `json:"money"`
	SellNumber int    `json:"sell_number"`
}

type Trolleys []Trolley

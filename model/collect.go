package model

type Collect struct {
	Username   string `json:"username"`
	Good       string `json:"good"`
	Money      int    `json:"money"`
	SellNumber int    `json:"sell_number"`
}

type Collects []Collect

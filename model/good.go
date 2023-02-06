package model

type Good struct {
	Name        string `json:"name"`
	Money       int    `json:"money"`
	Shop        string `json:"shop"`
	SellNum     int    `json:"sell_num"`
	CommentNum  int    `json:"comment_num"`
	Range       string `json:"range"`
	Color       string `json:"color"`
	Serve       string `json:"serve"`
	Message     string `json:"message"`
	SellMessage string `json:"sell_message"`
	Classify    string `json:"classify"`
}

type Goods []Good

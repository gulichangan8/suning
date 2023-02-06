package model

type Comment struct {
	Id          int
	ParentId    int
	Good        string `json:"good"`
	Writer      string `json:"writer"`
	Comment     string `json:"comment"`
	CommentTime string `json:"comment_time"`
	Respond     string `json:"respond"`
	RespondTime string `json:"respond_time"`
	CommentNum  int    `json:"comment_num"`
	ResWriter   string `json:"res_writer"`
}

type Comments []Comment

package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	TrueName string `json:"true_name"`
	Career   string `json:"career"`
	Money    int    `json:"money"`
}

type Users []User

type Addr struct {
	Username string `json:"username"`
	Addr     string `json:"addr"`
}

type Address []Addr

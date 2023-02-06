package main

import (
	"suning/api"
	"suning/dao"
)

func main() {
	dao.InitMySql()
	api.HttpInitEngine()
}

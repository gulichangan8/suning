package tool

import "suning/model"

func CreateUser(username string, password string) model.User {
	var U model.User
	U.Username = username
	U.Password = password
	return U
}

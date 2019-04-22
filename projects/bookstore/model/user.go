package model

import (
	"goweb_code/bookstore/utils"
)

type Users struct {
	ID       string
	Username string
	Password string
	Email    string
}

//生成CartID
func (user *Users) GetUserID() string {
	return utils.CreateUUID()
}

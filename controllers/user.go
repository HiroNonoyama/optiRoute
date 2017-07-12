package controllers

import (
	"github.com/hirononoyama/optiRoute/models"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (c User) SignUp(age int, sex int) interface{} {
	repo := models.NewUserRepository()
	user := repo.Create(age, sex)
	return user
}
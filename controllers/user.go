package controllers

import (
	"../models"
	"time"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// なぜかここにintが渡るのがただしいと思いこんでいる
func (c User) SignUp(birthday string, sex int) interface{} {
	repo := models.NewUserRepository()
	user := repo.Create(birthday, sex)
	return user
}

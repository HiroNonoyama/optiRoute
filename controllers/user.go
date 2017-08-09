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

func (c User) SignUp(birthday time.Time, sex int) interface{} {
	repo := models.NewUserRepository()
	user := repo.Create(birthday, sex)
	return user
}

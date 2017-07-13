package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type User struct {
	gorm.Model
	Age int `sql: "not null"`
	Sex int `sql: "not null"`
}

func init() {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	defer db.Close()
}

func NewUser(age int, sex int) User {
	return User{
		Age: age,
		Sex: sex,
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

type UserRepository struct {
}

func (m UserRepository) Create(age int, sex int) *User {
	db, err := gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	user := NewUser(age, sex)
	db.Create(&user)
	return &user
}
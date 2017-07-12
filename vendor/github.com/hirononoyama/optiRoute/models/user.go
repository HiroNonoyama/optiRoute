package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Age int `sql: "not null"`
	Sex int `sql: "not null"`
}

func init() {
	db, err := gorm.Open("mysql", "hirononoyama:hiro0117@/optiRoute?parseTime=true")
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
	db, _ := gorm.Open("mysql", "hirononoyama:hiro0117@/optiRoute?parseTime=true")
	user := NewUser(age, sex)
	db.Create(&user)
	return &user
}

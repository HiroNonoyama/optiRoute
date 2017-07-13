package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

type User struct {
	gorm.Model
	Age int `sql: "not null"`
	Sex int `sql: "not null"`
}

func init() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
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
	db, _ := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	user := NewUser(age, sex)
	db.Create(&user)
	return &user
}
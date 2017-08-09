package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	gorm.Model
	Birthday *time.Time
	Sex int `sql: "not null"`
}

func init() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.Model(&User{}).DropColumn("age")
	defer db.Close()
}

func NewUser(birthday time.Time, sex int) User {
	return User{
		Birthday: birthday,
		Sex: sex,
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

type UserRepository struct {
}

func (m UserRepository) Create(birthday time.Time, sex int) *User {
	// db, _ := gorm.Open("mysql", "hirononoyama:hiro0117@/optiRoute?parseTime=true")
	db, _ := gorm.Open("postgres", ENV["DATABASE_URL"])
	user := NewUser(birthday, sex)
	db.Create(&user)
	return &user
}

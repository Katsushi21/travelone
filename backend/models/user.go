package postgres

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	email    string
	password string
}

type Profile struct {
	gorm.Model
	name         string
	gender       string
	avatar       string
	introduction string
}

type Post struct {
	gorm.Model
	userno int
}

func main() {
	// db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
}

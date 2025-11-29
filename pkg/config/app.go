package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		panic("database environments are not set!")
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", user, pass, host, port, name)
	fmt.Println("DSN", dsn)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}

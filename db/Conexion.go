package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func GetConexion() *gorm.DB {
	var user, pass, host, port, db string
	user = os.Getenv("USER")
	pass = os.Getenv("PASSWORD")
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	db = os.Getenv("DATABASE")

	strconnect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)
	DB, _ = gorm.Open("mysql", strconnect)
	return DB
}

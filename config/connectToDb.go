package config

import (
	"backendgo/helper"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	//Getenv is taking from system var with name MYSQL_PASSWORD MYSQL_LOGIN
	dblogin := os.Getenv("MYSQL_LOGIN")
	dbpw := os.Getenv("MYSQL_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@/goschema?charset=utf8mb4&parseTime=True&loc=Local", dblogin, dbpw)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}

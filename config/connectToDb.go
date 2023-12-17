package config

import (
	"backendgo/helper"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	//Getenv is taking from system var with name MYSQL_PASSWORD
	fmt.Println("\nPASSWORD IS: " + os.Getenv("MYSQL_PASSWORD") + "\n")
	pw := os.Getenv("MYSQL_PASSWORD")

	dsn := fmt.Sprintf("root:%s@/goschema?charset=utf8mb4&parseTime=True&loc=Local", pw)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}

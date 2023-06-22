package config

import (
	"fmt"

	// "gorm.io/gorm"
	// _ "gorm.io/driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3308)/go_crud_app?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println("Hello")
	// 	panic(err)
	// }
	dsn := "root:root@tcp(127.0.0.1:3308)/go_crud_app?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	db = d
	return db

}

func GetDB() *gorm.DB {
	return db
}

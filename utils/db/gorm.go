package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root@tcp(127.0.0.1:3306)/tugas_sql_bri"), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}

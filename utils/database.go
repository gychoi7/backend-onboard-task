package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"onycom/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/onycom?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("데이터베이스 연결 실패", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Token{})
	if err != nil {
		log.Fatal("테이블 마이그레이션 실패", err)
	}
	fmt.Println("데이터베이스 연결 성공")
}

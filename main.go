package main

import (
	_ "onycom/docs" // 이 패키지를 import하여 Swagger 문서를 포함합니다.
	"onycom/routers"
	"onycom/utils"
)

// @title Onycom API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://localhost:8080/
// @contact.name API Support
// @contact.url http://localhost:8080/
// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	utils.InitDB()

	router := routers.InitRouter()

	router.Run(":8080")
}

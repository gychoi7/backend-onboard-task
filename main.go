package main

import (
	"onycom/routers"
	"onycom/utils"
)

func main() {
	utils.InitDB()

	router := routers.InitRouter()

	router.Run(":8080")
}

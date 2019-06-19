package main

import (
	"github.com/evsamsonov/trading-history-server/database"
	"github.com/evsamsonov/trading-history-server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	defer db.Close()

	database.Migrate(db)

	router := gin.Default()
	router.POST("/bars/", handler.CreateBars(db))

	router.Run()
}

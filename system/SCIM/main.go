package main

import (
	"github.com/acexra/scim/config"
	"github.com/acexra/scim/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db := config.ConnectDB();
	defer db.Close()

	router := gin.Default()


	router.GET("/test",handlers.Test)
	router.GET("/users/:id",handlers.GetById)
	router.POST("/users",handlers.Create)

	router.Run(":8080")
}
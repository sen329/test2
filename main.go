package main

import (
	"./config"
	"./controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetUser)
	router.GET("/persons", inDB.GetUsers)
	router.POST("/person", inDB.CreateUser)
	router.PUT("/person", inDB.UpdateUser)
	router.DELETE("/person/:id", inDB.DeleteUser)
	router.Run(":8000")
}

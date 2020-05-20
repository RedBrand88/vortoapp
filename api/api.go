package main

import (
	"database/sql"
	"fmt"

	"todoapp/vortoapp/api/todo"
	"todoapp/vortoapp/api/handler"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./ambie.db")
	todo.CheckError(err)
	todos := todo.GetDBConnection(db)

	router := gin.Default()

	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todo", handler.TodoGet(todos))
	router.POST("/todo", handler.TodoPost(todos))

	fmt.Println("Running on port: 5000")
	router.Run("0.0.0.0:5000")
}
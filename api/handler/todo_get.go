package handler

import (
	"net/http"
	"todoapp/vortoapp/api/todo"
	"github.com/gin-gonic/gin"
)

//TodoGet returns all todo items from todo_items table
func TodoGet(todos todo.Store) gin.HandlerFunc {
	return func (c *gin.Context)  {
		response := todos.Get()
		c.JSON(http.StatusOK, response)		
	}
}
package handler

import (
	"net/http"
	"todoapp/vortoapp/api/todo"
	"github.com/gin-gonic/gin"
)

// TodoPutRequest struct for [PUT] request
type TodoPutRequest struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Complete int `json:"complete`
}

// I'm missing something here I can't get the put request to work
// TodoPut [PUT] /todo
func TodoPut(myTodo todo.Store) gin.HandlerFunc {
	return func (c *gin.Context) {
		requestBody := TodoPutRequest{}
		err := c.Bind(&requestBody)
		todo.CheckError(err)
		item := todo.Item{
			ID: requestBody.ID,
			Text: requestBody.Text,
			Complete: requestBody.Complete,
		}

		myTodo.Update(item)

		c.Status(http.StatusNoContent)
	}
}
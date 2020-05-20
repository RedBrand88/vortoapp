package handler

import (
	"net/http"
	"todoapp/vortoapp/api/todo"
	"github.com/gin-gonic/gin"
)

// TodoPostRequest [POST] /todo
type TodoPostRequest struct {
	Text string `json:"text"`
	Complete int `json:"complete"`
}

//TodoPost is a post endpoint /todo
func TodoPost(myTodo todo.Store) gin.HandlerFunc {
	return func (c *gin.Context)  {
		requestBody := TodoPostRequest{}
		c.Bind(&requestBody)

		item := todo.Item{
			Text: requestBody.Text,
			Complete: requestBody.Complete,
		}

		myTodo.Set(item)

		c.Status(http.StatusNoContent)
	}
}
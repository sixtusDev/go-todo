package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []Todo {
	{Id: "1", Title: "Buy milk", Completed: true},
	{Id: "2", Title: "Go to the gym", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
    router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}


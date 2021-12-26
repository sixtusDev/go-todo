package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []Todo {
	{Id: uuid.NewString(), Title: "Buy milk", Completed: true},
	{Id: uuid.NewString(), Title: "Go to the gym", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodo)
	router.POST("/todo/post", postTodo)
    router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	newTodo.Id =  uuid.NewString()
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo);
}

func getTodo(c *gin.Context) {
	paramsId := c.Param("id")

	for _, todo := range todos {
		if todo.Id == paramsId{
			c.IndentedJSON(http.StatusFound, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo with the given id does not exist"});
}


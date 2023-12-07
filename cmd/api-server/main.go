package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumaresan1983/todoserver/pkg/controllers"
	"github.com/kumaresan1983/todoserver/pkg/initializers"
	"github.com/kumaresan1983/todoserver/pkg/middleware"
)

var server *gin.Engine

func init() {
	initializers.ConnectDB()

	server = gin.Default()

}

func main() {
	// version 1
	apiV1 := server.Group("/v1/api")

	server.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Simple To-Do Server in Golang"})
	})

	// User only can be added by authorized person
	authV1 := apiV1.Group("/")

	authV1.GET("/auth/google/login", controllers.GoogleLogin)
	authV1.GET("/auth/google/callback", controllers.GoogleCallback)

	authV1.GET("/todo", middleware.Auth(), controllers.GetTodos)
	authV1.GET("/todo/:id", middleware.Auth(), controllers.GetTodoByID)
	authV1.PUT("/todo", middleware.Auth(), controllers.CreateTodo)
	authV1.DELETE("/todo/:id", middleware.Auth(), controllers.DeleteTodoByID)
	authV1.PATCH("/todos/:id/complete", middleware.Auth(), controllers.CompleteTodoByID)

	log.Fatal(server.Run(":" + "8080"))

}

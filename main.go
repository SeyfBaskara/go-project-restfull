package main

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/go-project-restfull/initializers"
	"github.com/seyfBaskara/go-project-restfull/controllers"
	"github.com/seyfBaskara/go-project-restfull/routes"
)

var (
	server              *gin.Engine
	
	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)


func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = gin.Default()
}


func main (){
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PostRouteController.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
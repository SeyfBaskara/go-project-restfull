package main

import (
	"fmt"
	"log"
	"github.com/seyfBaskara/go-project-restfull/initializers"
	"github.com/seyfBaskara/go-project-restfull/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}


func main (){
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Migration complete")
}

/*
*FIXME 
* should create post table succesfuly 
*/
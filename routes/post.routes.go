package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seyfBaskara/go-project-restfull/controllers"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.FindPosts)
	router.GET("/:postId", pc.postController.FindPostById)
	router.PUT("/:postId", pc.postController.UpdatePost)
	router.DELETE("/:postId", pc.postController.DeletePost)

}
package router

import (
	"backendgo/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	api := router.Group("/api")
	{
		tagRouter := api.Group("/tags")
		tagRouter.GET("", tagsController.FindAll)
		tagRouter.GET("/:tagId", tagsController.FindById)
		tagRouter.POST("", tagsController.Create)
		tagRouter.PATCH("/:tagId", tagsController.Update)
		tagRouter.DELETE("/:tagId", tagsController.Delete)
	}

	return router

}

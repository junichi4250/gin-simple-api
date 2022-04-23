package route

import (
	"gin-api/handler"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/recipes", handler.NewRecipeHandler)
	router.GET("/recipes", handler.ListRecipesHandler)
	router.PUT("/recipes/:id", handler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", handler.DeleteRecipeHandler)
	router.GET("/recipes/search", handler.SearchRecipeHandler)
	return router
}
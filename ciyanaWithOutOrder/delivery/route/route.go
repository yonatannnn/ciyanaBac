package route

import (
	"ciyana/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl domain.Controller) *gin.Engine {
	r := gin.Default()
	r.POST("/items", ctrl.CreateItem)
	r.GET("/items/:id", ctrl.GetItem)
	r.GET("/items", ctrl.GetItems)
	r.PUT("/items/:id", ctrl.UpdateItem)
	r.DELETE("/items/:id", ctrl.DeleteItem)
	r.GET("/items/filter/category/:category_id", ctrl.FilterByCategory)
	r.GET("/items/search", ctrl.SearchItem)
	r.GET("/items/filter/tag/:tag", ctrl.FilterByTag)

	r.POST("/categories", ctrl.CreateCategory)
	r.GET("/categories/:id", ctrl.GetCategory)
	r.GET("/categories", ctrl.GetCategories)
	r.PUT("/categories/:id", ctrl.UpdateCategory)

	return r
}

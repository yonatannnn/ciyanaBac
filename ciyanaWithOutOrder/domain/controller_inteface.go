package domain

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateItem(c *gin.Context)
	GetItem(c *gin.Context)
	GetItems(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
	CreateCategory(c *gin.Context)
	GetCategory(c *gin.Context)
	GetCategories(c *gin.Context)
	UpdateCategory(c *gin.Context)
	FilterByCategory(c *gin.Context)
	SearchItem(c *gin.Context)
	FilterByTag(c *gin.Context)
}

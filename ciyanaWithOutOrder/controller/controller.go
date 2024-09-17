package controller

import (
	"ciyana/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type controller struct {
	itemUsecase     domain.ItemUsecase
	categoryUsecase domain.CategoryUsecase
}

func NewController(itemUsecase domain.ItemUsecase, categoryUsecase domain.CategoryUsecase) domain.Controller {
	return &controller{
		itemUsecase:     itemUsecase,
		categoryUsecase: categoryUsecase,
	}
}

func (ctrl *controller) CreateItem(c *gin.Context) {
	item := domain.Item{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.itemUsecase.CreateItem(&item); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": item})
}

func (ctrl *controller) GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := ctrl.itemUsecase.GetItem(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": item})
}

func (ctrl *controller) GetItems(c *gin.Context) {
	items, err := ctrl.itemUsecase.GetItems()
	fmt.Println("Retrieved items:", items)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items})
}

func (ctrl *controller) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	item := domain.Item{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.itemUsecase.UpdateItem(id, &item)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": item})

}

func (ctrl *controller) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.itemUsecase.DeleteItem(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": "Item deleted successfully"})
}

func (ctrl *controller) CreateCategory(c *gin.Context) {
	category := domain.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.categoryUsecase.CreateCategory(&category); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": category})
}

func (ctrl *controller) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := ctrl.categoryUsecase.GetCategory(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": category})
}

func (ctrl *controller) GetCategories(c *gin.Context) {
	categories, err := ctrl.categoryUsecase.GetCategories()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": categories})
}

func (ctrl *controller) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	category := domain.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := ctrl.categoryUsecase.UpdateCategory(id, &category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": category})
}

func (ctrl *controller) FilterByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")
	items, err := ctrl.itemUsecase.FilterByCategory(categoryID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items})
}

func (ctrl *controller) SearchItem(c *gin.Context) {
	query := c.Query("q")
	items, err := ctrl.itemUsecase.SearchItem(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items})
}

func (ctrl *controller) FilterByTag(c *gin.Context) {
	tag := c.Param("tag")
	items, err := ctrl.itemUsecase.FilterByTag(tag)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items})
}

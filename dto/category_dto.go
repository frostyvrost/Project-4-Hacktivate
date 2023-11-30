package dto

import (
	"net/http"
	"project-4/models"
	"project-4/pkg"
	"project-4/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(context *gin.Context) {
	var category models.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	categoryResponse, err := service.CategoryService.CreateCategory(&category)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"created_at":          categoryResponse.CreatedAt,
	})
}

func GetAllCategories(context *gin.Context) {
	categories, err := service.CategoryService.GetAllCategories()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, categories)
}

func UpdateCategory(context *gin.Context) {
	var categoryUpdated models.CategoryUpdate

	if err := context.ShouldBindJSON(&categoryUpdated); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	id, _ := pkg.GetIdParam(context, "categoryId")

	categoryResponse, err := service.CategoryService.UpdateCategory(&categoryUpdated, id)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":                  categoryResponse.ID,
		"type":                categoryResponse.Type,
		"sold_product_amount": categoryResponse.SoldProductAmount,
		"updated_at":          categoryResponse.UpdatedAt,
	})
}

func DeleteCategory(context *gin.Context) {
	categoryId, _ := pkg.GetIdParam(context, "categoryId")

	err := service.CategoryService.DeleteCategory(categoryId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}

package dto

import (
	"net/http"
	"project-4/models"
	"project-4/pkg"
	"project-4/service"

	"github.com/gin-gonic/gin"
)

func CreateProduct(context *gin.Context) {
	var product models.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	productResponse, err := service.ProductService.CreateProduct(&product)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":          productResponse.ID,
		"title":       productResponse.Title,
		"price":       productResponse.Price,
		"stock":       productResponse.Stock,
		"category_id": productResponse.CategoryID,
		"created_at":  productResponse.CreatedAt,
	})
}

func GetAllProducts(context *gin.Context) {
	products, err := service.ProductService.GetAllProducts()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	var productMaps []map[string]interface{}

	for _, product := range products {
		productMap := map[string]interface{}{
			"id":          product.ID,
			"title":       product.Title,
			"price":       product.Price,
			"stock":       product.Stock,
			"category_Id": product.CategoryID,
			"created_at":  product.CreatedAt,
		}

		productMaps = append(productMaps, productMap)
	}

	context.JSON(http.StatusOK, productMaps)
}

func UpdateProduct(context *gin.Context) {
	var productUpdated models.ProductUpdate

	if err := context.ShouldBindJSON(&productUpdated); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	id, _ := pkg.GetIdParam(context, "productId")

	productResponse, err := service.ProductService.UpdateProduct(&productUpdated, id)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result := map[string]interface{}{
		"id":         productResponse.ID,
		"title":      productResponse.Title,
		"price":      productResponse.Price,
		"stock":      productResponse.Stock,
		"CategoryId": productResponse.CategoryID,
		"createdAt":  productResponse.CreatedAt,
		"updatedAt":  productResponse.UpdatedAt,
	}

	context.JSON(http.StatusOK, gin.H{
		"product": result,
	})
}

func DeleteProduct(context *gin.Context) {
	productId, _ := pkg.GetIdParam(context, "productId")

	err := service.ProductService.DeleteProduct(productId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Product has been successfully deleted",
	})
}

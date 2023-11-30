package dto

import (
	"net/http"
	"project-4/models"
	"project-4/pkg"
	"project-4/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateTransaction(context *gin.Context) {
	var transaction models.TransactionCreate

	if err := context.ShouldBindJSON(&transaction); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	transactionResponse, err := service.TransactionService.CreateTransaction(&transaction, userId)
	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "You have successfully purchased product",
		"transaction_bill": gin.H{
			"total_price":   transactionResponse.TotalPrice,
			"quantity":      transactionResponse.Quantity,
			"product_title": transactionResponse.Product.Title,
		},
	})
}

func GetTransactionsByUserID(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	results, err := service.TransactionService.GetTransactionsByUserID(userId)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	transactions := make([]gin.H, 0, len(results))

	for _, result := range results {
		transaction := gin.H{
			"id":          result.ID,
			"product_id":  result.ProductID,
			"user_id":     result.UserID,
			"quantity":    result.Quantity,
			"total_price": result.TotalPrice,
			"Product": gin.H{
				"id":          result.Product.ID,
				"title":       result.Product.Title,
				"price":       result.Product.Price,
				"stock":       result.Product.Stock,
				"category_Id": result.Product.CategoryID,
				"created_at":  result.Product.CreatedAt,
				"updated_at":  result.Product.UpdatedAt,
			},
		}

		transactions = append(transactions, transaction)
	}

	context.JSON(http.StatusOK, transactions)

}

func GetAllTransaction(context *gin.Context) {

	results, err := service.TransactionService.GetAllTransaction()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	transactions := make([]gin.H, 0, len(results))

	for _, result := range results {
		transaction := gin.H{
			"id":          result.ID,
			"product_id":  result.ProductID,
			"user_id":     result.UserID,
			"quantity":    result.Quantity,
			"total_price": result.TotalPrice,
			"Product": gin.H{
				"id":          result.Product.ID,
				"title":       result.Product.Title,
				"price":       result.Product.Price,
				"stock":       result.Product.Stock,
				"category_Id": result.Product.CategoryID,
				"created_at":  result.Product.CreatedAt,
				"updated_at":  result.Product.UpdatedAt,
			},
			"User": gin.H{
				"id":         result.User.ID,
				"email":      result.User.Email,
				"full_name":  result.User.FullName,
				"balance":    result.User.Balance,
				"created_at": result.User.CreatedAt,
				"updated_at": result.User.UpdatedAt,
			},
		}

		transactions = append(transactions, transaction)
	}

	context.JSON(http.StatusOK, transactions)
}

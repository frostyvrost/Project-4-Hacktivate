package dto

import (
	"fmt"
	"net/http"
	"project-4/models"
	"project-4/pkg"
	"project-4/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	user.Role = "customer"

	userResponse, err := service.UserService.Register(&user)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         userResponse.ID,
		"full_name":  userResponse.FullName,
		"email":      userResponse.Email,
		"password":   userResponse.Password,
		"balance":    userResponse.Balance,
		"created_at": userResponse.CreatedAt,
	})
}

func Login(context *gin.Context) {
	var userLogin models.LoginCredential

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	token, err := service.UserService.Login(&userLogin)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateBalance(context *gin.Context) {
	var balance models.BalanceUpdate

	if err := context.ShouldBindJSON(&balance); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userId := int(userData["id"].(float64))

	updatedBalance, err := service.UserService.UpdateBalance(&balance, userId)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp %d", updatedBalance),
	})
}

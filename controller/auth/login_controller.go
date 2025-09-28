package auth

import (
	"api-book/middleware"
	"api-book/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {
	var data model.User

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
		return
	}

	if data.Username == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Need Valid Input",
		})
		return
	}

	if data.Username != "admin" || data.Password != "12345" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credential",
		})
		return
	}

	token, err := middleware.GenerateJwt(data.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":    token,
		"created":  data.Created_at,
		"modified": data.Modified_at,
		"message":  "Success",
	})
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/back-go-barbershop-core/backend/models/user"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u user.User
		err := ctx.ShouldBindJSON(&u)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = user.Register(u, db)
		if err != nil {
			log.Fatal(err)
			ctx.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
			"status":  200,
		})
	}
}

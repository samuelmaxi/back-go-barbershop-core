package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelmaxi/back-go-barbershop-core/backend/utils/user"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var u user.User
		fmt.Println("PASSOU PELA API")
		err := ctx.ShouldBindJSON(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("PASSOU PELA API")
		u.Register(db)
	}
}

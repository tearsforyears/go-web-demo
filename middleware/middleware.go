package middleware

import (
	"github.com/gin-gonic/gin"
)

func Baned() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(500,"you baned")
		ctx.Abort()
	}
}

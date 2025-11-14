package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AdminLoggin() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		log.Println("Loog middleware runing ")
		ctx.Next()
	}
}

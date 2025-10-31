package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// rencanyanya semua middleware di simpan di sini aja

func UserLoggin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Loog middleware runing ")
		ctx.Next()
	}
}
func AdminLoggin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Loog middleware runing ")
		ctx.Next()
	}
}

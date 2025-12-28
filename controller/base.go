package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// / halaman utaman app
func Dashboard(context *gin.Context) {

	context.HTML(http.StatusOK, "base/dashboard.html",nil)
}

// / handle untuk about atau block aplikasi
func Block(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base/block.html", nil)
}

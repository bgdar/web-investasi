package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/// route utaman halaman adamin
func Admin(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"admin/admin.html",nil)
	
}

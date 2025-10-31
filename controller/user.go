package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(context *gin.Context) {
	context.HTML(http.StatusOK,"user/profile.html",nil)
}

func Sign_in(context *gin.Context) {

	context.HTML(http.StatusOK, "user/sign-in.html", nil)

}
func Sign_up(context *gin.Context) {

	context.HTML(http.StatusOK, "user/sign-up.html", nil)

}

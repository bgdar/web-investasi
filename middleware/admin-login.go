package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ClaimsAdmin struct {
	Adminname string `json:"adminname"`
	jwt.RegisteredClaims
}

func AdminLoggin(jwtKey []byte) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// pastikan di simpan di login admin
	tokenString, err := ctx.Cookie("admin-token")

		if err != nil {
			ctx.Redirect(http.StatusFound, "/admin/sign-in")
			ctx.Abort()
			return
		}
		claims := &ClaimsAdmin{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.Redirect(http.StatusFound, "/admin/sign-in")
			ctx.Abort()
			return
		}

		// Simpan adminname dari token ke context agar bisa digunakan di handler berikutnya
		ctx.Set("adminname", claims.Adminname)
		ctx.Next()
	}
}

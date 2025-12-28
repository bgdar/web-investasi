package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Struktur klaim token
type ClaimsUser struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func UserLoggin(jwtKey []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ambil token yang di kirim lewat header | sudha menggunakn Cookie |
		// tokenString := ctx.GetHeader("Authorization")
		// if tokenString == "" {
		// 	// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token","errro":"error sementara"})
		// 	ctx.Redirect(http.StatusFound, "/sign-in")
		// 	ctx.Abort()
		// 	return
		// }
		// if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		// 	tokenString = tokenString[7:] // ambil setelah "Bearer
		// }

		tokenString, err := ctx.Cookie("user-token")

		if err != nil {
			ctx.Redirect(http.StatusFound, "/user/sign-in")
			ctx.Abort()
			return
		}
		claims := &ClaimsUser{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		// Jika parsing gagal atau token tidak valid
		if err != nil || !token.Valid {
			// ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Redirect(http.StatusFound, "/user/sign-in")
			ctx.Abort()
			return
		}

		// Simpan username dari token ke context agar bisa digunakan di handler berikutnya
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}

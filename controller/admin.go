package controller

import (
	"log"
	"net/http"
	"time"
	"web-investasi/middleware"
	"web-investasi/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

/// route utaman halaman adamin
func AdminHome(ctx *gin.Context)  {

	ctx.HTML(http.StatusOK,"admin/admin.html",gin.H{
		"title":"admin",
	})
	
}

func AdminProfile(ctx *gin.Context) { 
ctx.HTML(http.StatusOK, "admin/profile.hml", gin.H{
"title" : "profile",

	})
}

func AdminSignInPage(ctx *gin.Context) { 
ctx.HTML(http.StatusOK , "admin/sign-in.html",nil)
}
func AdminSignInPost(jwtKey []byte) gin.HandlerFunc { 
	return func(ctx *gin.Context){
			name := ctx.PostForm("username")
		password := ctx.PostForm("password")

	isAdmin , err := model.IsAdminExist(name,password) 
	if err != nil {
		log.Println("[x] ada masalah saat pengecekan user ")
		return
	}
			if !isAdmin {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "anda belum sign-up , silakahn buat akun baru",
			})
			return
		}

	
		// waktu expired token JWT
		expiredTime := time.Now().Add(1 * time.Hour)

		// isi claim token JWT
		claims := &middleware.ClaimsAdmin{
			Adminname :  name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiredTime),
			},
		}

		// buat token baru | Untuk HS256, key adalah string biasa / []byte
		tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// tanda tangani
		tokenString, err := tokenJWT.SignedString(jwtKey)
		if err != nil {
			log.Printf("SignedString error: %T %v\n", jwtKey, err) // tunjukkan tipe jwtKey juga
			ctx.JSON(http.StatusInternalServerError, gin.H{

				"error": "Tidak bisa membuat token utuk admin",
			})
			return
		}

		// hitung maxAge cookie
		maxAge := int(time.Until(expiredTime).Seconds())

		// set cookie token dan username
		ctx.SetCookie("admin-token", tokenString, maxAge, "/", "", false, true)
		ctx.SetCookie("user", name, maxAge, "/", "", false, true)

		ctx.Redirect(http.StatusFound, "/admin/")
	}
}
func AdminSignUpPage(ctx *gin.Context) { 
	ctx.HTML(http.StatusOK,"admin/sign-up.html",nil)

}
func AdminSignUpPost(ctx *gin.Context) { 

}

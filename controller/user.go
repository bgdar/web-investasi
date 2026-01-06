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

func UserHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/home.html", nil)
}

func UserProfile(context *gin.Context) {

	username, err := context.Cookie("username")

	if err != nil {
		log.Println("[x] Tidak mendapakan nama user di cookie, err : ", err)
	}

	user, err := model.GetUserByName(username)
	if err != nil {
		log.Printf("[x] gagal mendapatkan user dari database , file controller/user : %v", err)
	}

	context.HTML(http.StatusOK, "user/profile.html", gin.H{
		"title":    "profile",
		"username": user.Name,
		"email":    user.Email,
		"saldo":    user.Saldo,
	})
}

func UserSignInPage(context *gin.Context) {
	context.HTML(http.StatusOK, "user/sign-in.html", nil)
}
func UserSignIpPost(jwtKey []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		name := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// cek apakah user sudah ada
		isUser, err := model.IsUserExists(name, password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Kesalahan saat pengecekan user",
			})
			return
		}

		if !isUser {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "anda belum sign-up , silakahn buat akun baru",
			})
			return
		}

		// waktu expired token JWT
		expiredTime := time.Now().Add(1 * time.Hour)

		// isi claim token JWT
		claims := &middleware.ClaimsUser{
			Username: name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiredTime),
			},
		}

		// buat token baru | Untuk HS256, key adalah string biasa / []byte
		tokenJWT := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

		// tanda tangani
		tokenString, err := tokenJWT.SignedString(jwtKey)
		if err != nil {
			log.Printf("SignedString error: %T %v\n", jwtKey, err) // tunjukkan tipe jwtKey juga
			ctx.JSON(http.StatusInternalServerError, gin.H{

				"error": "Tidak bisa membuat token",
			})
			return
		}

		// hitung maxAge cookie
		maxAge := int(time.Until(expiredTime).Seconds())

		// set cookie token dan username
		ctx.SetCookie("user-token", tokenString, maxAge, "/", "", false, true)
		ctx.SetCookie("user", name, maxAge, "/", "", false, true)

		ctx.Redirect(http.StatusFound, "/user/")

	}
}

func UserSignUpPage(context *gin.Context) {
	context.HTML(http.StatusOK, "user/sign-up.html", nil)

}
func UserSignUpPost(ctx *gin.Context) {
	// panggil atau buat database atas nama user saat pertama kali reqiester

	name := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	log.Println("data user sign-up",name , email , password)

	isUser, err := model.IsUserExists(name, password)
	if err != nil {
		log.Println("[x] ada yang sasalh saat pengecekan user :", err)
	}
	if isUser {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "user sudah ada silahkan login",
		})
		return

	}

	// // buat table atas nama user ini ||  SUDA  TIDAK  di perluakn lagi
	// model.CreateProductForUser(name)
	// // simpan adata nya
	// if err := model.AddUser(name, password, email); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "ada yang salah saat menambah user baru",
	// 	})
	// 	return
	// }

}

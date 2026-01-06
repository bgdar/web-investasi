package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"web-investasi/model"

	"github.com/gin-gonic/gin"
)

func ProductForUserGet(ctx *gin.Context) {

	table_name, err := ctx.Cookie("user")
	if err != nil {
		log.Println("[x] gagal medapatkan nama user")
		return
	}

	userProduct, err := model.GetAllProduct(table_name)
	if err != nil {
		log.Println("[x] gagal mendapatka product untuk user")
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "product/product.html", gin.H{
		"user_product": userProduct,
	})

}
// / update product untuk usernya
func ProductForUserPost(ctx *gin.Context) {

	var username, errName = ctx.Cookie("username")
	if errName != nil {

		log.Println("gagal mendapatkan username")
	}
	type DataType struct {
		Name         string `json:"name"`
		TotalProduct int16  `json:"total_product"`
	}
	var data DataType
	log.Println("Dta yang di kirim ,:", data)
	var err = ctx.ShouldBindBodyWithJSON(&data)
	if err != nil {

		// nantik redirec dengan Flask message
		// kirm jsn untuk sekaranng
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//user di simpan di cookies cookies
	// user, err := ctx.Cookie("user")
	// if err != nil {
	// 	log.Println("[x] tidak di dapatkan user ", err)
	// }

	// simpan ke table product user  | sudha beda teknik 
	// model.Add(fmt.Sprintf("table_%s", user), data.Name, data.TotalProduct, time.Now(), username)

	// simpan product baru atas nama user 
	model.AddProduct(data.Name,data.TotalProduct,time.Now(),username)
}

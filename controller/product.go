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

// / tampilkan sebagai halaman utama apliaksi
func Product(cxt *gin.Context) {

	var products, err = model.GetAllProdct()
	if err != nil {
		log.Printf("[x] ada yg salah di model product saat mendapatkan semua product : %v", err)
	}
	// log.Println("produtcs : ", products)
	cxt.HTML(http.StatusOK, "product/product.html", gin.H{
		"products":    products,
		"curent_path": cxt.Request.URL.Path,
	})
}

// / function untuk mengupdate product yang di pilih user
func ProductGet(ctx *gin.Context) {
	id := ctx.Param("id")
	idconver, _ := strconv.ParseInt(id, 10, 64)
	product, err := model.GetProductByID(idconver)

	if err != nil {
		log.Println("[x] error dapatkan id ", err)
		return
	}

	ctx.HTML(http.StatusOK, "product/product-id.html", gin.H{
		"title ":              fmt.Sprintf(" product | %s", product.Name),
		"product_id":          product.ID,
		"product_name":        product.Name,
		"product_idr":         product.IDR,
		"product_description": product.Description,
	})

}

// / update product untuk usernya
func ProductAddForUser(ctx *gin.Context) {

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
	user, err := ctx.Cookie("user")
	if err != nil {
		log.Println("[x] tidak di dapatkan user ", err)
	}

	// simpan ke table product user
	model.AddProductForUser(fmt.Sprintf("table_%s", user), data.Name, data.TotalProduct, time.Now())
}

func ProductForUser(ctx *gin.Context) {

	table_name, err := ctx.Cookie("user")
	if err != nil {
		log.Println("[x] gagal medapatkan nama user")
		return
	}

	userProduct, err := model.GetAllProductForUser(table_name)
	if err != nil {
		log.Panicln("[x] gagal mendapatka product untuk user")
		return
	}
	ctx.HTML(http.StatusOK, "product/product-user.html", gin.H{
		"user_product": userProduct,
	})

}

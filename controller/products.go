package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"web-investasi/model"

	"github.com/gin-gonic/gin"
)

// / tampilkan sebagai halaman utama apliaksi
func ProductsGet(cxt *gin.Context) {

	var products, err = model.GetAllProdcts()
	if err != nil {
		log.Printf("[x] ada yg salah di model product saat mendapatkan semua product : %v", err)
	}
	// log.Println("produtcs : ", products)
	cxt.HTML(http.StatusOK, "products/products.html", gin.H{
		"products":    products,
		"curent_path": cxt.Request.URL.Path,
	})
}

// / function untuk mengupdate product yang di pilih user
func ProductsIdGet(ctx *gin.Context) {
	id := ctx.Param("id")
	idconver, _ := strconv.ParseInt(id, 10, 64)
	product, err := model.GetProductsByID(idconver)

	if err != nil {
		log.Println("[x] error dapatkan id ", err)
		return
	}

	ctx.HTML(http.StatusOK, "products/products-id.html", gin.H{
		"title ":              fmt.Sprintf(" product | %s", product.Name),
		"product_id":          product.ID,
		"product_name":        product.Name,
		"product_idr":         product.IDR,
		"product_description": product.Description,
	})
}

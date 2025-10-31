package controller

import (
	"log"
	"net/http"
	"strconv"
	"web-investasi/model"

	"github.com/gin-gonic/gin"
)

/// tampilkan sebagai halaman utama apliaksi
func Product(cxt *gin.Context) {

	var data = map[string]string{
		"img" : ".",
	}
cxt.HTML(http.StatusOK,"product/product.html",data)
}
 

/// function untuk mengupdate product yang di pilih user
func Product_get(ctx *gin.Context) {

	id := ctx.Param("id")

	idconver , _ := strconv.ParseInt(id,10,64)
	product , err := model.GetProductId(idconver)

	if err != nil {
		log.Println("[x] error dapatkan id ",err)
	}

	log.Print("id : ",id)
	
	ctx.HTML(http.StatusOK,"product/product-id.html",gin.H{
		"title " : "id-product",
		"product_name" : product.Name,
		"product_idr" : product.IDR,
		"product_description" : product.Description,
	})
}

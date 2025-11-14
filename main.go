package main

import (
	"web-investasi/config"
	"web-investasi/controller"
	"web-investasi/middleware"

	// "web-investasi/seeder"

	"github.com/gin-gonic/gin"
)

func main() {

	var jwtKey = []byte("ini_secret_key2")

	var route = gin.Default()

	route.Static("/static/", "./static/")
	route.LoadHTMLGlob("views/**/*")

	// koneksi ke database
	config.ConnectDatabase()

	//<-----JALANKAN DATA DUMMY DALAM ARTIAN DI BUTUHKAN ---->
	// seeder.StartCreateTable()
	// seeder.FillDataTable()

	base_route := route.Group("/")
	{
		base_route.GET("/", controller.Dashboard)
		base_route.GET("/block", controller.Block)
	}

	product_route := route.Group("/product")
	product_route.Use(middleware.UserLoggin(jwtKey))
	{
		product_route.GET("/", controller.Product)
		product_route.GET("/:id", controller.ProductGet) // kirim id
		product_route.POST("/product-user", controller.ProductAddForUser)
		product_route.GET("/product-user", controller.ProductForUser)
	}

	user_route := route.Group("/user")
	{
		user_route.GET("/", middleware.UserLoggin(jwtKey), controller.Profile)
		user_route.GET("/sign-in", controller.SignInPage)
		user_route.POST("/sign-in", controller.SignIpPost(jwtKey))
		user_route.GET("/sign-up", controller.SignUpPage)
		user_route.POST("/sign-up", controller.SignUpPost)
	}

	admin_route := route.Group("/admin")
	{
		admin_route.GET("/", controller.Admin)
	}

	// jalana loalhost:5002
	route.Run(":5002")

}

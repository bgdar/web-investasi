package main

import (
	"web-investasi/config"
	"web-investasi/controller"
	// "web-investasi/seeder"

	"github.com/gin-gonic/gin"
)

func main() {

	var route = gin.Default()

	// route.Static("/static/","./static/") // matiin dulu untuk file static
	route.LoadHTMLGlob("views/**/*")

	// koneksi ke database 
	config.ConnectDatabase()
	// buat table user pertama kali  
	// seeder.StartCreateTable()


	base_route := route.Group("/")
	{
		base_route.GET("/", controller.Dashboard)
	}

	
	product_route := route.Group("/product")
	{ 
		product_route.GET("/",controller.Product)
		product_route.GET("/:id",controller.Product_get) // kirim id
	}


	user_route := route.Group("/user")
	{
		user_route.GET("/",controller.Profile)
		user_route.GET("/sign-in", controller.Sign_in)
		user_route.GET("/sign-un", controller.Sign_up)
	}

	// jalana loalhost:5002
	route.Run(":5002")

}

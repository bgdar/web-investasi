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
	product_route.Use(middleware.UserLoggin(jwtKey)) // UserLogginsemua menggunaakn middleware
	{
		product_route.GET("/", controller.ProductGet)
		product_route.GET("/:id", controller.ProductIdGet) // kirim id
		product_route.GET("/product-user", controller.ProductForUserGet)
		product_route.POST("/product-user", controller.ProductForUserPost)
	}

	user_route := route.Group("/user")
	{
		user_route.GET("/", middleware.UserLoggin(jwtKey), controller.UserHome)
		user_route.GET("/profile", middleware.UserLoggin(jwtKey), controller.UserProfile)

		user_route.GET("/sign-in", controller.UserSignInPage)
		user_route.POST("/sign-in", controller.UserSignIpPost(jwtKey))
		user_route.GET("/sign-up", controller.UserSignUpPage)
		user_route.POST("/sign-up", controller.UserSignUpPost)
	}

	admin_route := route.Group("/admin")
	{
		admin_route.GET("/", middleware.AdminLoggin(jwtKey), controller.AdminHome)
		admin_route.GET("/profile", middleware.AdminLoggin(jwtKey), controller.AdminProfile)

		admin_route.GET("/sign-in", controller.AdminSignInPage)
		admin_route.POST("/sign-in", controller.AdminSignInPost(jwtKey))
		admin_route.GET("/sign-up", controller.AdminSignUpPage)
		admin_route.POST("/sign-up", controller.AdminSignUpPost)
	}

	// jalana loalhost:5002
	route.Run(":5002")

}

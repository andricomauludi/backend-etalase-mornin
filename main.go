package main

import (
	"github.com/andricomauludi/backend-etalase-mornin/controllers/authcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/controllers/productcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/initializers"
	"github.com/andricomauludi/backend-etalase-mornin/middleware"
	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	models.ConnectDatabase() //melakukan koneksi database

	// //URL CONTROLLER DAN FUNCTIONNYA

	//AUTH

	//1 superadmin
	//2 admin
	//3 cashier
	//4 dashboard management
	//5 customer
	api := r.Group("/api")

	api.Use(middleware.RequireAuth)

	product := api.Group("/product")
	auth := api.Group("/auth")

	product.Use(middleware.Authorization([]int{1, 2, 4}))
	auth.Use(middleware.Authorization([]int{1, 2}))

	//CRUD Product ((RBAC and customer))
	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/sandwich", productcontroller.Show_sandwich)
	r.GET("/api/product/rice", productcontroller.Show_rice)
	r.GET("/api/product/coffee", productcontroller.Show_coffee)
	r.GET("/api/product/:id", productcontroller.Show) //terdapat id yg params nya dapat diambil oleh controller
	product.POST("/", productcontroller.Create)
	product.PUT("/:id", productcontroller.Update)
	product.DELETE("/", productcontroller.Delete)

	r.POST("api/auth/signup", authcontroller.Signup)
	r.POST("api/auth/login", authcontroller.Login)
	auth.GET("/validate", authcontroller.Validate)
	auth.GET("/showall", middleware.RequireAuth, authcontroller.Showall)
	// r := gin.Default()
	// r.POST("/login", handler.LoginHandler)

	// api:=r.Group("/api")

	// api.Use(middleware.ValidateToken())

	// product:=api.Group("/product")

	// product.Use(middleware.Authorization([]int{1}))

	// product.GET("/",handler.GetAll)
	// product.POST("/",middleware.Authorization([]int{4}), handler.AddProduct)

	// user := api.Group("/User")
	// user.GET("/",func(c *gin.Context) {
	// 	c.AbortWithStatusJSON(200, gin.H{
	// 		"status":"ok",
	// 	})
	// })
	// r.Run("localhost:8080")

	r.Run() //WAJIB ADA agar controller terjalani
	// log.Fatal(http.ListenAndServe(":8080", l))
}

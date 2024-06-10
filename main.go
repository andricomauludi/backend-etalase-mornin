package main

import (
	"github.com/andricomauludi/backend-etalase-mornin/controllers/authcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/controllers/productcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/controllers/transactioncontroller"
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
	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	// api.Use(middleware.RequireAuth)

	product := api.Group("/product")
	transaction := api.Group("/transaction")
	auth := api.Group("/auth")

	// product.Use(middleware.Authorization([]int{1, 2, 4}))
	// auth.Use(middleware.Authorization([]int{1, 2}))

	//CRUD Product ((RBAC and customer))
	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/makanan", productcontroller.Show_makanan)
	r.GET("/api/product/minuman", productcontroller.Show_minuman)
	r.GET("/api/product/cemilan", productcontroller.Show_cemilan)
	r.GET("/api/product/barbershop", productcontroller.Show_barbershop)
	r.GET("/api/product/:id", productcontroller.Show) //terdapat id yg params nya dapat diambil oleh controller
	product.POST("/create", productcontroller.Create)
	// product.POST("/base64convert", productcontroller.Base64converter)
	product.PUT("/:id", productcontroller.Update)
	product.DELETE("/", productcontroller.Delete)

	//TRANSACTION
	transaction.POST("/create_bill", transactioncontroller.CreateOrUpdateBill)
	transaction.POST("/create_detail_bill", transactioncontroller.Create_detail_bill)
	transaction.POST("/create_detail_bill_json", transactioncontroller.Create_detail_bill2)
	transaction.POST("/create_klien", transactioncontroller.Create_klien)
	transaction.GET("/show_transaction", transactioncontroller.Show_transaction)
	transaction.GET("/show_detail_bill/:id", transactioncontroller.Show_detail_bill)
	transaction.GET("/show_saved_bill", transactioncontroller.Show_saved_bill)
	transaction.PUT("/edit_bill/:id", transactioncontroller.Update_bill)
	transaction.PUT("/edit_detail_bill/:id", transactioncontroller.Update_detail_bill)
	transaction.PUT("/edit_klien/:id", transactioncontroller.Update_klien)
	transaction.POST("/delete_bill", transactioncontroller.Delete_bill)
	transaction.POST("/delete_detail_bill", transactioncontroller.Delete_detail_bill)
	transaction.DELETE("/delete_klien", transactioncontroller.Delete_klien)

	r.POST("api/auth/signup", authcontroller.Signup)
	r.POST("api/auth/login", authcontroller.Login)
	auth.POST("/logout", middleware.RequireAuth, authcontroller.Logout)
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

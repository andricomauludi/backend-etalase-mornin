package main

import (
	"github.com/andricomauludi/backend-etalase-mornin/controllers/authcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/controllers/pendapatancontroller"
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
	pendapatan := api.Group("/pendapatan")
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
	transaction.POST("/create_pengeluaran", transactioncontroller.Create_pengeluaran)
	transaction.POST("/create_klien", transactioncontroller.Create_klien)
	transaction.POST("/excel_export", transactioncontroller.Excel_export)
	transaction.GET("/show_transaction", transactioncontroller.Show_transaction)
	transaction.GET("/show_transaction_cvj", transactioncontroller.Show_transaction_cvj)
	transaction.GET("/detail_transaction/:id", transactioncontroller.Detail_transaction)
	transaction.GET("/show_bill/:id", transactioncontroller.Show_bill)
	transaction.GET("/show_pengeluaran", transactioncontroller.Show_pengeluaran)
	transaction.GET("/show_pengeluaran_cvj", transactioncontroller.Show_pengeluaran_cvj)
	transaction.GET("/show_detail_bill/:id", transactioncontroller.Show_detail_bill)
	transaction.GET("/show_saved_bill", transactioncontroller.Show_saved_bill)
	transaction.PUT("/edit_bill/:id", transactioncontroller.Update_bill)
	transaction.PUT("/edit_pengeluaran/:id", transactioncontroller.Update_pengeluaran)
	transaction.PUT("/edit_detail_bill/:id", transactioncontroller.Update_detail_bill)
	transaction.PUT("/edit_klien/:id", transactioncontroller.Update_klien)
	transaction.POST("/delete_bill", transactioncontroller.Delete_bill)
	transaction.POST("/delete_detail_bill", transactioncontroller.Delete_detail_bill)
	transaction.DELETE("/delete_klien", transactioncontroller.Delete_klien)
	transaction.POST("/delete_pengeluaran", transactioncontroller.Delete_pengeluaran)

	r.POST("api/auth/signup", authcontroller.Signup)
	r.POST("api/auth/login", authcontroller.Login)
	r.POST("api/auth/check-auth", authcontroller.CheckAuthHandler)

	auth.POST("/logout", middleware.RequireAuth, authcontroller.Logout)
	auth.GET("/validate", authcontroller.Validate)
	auth.GET("/showall", middleware.RequireAuth, authcontroller.Showall)

	pendapatan.GET("/show_pendapatan_bulanan", pendapatancontroller.TotalCurrentMonth)
	pendapatan.POST("/show_pendapatan_bulanan_pembayaran", pendapatancontroller.TotalCurrentMonthJenisPembayaran)
	pendapatan.GET("/show_pendapatan_bulanan_cvj", pendapatancontroller.TotalCurrentMonthCvj)
	pendapatan.GET("/show_pendapatan_harian", pendapatancontroller.TotalToday)
	pendapatan.POST("/show_pendapatan_harian_pembayaran", pendapatancontroller.TotalTodayJenisPembayaran)
	pendapatan.GET("/show_pendapatan_harian_cvj", pendapatancontroller.TotalTodayCvj)
	pendapatan.GET("/show_pengeluaran_bulanan", pendapatancontroller.TotalPengeluaranCurrentMonth)
	pendapatan.POST("/show_pengeluaran_bulanan_jenis", pendapatancontroller.TotalPengeluaranCurrentMonthJenis)
	pendapatan.GET("/show_pengeluaran_bulanan_cvj", pendapatancontroller.TotalPengeluaranCurrentMonthCvj)
	pendapatan.GET("/show_pengeluaran_harian", pendapatancontroller.TotalPengeluaranToday)
	pendapatan.GET("/show_pengeluaran_harian_cvj", pendapatancontroller.TotalPengeluaranTodayCvj)
	pendapatan.GET("/show_keuntungan_bulanan", pendapatancontroller.TotalKeuntunganBersihCurrentMonth)
	pendapatan.GET("/show_keuntungan_bulanan_cvj", pendapatancontroller.TotalKeuntunganBersihCurrentMonthCvj)
	pendapatan.GET("/show_keuntungan_harian", pendapatancontroller.TotalKeuntunganBersihCurrentDay)
	pendapatan.GET("/show_keuntungan_harian_cvj", pendapatancontroller.TotalKeuntunganBersihCurrentDayCvj)

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

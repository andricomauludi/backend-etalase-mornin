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
	r.POST("api/auth/signup", authcontroller.Signup)
	r.POST("api/auth/login", authcontroller.Login)
	r.GET("api/auth/validate", middleware.RequireAuth, authcontroller.Validate)

	//CRUD Product
	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show) //terdapat id yg params nya dapat diambil oleh controller
	r.POST("api/products", productcontroller.Create)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products", productcontroller.Delete)

	r.Run() //WAJIB ADA agar controller terjalani
	// log.Fatal(http.ListenAndServe(":8080", l))
}

package main

import (
	"github.com/andricomauludi/backend-etalase-mornin/controllers/productcontroller"
	"github.com/andricomauludi/backend-etalase-mornin/models"
<<<<<<< HEAD
=======

>>>>>>> development
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
<<<<<<< HEAD
	models.ConnectDatabase()

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show)
	r.POST("api/products", productcontroller.Create)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products/:id", productcontroller.Delete)

	r.Run()
=======
	models.ConnectDatabase() //melakukan koneksi database

	//URL CONTROLLER DAN FUNCTIONNYA

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show) //terdapat id yg params nya dapat diambil oleh controller
	r.POST("api/products", productcontroller.Create)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products", productcontroller.Delete)

	r.Run() //WAJIB ADA agar controller terjalani
>>>>>>> development
}

package main

import (
	"github.com/andricomauludi/backend-etalase-mornin/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// l := mux.NewRouter()

	// models.ConnectDatabase() //melakukan koneksi database

	// //URL CONTROLLER DAN FUNCTIONNYA

	// l.HandleFunc("login", authcontroller.Login).Methods("POST")
	// l.HandleFunc("register", authcontroller.Register).Methods("POST")
	// l.HandleFunc("logout", authcontroller.Logout).Methods("POST")

	// log.Fatal(http.ListenAndServe(":8080", l))

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show) //terdapat id yg params nya dapat diambil oleh controller
	r.POST("api/products", productcontroller.Create)
	r.PUT("api/products/:id", productcontroller.Update)
	r.DELETE("api/products", productcontroller.Delete)

	r.Run() //WAJIB ADA agar controller terjalani
}

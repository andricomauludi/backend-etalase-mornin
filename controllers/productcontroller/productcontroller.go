package productcontroller

import (
<<<<<<< HEAD
	"net/http"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product //ambl model product

	models.DB.Find(&products) //ambil semua data di product
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func Show(c *gin.Context) {
	var product models.Product //ambil model product
	id := c.Param("id")        // params diambil dari url main.go

	if err := models.DB.First(&product, id).Error; err != nil { //mengecek saat pengambilan satu data pada product berdasarkan id, dan pengecekan apabila terdapat error
		switch err {
		case gorm.ErrRecordNotFound: //apabila tidak terdapat error record saat eksekusi error
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		default: //terdapat error dan terdapat error record nya
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
=======
	"encoding/json"
	"net/http"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Product //array dan ambil model product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products}) //untuk return json nya
}
func Show(c *gin.Context) {
	var product models.Product //ambil model product
	id := c.Param("id")        //ngambil params dari URL main.go

	if err := models.DB.First(&product, id).Error; err != nil { //mencari 1 data yg memiliki id yg sama dengan yg dicari, apabila tidak dapat maka masuk ke if ini(error)
		switch err {
		case gorm.ErrRecordNotFound: //apabila tidak terdapat error record
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "ERROR Data not found"})
			return
		default: //apabilla terdapat error record, mengembalikan message dengan error record
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error})
>>>>>>> development
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Create(c *gin.Context) {

	var product models.Product

<<<<<<< HEAD
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error})
=======
	if err := c.ShouldBindJSON(&product); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
>>>>>>> development
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
<<<<<<< HEAD

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

=======
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&product); err != nil { //mengecek apakah product ada atau tidak (input menggunakan JSON)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data is not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Data updated"}) //mengembalikan status yg benar
}
func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	// input := map[string]string{"id": "0"} //mengambil input tipe map dengan bentuk string dari json

	if err := c.ShouldBindJSON(&input); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	// id, _ := strconv.ParseInt(input["id"], 10, 64) //melakukan perubahan string to integer, dengan ukuran integer 10 dan size integer 64
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
>>>>>>> development
}

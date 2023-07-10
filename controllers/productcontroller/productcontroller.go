package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product //array dan ambil model product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_sandwich(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "menu_type = ?", "sandwich")
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_rice(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "menu_type = ?", "rice")
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_coffee(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "menu_type = ?", "coffee")
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show(c *gin.Context) {
	var product models.Product //ambil model product
	id := c.Param("id")        //ngambil params dari URL main.go

	if err := models.DB.First(&product, id).Error; err != nil { //mencari 1 data yg memiliki id yg sama dengan yg dicari, apabila tidak dapat maka masuk ke if ini(error)
		switch err {
		case gorm.ErrRecordNotFound: //apabila tidak terdapat error record
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "data": "ERROR data not found"})
			return
		default: //apabilla terdapat error record, mengembalikan data dengan error record
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": err.Error})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": product})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Item created successfully!"})
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&product); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Data is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data updated"}) //mengembalikan status yg benar
}
func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}
	// id, _ := strconv.ParseInt(input["id"], 10, 64) //melakukan perubahan string to integer, dengan ukuran integer 10 dan size integer 64

	if err := c.ShouldBindJSON(&input); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Product is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}

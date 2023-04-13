package productcontroller

import (
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
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error})
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}

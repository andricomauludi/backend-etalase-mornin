package productcontroller

import (
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

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "BAD REQUEST"})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}

package productcontroller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product //array dan ambil model product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_makanan(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "jenis_menu = ?", "makanan")
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_cemilan(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "jenis_menu = ?", "cemilan")
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products}) //untuk return json nya
}
func Show_minuman(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	models.DB.Find(&products, "jenis_menu = ?", "minuman")
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

func CreateDummy(c *gin.Context) {

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Access form values
	formData := make(map[string]string)
	for key, values := range c.Request.MultipartForm.Value {
		if len(values) > 0 {
			formData[key] = values[0]
		}
	}

	// Process form data
	// Here you can do whatever you need with the form data
	// For example, you can save it to a database, perform some validation, etc.

	// Access uploaded files
	files := c.Request.MultipartForm.File["files"]
	for _, file := range files {
		// Generate a unique filename with timestamp
		timestamp := time.Now().UnixNano()
		filename := file.Filename
		ext := filepath.Ext(filename)
		filename = filename[:len(filename)-len(ext)] + "_" + strconv.FormatInt(timestamp, 10) + ext

		// Specify the destination directory
		dest := filepath.Join("assets/photo/products", filename)

		// Save the uploaded file to the destination directory
		if err := c.SaveUploadedFile(file, dest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save file"})
			return
		}
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Form submitted successfullyrrr", "data": formData, "File": files})
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

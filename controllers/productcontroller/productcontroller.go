package productcontroller

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	// Retrieve the products from the database
	if err := models.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		products[i].Photo = base64String
		if err != nil {
			c.Next()

		}

		// base64String := base64.StdEncoding.EncodeToString([]byte("assets/photo/products/"+product.Photo))
	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products})
}
func ConvertFileToBase64(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read file content into a byte slice
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convert file content to base64
	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	return base64String, nil
}
func Show_makanan(c *gin.Context) {

	var products []models.Product //array dan ambil model product

	if err := models.DB.Find(&products, "jenis_menu = ?", "makanan").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	var base64Strings []string
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": "error on base 64", "base64": base64Strings, "data": products})

		}
		products[i].Photo = base64String

		// base64String := base64.StdEncoding.EncodeToString([]byte("assets/photo/products/"+product.Photo))
	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products})
}
func Show_barbershop(c *gin.Context) {

	var products []models.Product //array dan ambil model product

	if err := models.DB.Find(&products, "jenis_menu = ?", "barbershop").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	var base64Strings []string
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": "error on base 64", "base64": base64Strings, "data": products})

		}
		products[i].Photo = base64String

		// base64String := base64.StdEncoding.EncodeToString([]byte("assets/photo/products/"+product.Photo))
	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products})
}
func Show_cemilan(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	if err := models.DB.Find(&products, "jenis_menu = ?", "cemilan").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	var base64Strings []string
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": "error on base 64", "base64": base64Strings, "data": products})

		}
		products[i].Photo = base64String

		// base64String := base64.StdEncoding.EncodeToString([]byte("assets/photo/products/"+product.Photo))
	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products})
}
func Show_minuman(c *gin.Context) {

	var products []models.Product //array dan ambil model product
	if err := models.DB.Find(&products, "jenis_menu = ?", "minuman").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	var base64Strings []string
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": "error on base 64", "base64": base64Strings, "data": products})

		}
		products[i].Photo = base64String

		// base64String := base64.StdEncoding.EncodeToString([]byte("assets/photo/products/"+product.Photo))
	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": products})
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

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	// Access the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "No file uploaded or invalid file field name"})
		return
	}

	// Generate a unique filename with timestamp
	timestamp := time.Now().UnixNano()
	filename := file.Filename
	ext := filepath.Ext(filename)
	filename = filename[:len(filename)-len(ext)] + "_" + strconv.FormatInt(timestamp, 10) + ext

	// Specify the destination directory
	dest := filepath.Join("assets/photo/products", filename)

	// Save the uploaded file to the destination directory
	if err := c.SaveUploadedFile(file, dest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to save file"})
		return
	}
	NamaMenuPost := c.PostForm("nama_menu")
	JenisMenuPost := c.PostForm("jenis_menu")
	DeskripsiMenuPost := c.PostForm("deskripsi_menu")
	HargaStr := c.PostForm("harga")

	HargaInt, err := strconv.ParseInt(HargaStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jenis_menu value"})
		return
	}
	// Save file details to the database
	product := models.Product{
		Photo:         filename,
		NamaMenu:      NamaMenuPost,
		JenisMenu:     JenisMenuPost,
		DeskripsiMenu: DeskripsiMenuPost,
		Harga:         HargaInt,
		// ContentType: file.Header.Get("Content-Type"),
		// Add other fields from the form data as needed
	}
	// Assuming models.DB is your database connection
	if err := models.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": product, "message": "File uploaded successfully"})
}

func CreateOld(c *gin.Context) {

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

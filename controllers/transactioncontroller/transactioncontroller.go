package transactioncontroller

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
	var base64Strings []string
	for i, product := range products {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		base64String, err := ConvertFileToBase64("assets/photo/products/" + product.Photo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": products, "base64": base64Strings})

		}
		products[i].Photo = base64String

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
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": products, "base64": base64Strings})

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
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": products, "base64": base64Strings})

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
			c.JSON(http.StatusOK, gin.H{"status": -1, "message": products, "base64": base64Strings})

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

func Create_bill(c *gin.Context) {

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	PaidPost := c.PostForm("paid")
	NamaBillPost := c.PostForm("nama_bill")
	TimestampPost := time.Now()
	IdJenisPembayaranPost := c.PostForm("id_jenis_pembayaran")
	JenisPembayaranPost := c.PostForm("jenis_pembayaran")
	TotalPostStr := c.PostForm("total")
	CashInPostStr := c.PostForm("cash_in")
	CashOutPostStr := c.PostForm("cash_out")
	IdKlienPostStr := c.PostForm("id_klien")
	NamaKlienPost := c.PostForm("nama_klien")

	IdKlienPost, err := strconv.ParseInt(IdKlienPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	TotalPost, err := strconv.ParseInt(TotalPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	CashInPost, err := strconv.ParseInt(CashInPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cash_in value"})
		return
	}
	CashOutPost, err := strconv.ParseInt(CashOutPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cash_out value"})
		return
	}
	// Save file details to the database
	bill := models.Bill{
		IdKlien:           IdKlienPost,
		NamaKlien:         NamaKlienPost,
		NamaBill:          NamaBillPost,
		Paid:              PaidPost,
		Timestamp:         TimestampPost,
		IdJenisPembayaran: IdJenisPembayaranPost,
		JenisPembayaran:   JenisPembayaranPost,
		Total:             TotalPost,
		CashIn:            CashInPost,
		CashOut:           CashOutPost,

		// Add other fields from the form data as needed
	}
	// Assuming models.DB is your database connection
	if err := models.DB.Create(&bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": bill, "message": "Your bill is successfully created!"})
}
func Create_detail_bill(c *gin.Context) {

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	IdBillPostStr := c.PostForm("id_bill")
	IdMenuPostStr := c.PostForm("id_menu")
	NamaMenuPost := c.PostForm("nama_menu")
	HargaPostStr := c.PostForm("harga")
	JumlahPostStr := c.PostForm("jumlah")
	TotalHargaPostStr := c.PostForm("total_harga")

	IdBillPost, err := strconv.ParseInt(IdBillPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	IdMenuPost, err := strconv.ParseInt(IdMenuPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	HargaPost, err := strconv.ParseInt(HargaPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	JumlahPost, err := strconv.ParseInt(JumlahPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total value"})
		return
	}
	TotalHargaPost, err := strconv.ParseInt(TotalHargaPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cash_in value"})
		return
	}

	// Save file details to the database
	detail_bill := models.Detail_bill{
		IdBill:     IdBillPost,
		IdMenu:     IdMenuPost,
		NamaMenu:   NamaMenuPost,
		Harga:      HargaPost,
		Jumlah:     JumlahPost,
		TotalHarga: TotalHargaPost,
	}
	// Assuming models.DB is your database connection
	if err := models.DB.Create(&detail_bill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": detail_bill, "message": "Your detail bill is successfully created!"})
}

func Update_detail_bill(c *gin.Context) {
	var detail_bill models.Detail_bill
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&detail_bill); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	if models.DB.Model(&detail_bill).Where("id = ?", id).Updates(&detail_bill).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Data is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data updated"}) //mengembalikan status yg benar
}
func Update_bill(c *gin.Context) {
	var bill models.Bill
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&bill); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	if models.DB.Model(&bill).Where("id = ?", id).Updates(&bill).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Data is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data updated"}) //mengembalikan status yg benar
}
func Delete_bill(c *gin.Context) {
	var bill models.Bill

	var input struct {
		Id json.Number
	}
	// id, _ := strconv.ParseInt(input["id"], 10, 64) //melakukan perubahan string to integer, dengan ukuran integer 10 dan size integer 64

	if err := c.ShouldBindJSON(&input); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&bill, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Bill is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}
func Delete_detail_bill(c *gin.Context) {
	var detail_bill models.Detail_bill

	var input struct {
		Id json.Number
	}
	// id, _ := strconv.ParseInt(input["id"], 10, 64) //melakukan perubahan string to integer, dengan ukuran integer 10 dan size integer 64

	if err := c.ShouldBindJSON(&input); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&detail_bill, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Detail bill is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}

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
func Show_transaction(c *gin.Context) {
	var bills []models.Bill // array to hold all bills
	// UserResponse struct represents the custom JSON response
	type BillResponse struct {
		Bill        models.Bill
		Detail_bill []models.Detail_bill
	}

	// Fetch all bills from the database
	if err := models.DB.Where("tipe = ?", 0).Find(&bills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the bills and fetch their corresponding detail bills
	var billResponses []BillResponse
	for i := range bills {
		var detailBills []models.Detail_bill // array to hold detail bills for each bill
		if err := models.DB.Find(&detailBills, "id_bill = ?", bills[i].Id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := BillResponse{
			Bill:        bills[i],
			Detail_bill: detailBills,
		}
		billResponses = append(billResponses, response)
	}

	// Return the JSON response with bills and their associated detail bills
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": billResponses})
}
func Show_transaction_cvj(c *gin.Context) {
	var bills []models.Bill // array to hold all bills
	// UserResponse struct represents the custom JSON response
	type BillResponse struct {
		Bill        models.Bill
		Detail_bill []models.Detail_bill
	}

	// Fetch all bills from the database
	if err := models.DB.Where("tipe = ?", 1).Find(&bills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the bills and fetch their corresponding detail bills
	var billResponses []BillResponse
	for i := range bills {
		var detailBills []models.Detail_bill // array to hold detail bills for each bill
		if err := models.DB.Find(&detailBills, "id_bill = ?", bills[i].Id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := BillResponse{
			Bill:        bills[i],
			Detail_bill: detailBills,
		}
		billResponses = append(billResponses, response)
	}

	// Return the JSON response with bills and their associated detail bills
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": billResponses})
}
func Detail_transaction(c *gin.Context) {
	id := c.Param("id")
	var bills []models.Bill // array to hold all bills
	// UserResponse struct represents the custom JSON response
	type BillResponse struct {
		Bill        models.Bill
		Detail_bill []models.Detail_bill
	}

	// Fetch all bills from the database
	if err := models.DB.Find(&bills, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the bills and fetch their corresponding detail bills
	var billResponses []BillResponse
	for i := range bills {
		var detailBills []models.Detail_bill // array to hold detail bills for each bill
		if err := models.DB.Find(&detailBills, "id_bill = ?", bills[i].Id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := BillResponse{
			Bill:        bills[i],
			Detail_bill: detailBills,
		}
		billResponses = append(billResponses, response)
	}

	// Return the JSON response with bills and their associated detail bills
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": billResponses})
}
func Show_bill(c *gin.Context) {
	id := c.Param("id")
	var bills []models.Bill // array to hold all bills
	// UserResponse struct represents the custom JSON response

	if err := models.DB.Find(&bills, id).Error; err != nil { //mencari 1 data yg memiliki id yg sama dengan yg dicari, apabila tidak dapat maka masuk ke if ini(error)
		switch err {
		case gorm.ErrRecordNotFound: //apabila tidak terdapat error record
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "data": "ERROR data not found"})
			return
		default: //apabilla terdapat error record, mengembalikan data dengan error record
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": err.Error})
			return
		}
	}

	// Return the JSON response with bills and their associated bills
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": bills})
}
func Show_pengeluaran(c *gin.Context) {

	var pengeluarans []models.Pengeluaran

	// Retrieve the products from the database
	if err := models.DB.Where("tipe = ?", 0).Find(&pengeluarans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": pengeluarans})
}
func Show_pengeluaran_cvj(c *gin.Context) {

	var pengeluarans []models.Pengeluaran

	// Retrieve the products from the database
	if err := models.DB.Where("tipe = ?", 1).Find(&pengeluarans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": pengeluarans})
}

func Show_saved_bill(c *gin.Context) {

	var bill []models.Bill               //array dan ambil model product
	var detail_bill []models.Detail_bill //array dan ambil model product

	// UserResponse struct represents the custom JSON response
	type BillResponse struct {
		Bill        models.Bill
		Detail_bill []models.Detail_bill
	}

	// if err := models.DB.Find(&bill, "jenis_menu = ?", "makanan").Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := models.DB.Find(&bill, "paid = ?", 0).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Loop through the products and convert their desired fields to base64
	var billResponses []BillResponse
	for i, _ := range bill {
		// Assuming you want to convert the product name to base64
		// Adjust this to convert the appropriate field
		if err := models.DB.Find(&detail_bill, "id_bill = ?", bill[i].Id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := BillResponse{
			Bill:        bill[i],
			Detail_bill: detail_bill,
		}
		billResponses = append(billResponses, response)

	}

	// Return the JSON response with products and their base64 encoded fields
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": billResponses})
}
func Show_detail_bill(c *gin.Context) {
	var detail_bill []models.Detail_bill //ambil model product
	id := c.Param("id")                  //ngambil params dari URL main.go

	if err := models.DB.Find(&detail_bill, "id_bill = ?", id).Error; err != nil { //mencari 1 data yg memiliki id yg sama dengan yg dicari, apabila tidak dapat maka masuk ke if ini(error)
		switch err {
		case gorm.ErrRecordNotFound: //apabila tidak terdapat error record
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "data": "ERROR data not found"})
			return
		default: //apabilla terdapat error record, mengembalikan data dengan error record
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": err.Error})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": detail_bill})
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

func CreateOrUpdateBill(c *gin.Context) {
	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	IdStr := c.PostForm("id")
	PaidPost := c.PostForm("paid")
	TipePost := c.PostForm("tipe")
	NamaBillPost := c.PostForm("nama_bill")
	TimestampPost := time.Now()
	JenisPembayaranPost := c.PostForm("jenis_pembayaran")
	TotalPostStr := c.PostForm("total")
	CashInPostStr := c.PostForm("cash_in")
	CashOutPostStr := c.PostForm("cash_out")
	IdKlienPostStr := c.PostForm("id_klien")
	NamaKlienPost := c.PostForm("nama_klien")

	Id, err := strconv.ParseInt(IdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id value"})
		return
	}
	IdKlienPost, err := strconv.ParseInt(IdKlienPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IdKlien value"})
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

	var bill models.Bill
	if Id != 0 {
		// Check if the bill with the given ID exists
		if err := models.DB.Where("id = ?", Id).First(&bill).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Bill with given ID does not exist
				c.JSON(http.StatusNotFound, gin.H{"status": -1, "error": "Bill not found"})
				return
			}
			// Some other error occurred
			c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to query database"})
			return
		}

		// Bill exists, update it
		bill.NamaKlien = NamaKlienPost
		bill.NamaBill = NamaBillPost
		bill.Paid = PaidPost
		bill.Tipe = TipePost
		bill.Timestamp = TimestampPost
		bill.JenisPembayaran = JenisPembayaranPost
		bill.Total = TotalPost
		bill.CashIn = CashInPost
		bill.CashOut = CashOutPost
		bill.IdKlien = IdKlienPost
		bill.NamaKlien = NamaKlienPost

		if err := models.DB.Save(&bill).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to update data in database"})
			return
		}

		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": bill, "message": "Your bill is successfully updated!"})
	} else {
		// Bill does not exist, create a new one
		bill = models.Bill{
			IdKlien:         IdKlienPost,
			NamaKlien:       NamaKlienPost,
			NamaBill:        NamaBillPost,
			Paid:            PaidPost,
			Tipe:            TipePost,
			Timestamp:       TimestampPost,
			JenisPembayaran: JenisPembayaranPost,
			Total:           TotalPost,
			CashIn:          CashInPost,
			CashOut:         CashOutPost,
		}

		if err := models.DB.Create(&bill).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
			return
		}

		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{"status": 1, "data": bill, "message": "Your bill is successfully created!"})
	}
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

func Create_pengeluaran(c *gin.Context) {

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	NamaPengeluaranPost := c.PostForm("nama_pengeluaran")
	JenisPengeluaranPost := c.PostForm("jenis_pengeluaran")
	HargaPengeluaranPostStr := c.PostForm("harga_pengeluaran")
	JumlahBarangPostStr := c.PostForm("jumlah_barang")
	SatuanPost := c.PostForm("satuan")
	TipePost := c.PostForm("tipe")
	TotalPengeluaranPostStr := c.PostForm("total_pengeluaran")
	WaktuPengeluaranPost := time.Now()

	HargaPengeluaranPost, err := strconv.ParseInt(HargaPengeluaranPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Harga Pengeluaran value"})
		return
	}
	JumlahBarangPost, err := strconv.ParseInt(JumlahBarangPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Jumlah Barang value"})
		return
	}
	TotalPengeluaranPost, err := strconv.ParseInt(TotalPengeluaranPostStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Total Pengeluaran value"})
		return
	}

	// Save file details to the database
	pengeluaran := models.Pengeluaran{
		NamaPengeluaran:  NamaPengeluaranPost,
		JenisPengeluaran: JenisPengeluaranPost,
		HargaPengeluaran: HargaPengeluaranPost,
		JumlahBarang:     JumlahBarangPost,
		Satuan:           SatuanPost,
		Tipe:             TipePost,
		TotalPengeluaran: TotalPengeluaranPost,
		WaktuPengeluaran: WaktuPengeluaranPost,
	}
	// Assuming models.DB is your database connection
	if err := models.DB.Create(&pengeluaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": pengeluaran, "message": "Your pengeluaran is successfully created!"})
}

func Create_detail_bill2(c *gin.Context) {
	var detail_bills []models.Detail_bill

	if err := c.ShouldBindJSON(&detail_bills); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	// Extract unique id_bill values from the incoming data
	idBills := make(map[int64]struct{})
	for _, detail_bill := range detail_bills {
		idBills[detail_bill.IdBill] = struct{}{}
	}

	// Delete existing detail_bill records with the same id_bill values
	for idBill := range idBills {
		if err := models.DB.Where("id_bill = ?", idBill).Delete(&models.Detail_bill{}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to delete existing detail_bill records"})
			return
		}
	}

	// Save new detail_bill records to the database
	if err := models.DB.Create(&detail_bills).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": detail_bills, "message": "Your detail bills are successfully created!"})
}

func Create_klien(c *gin.Context) {

	// Parse multipart form data
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	NamaKlienPost := c.PostForm("nama_klien")
	NomorHpPost := c.PostForm("nomor_hp")
	EmailKlienPost := c.PostForm("email_klien")

	// Save file details to the database
	klien := models.Klien{
		NamaKlien:  NamaKlienPost,
		NomorHp:    NomorHpPost,
		EmailKlien: EmailKlienPost,

		// Add other fields from the form data as needed
	}
	// Assuming models.DB is your database connection
	if err := models.DB.Create(&klien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": -1, "error": "Failed to save data to database"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": klien, "message": "Your klien data is successfully created!"})
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
func Update_pengeluaran(c *gin.Context) {
	var pengeluaran models.Pengeluaran
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&pengeluaran); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	if models.DB.Model(&pengeluaran).Where("id = ?", id).Updates(&pengeluaran).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Data is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data updated"}) //mengembalikan status yg benar
}
func Update_klien(c *gin.Context) {
	var klien models.Klien
	id := c.Param("id") //mengambil params dari url yg disediakan main.go

	if err := c.ShouldBindJSON(&klien); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	if models.DB.Model(&klien).Where("id = ?", id).Updates(&klien).RowsAffected == 0 { //mengecek apakah id yg diinput ada di database atau tidak
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Data is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data updated"}) //mengembalikan status yg benar
}
func Delete_bill_old(c *gin.Context) {
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

	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	IdPostStr := c.PostForm("id")

	id, err := strconv.ParseInt(IdPostStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Invalid detail bill ID"})
		return
	}

	if models.DB.Delete(&detail_bill, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Detail bill is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}

func Delete_bill(c *gin.Context) {
	var bill models.Bill

	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	IdPostStr := c.PostForm("id")

	id, err := strconv.ParseInt(IdPostStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Invalid bill ID"})
		return
	}

	// Begin a transaction
	tx := models.DB.Begin()
	if tx.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": "Failed to start transaction"})
		return
	}

	// Find the bill
	if err := tx.Where("id = ?", id).First(&bill).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Bill not found"})
		return
	}

	// Delete associated detail bills
	if err := tx.Where("id_bill = ?", id).Delete(&models.Detail_bill{}).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": "Failed to delete detail bills"})
		return
	}

	// Delete the bill
	if err := tx.Delete(&bill).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": "Failed to delete bill"})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": -1, "data": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}

func Delete_klien(c *gin.Context) {
	var klien models.Klien

	var input struct {
		Id json.Number
	}
	// id, _ := strconv.ParseInt(input["id"], 10, 64) //melakukan perubahan string to integer, dengan ukuran integer 10 dan size integer 64

	if err := c.ShouldBindJSON(&input); err != nil { //create menggunakan input json sehinggap pengecekan juga menggunakan json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": err.Error()})
		return
	}

	id, _ := input.Id.Int64()

	if models.DB.Delete(&klien, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Klien is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}
func Delete_pengeluaran(c *gin.Context) {

	var pengeluarans models.Pengeluaran

	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	IdPostStr := c.PostForm("id")

	id, err := strconv.ParseInt(IdPostStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Invalid pengeluaran ID"})
		return
	}

	if models.DB.Delete(&pengeluarans, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "data": "Pengeluaran is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 1, "data": "Data deleted successfully"})
}

package pendapatancontroller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/andricomauludi/backend-etalase-mornin/models"

	"github.com/gin-gonic/gin"
)

// Handler function for /api/total-current-month endpoint
func TotalCurrentMonth(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total of 'Total' field for bills within the current month
	if err := models.DB.Model(&Bill).
		Select("SUM(total) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfMonth, endOfMonth, 0).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_current_month": total.Int64})
}

// Handler function for /api/total-current-month endpoint
func TotalCurrentMonthJenisPembayaran(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPembayaranPost := c.PostForm("jenis_pembayaran")
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Start building the query
	query := models.DB.Model(&Bill).Select("SUM(total) as total").Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfMonth, endOfMonth, 0)

	// Only add the jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "Semua Jenis" {
		query = query.Where("jenis_pembayaran = ?", JenisPembayaranPost)
	}

	// Execute the query
	if err := query.Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Set the response type based on JenisPembayaranPost
	responseType := JenisPembayaranPost
	if JenisPembayaranPost == "Semua Jenis" {
		responseType = "Semua Jenis Pembayaran"
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{
		"status":              1,
		"total_current_month": total.Int64,
		"type":                responseType,
	})
}

func TotalCurrentMonthCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total of 'Total' field for bills within the current month
	if err := models.DB.Model(&Bill).
		Select("SUM(total) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 1, startOfMonth, endOfMonth, 0).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_current_month": total.Int64})
}

// Handler function for /api/total-today endpoint
func TotalToday(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get today's start and end dates
	today := time.Now()
	startOfToday := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfToday := startOfToday.AddDate(0, 0, 1).Add(-time.Nanosecond)

	// Query the total of 'Total' field for bills created today
	if err := models.DB.Model(&Bill).
		Select("SUM(total) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfToday, endOfToday, 0).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for today
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_today": total.Int64})
}
func TotalTodayJenisPembayaran(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPembayaranPost := c.PostForm("jenis_pembayaran")
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get today's start and end dates
	today := time.Now()
	startOfToday := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfToday := startOfToday.AddDate(0, 0, 1).Add(-time.Nanosecond)

	// Start building the query
	query := models.DB.Model(&Bill).Select("SUM(total) as total").Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfToday, endOfToday, 0)

	// Only add the jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "Semua Jenis" {
		query = query.Where("jenis_pembayaran = ?", JenisPembayaranPost)
	}

	// Execute the query
	if err := query.Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Set the response type based on JenisPembayaranPost
	responseType := JenisPembayaranPost
	if JenisPembayaranPost == "Semua Jenis" {
		responseType = "Semua Jenis Pembayaran"
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{
		"status":      1,
		"total_today": total.Int64,
		"type":        responseType,
	})
}
func TotalTodayCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var Bill models.Bill
	var total sql.NullInt64

	// Get today's start and end dates
	today := time.Now()
	startOfToday := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfToday := startOfToday.AddDate(0, 0, 1).Add(-time.Nanosecond)

	// Query the total of 'Total' field for bills created today
	if err := models.DB.Model(&Bill).
		Select("SUM(total) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 1, startOfToday, endOfToday, 0).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for today
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_today": total.Int64})
}

// Handler function for /api/total-pengeluaran-current-month endpoint
func TotalPengeluaranCurrentMonth(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total of 'TotalPengeluaran' field for expenses within the current month
	if err := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfMonth, endOfMonth).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_pengeluaran_current_month": total.Int64})
}
func TotalPengeluaranCurrentMonthJenis(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	// Parse the multipart form to get the 'jenis_pengeluaran'
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPengeluaranPost := c.PostForm("jenis_pengeluaran")
	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Start building the query
	query := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfMonth, endOfMonth)

	// Only add the jenis_pengeluaran filter if JenisPengeluaranPost is not empty
	if JenisPengeluaranPost != "" && JenisPengeluaranPost != "Semua Jenis" {
		query = query.Where("jenis_pengeluaran = ?", JenisPengeluaranPost)
	}

	// Execute the query
	if err := query.Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Set the response type based on JenisPengeluaranPost
	responseType := JenisPengeluaranPost
	if JenisPengeluaranPost == "Semua Jenis" {
		responseType = "Semua Jenis Pengeluaran"
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{
		"status":                          1,
		"total_pengeluaran_current_month": total.Int64,
		"type":                            responseType,
	})
}

func TotalPengeluaranCurrentMonthCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total of 'TotalPengeluaran' field for expenses within the current month
	if err := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 1, startOfMonth, endOfMonth).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for the current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_pengeluaran_current_month": total.Int64})
}

// Handler function for /api/total-pengeluaran-today endpoint
func TotalPengeluaranToday(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the start and end dates for today
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Query the total of 'TotalPengeluaran' field for expenses within today
	if err := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfDay, endOfDay).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for today
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_pengeluaran_today": total.Int64})
}
func TotalPengeluaranTodayJenis(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	// Parse the multipart form to get the 'jenis_pengeluaran' (optional)
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPengeluaranPost := c.PostForm("jenis_pengeluaran")
	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the start and end dates for today
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Start building the query
	query := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfDay, endOfDay)

	// Only add the jenis_pengeluaran filter if JenisPengeluaranPost is not empty
	if JenisPengeluaranPost != "" && JenisPengeluaranPost != "Semua Jenis" {
		query = query.Where("jenis_pengeluaran = ?", JenisPengeluaranPost)
	}

	// Execute the query
	if err := query.Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Set the response type based on JenisPengeluaranPost
	responseType := JenisPengeluaranPost
	if JenisPengeluaranPost == "" || JenisPengeluaranPost == "Semua Jenis" {
		responseType = "Semua Jenis Pengeluaran"
	}

	// Return JSON response with total amount for today
	c.JSON(http.StatusOK, gin.H{
		"status":                  1,
		"total_pengeluaran_today": total.Int64,
		"type":                    responseType,
	})
}

func TotalPengeluaranTodayCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var total sql.NullInt64

	// Get the start and end dates for today
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Query the total of 'TotalPengeluaran' field for expenses within today
	if err := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 1, startOfDay, endOfDay).
		Scan(&total).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If total is NULL, set it to 0
	if !total.Valid {
		total.Int64 = 0
	}

	// Return JSON response with total amount for today
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_pengeluaran_today": total.Int64})
}

func TotalKeuntunganBersihCurrentMonth(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var bill models.Bill

	var totalCurrentMonth, totalPengeluaranCurrentMonth, totalKeuntunganBersih int64
	var err error

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total current month
	if err = models.DB.Model(&bill).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfMonth, endOfMonth, 0).
		Scan(&totalCurrentMonth).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total current month"})
		return
	}

	// Query the total pengeluaran current month
	if err = models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfMonth, endOfMonth).
		Scan(&totalPengeluaranCurrentMonth).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total pengeluaran current month"})
		return
	}

	// Calculate total keuntungan bersih current month
	totalKeuntunganBersih = totalCurrentMonth - totalPengeluaranCurrentMonth

	// Return JSON response with total keuntungan bersih current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_month": totalKeuntunganBersih})
}
func TotalKeuntunganBersihCurrentMonthJenis(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var bill models.Bill

	var totalCurrentMonth, totalPengeluaranCurrentMonth, totalKeuntunganBersih int64
	var err error

	// Parse the multipart form to get the 'jenis_pembayaran' (optional)
	err = c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPembayaranPost := c.PostForm("jenis_pembayaran")

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total current month
	queryBill := models.DB.Model(&bill).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfMonth, endOfMonth, 0)

	// Only add jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "" && JenisPembayaranPost != "Semua Jenis" {
		queryBill = queryBill.Where("jenis_pembayaran = ?", JenisPembayaranPost)
	}

	if err = queryBill.Scan(&totalCurrentMonth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total current month"})
		return
	}

	// Query the total pengeluaran current month
	queryPengeluaran := models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfMonth, endOfMonth)

	// Only add jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "" && JenisPembayaranPost != "Semua Jenis" {
		queryPengeluaran = queryPengeluaran.Where("jenis_pengeluaran = ?", JenisPembayaranPost)
	}

	if err = queryPengeluaran.Scan(&totalPengeluaranCurrentMonth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total pengeluaran current month"})
		return
	}

	// Calculate total keuntungan bersih current month
	totalKeuntunganBersih = totalCurrentMonth - totalPengeluaranCurrentMonth

	// Return JSON response with total keuntungan bersih current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_month": totalKeuntunganBersih})
}

func TotalKeuntunganBersihCurrentMonthCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var pengeluaran models.Pengeluaran
	var bill models.Bill

	var totalCurrentMonth, totalPengeluaranCurrentMonth, totalKeuntunganBersih int64
	var err error

	// Get the current month start and end dates
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Query the total current month
	if err = models.DB.Model(&bill).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 1, startOfMonth, endOfMonth, 0).
		Scan(&totalCurrentMonth).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total current month"})
		return
	}

	// Query the total pengeluaran current month
	if err = models.DB.Model(&pengeluaran).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 1, startOfMonth, endOfMonth).
		Scan(&totalPengeluaranCurrentMonth).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total pengeluaran current month"})
		return
	}

	// Calculate total keuntungan bersih current month
	totalKeuntunganBersih = totalCurrentMonth - totalPengeluaranCurrentMonth

	// Return JSON response with total keuntungan bersih current month
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_month": totalKeuntunganBersih})
}

// Handler function for /api/total-keuntungan-bersih-current-day endpoint
func TotalKeuntunganBersihCurrentDay(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var totalIncome int64
	var totalExpense int64
	var netProfit int64

	// Get the start and end of the current day
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Calculate total income for the current day
	if err := models.DB.Model(&models.Bill{}).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfDay, endOfDay, 0).
		Scan(&totalIncome).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total expense for the current day
	if err := models.DB.Model(&models.Pengeluaran{}).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfDay, endOfDay).
		Scan(&totalExpense).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate net profit for the current day
	netProfit = totalIncome - totalExpense

	// Return JSON response with net profit for the current day
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_day": netProfit})
}
func TotalKeuntunganBersihCurrentDayJenis(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var totalIncome int64
	var totalExpense int64
	var netProfit int64

	// Parse the multipart form to get the 'jenis_pembayaran' (optional)
	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "error": "Failed to parse multipart form"})
		return
	}

	JenisPembayaranPost := c.PostForm("jenis_pembayaran")

	// Get the start and end of the current day
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Calculate total income for the current day
	queryIncome := models.DB.Model(&models.Bill{}).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 0, startOfDay, endOfDay, 0)

	// Only add jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "" && JenisPembayaranPost != "Semua Jenis" {
		queryIncome = queryIncome.Where("jenis_pembayaran = ?", JenisPembayaranPost)
	}

	if err := queryIncome.Scan(&totalIncome).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total expense for the current day
	queryExpense := models.DB.Model(&models.Pengeluaran{}).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 0, startOfDay, endOfDay)

	// Only add jenis_pembayaran filter if JenisPembayaranPost is not empty
	if JenisPembayaranPost != "" && JenisPembayaranPost != "Semua Jenis" {
		queryExpense = queryExpense.Where("jenis_pengeluaran = ?", JenisPembayaranPost)
	}

	if err := queryExpense.Scan(&totalExpense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate net profit for the current day
	netProfit = totalIncome - totalExpense

	// Return JSON response with net profit for the current day
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_day": netProfit})
}

func TotalKeuntunganBersihCurrentDayCvj(c *gin.Context) {
	if models.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection is nil"})
		return
	}

	var totalIncome int64
	var totalExpense int64
	var netProfit int64

	// Get the start and end of the current day
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	// Calculate total income for the current day
	if err := models.DB.Model(&models.Bill{}).
		Select("COALESCE(SUM(total), 0) as total").
		Where("tipe = ? AND timestamp BETWEEN ? AND ? AND paid != ?", 1, startOfDay, endOfDay, 0).
		Scan(&totalIncome).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total expense for the current day
	if err := models.DB.Model(&models.Pengeluaran{}).
		Select("COALESCE(SUM(total_pengeluaran), 0) as total").
		Where("tipe = ? AND waktu_pengeluaran BETWEEN ? AND ?", 1, startOfDay, endOfDay).
		Scan(&totalExpense).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate net profit for the current day
	netProfit = totalIncome - totalExpense

	// Return JSON response with net profit for the current day
	c.JSON(http.StatusOK, gin.H{"status": 1, "total_keuntungan_bersih_current_day": netProfit})
}

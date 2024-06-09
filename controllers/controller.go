package controllers

import (
	"backend/database"
	"backend/models"
	"backend/query"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.GetMenu()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	var results []models.DataMenu

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var dataMenu models.DataMenu

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&dataMenu.Id, &dataMenu.Nama, &dataMenu.CreateDate, &dataMenu.IdKategori, &dataMenu.NamaKategori, &dataMenu.Image, &dataMenu.Status, &dataMenu.Harga)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, dataMenu)
	}

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func MenuOrder(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.GetMenuOrder()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	var results []models.DataMenu

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var dataMenu models.DataMenu

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&dataMenu.Id, &dataMenu.Nama, &dataMenu.CreateDate, &dataMenu.IdKategori, &dataMenu.NamaKategori, &dataMenu.Image, &dataMenu.Status, &dataMenu.Harga)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, dataMenu)
	}

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func Kategori(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.GetKategori()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	var results []models.DataKategori

	// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
	for rows.Next() {

		var dataKategori models.DataKategori

		// Scan data dari baris ke variabel-variabel
		err := rows.Scan(&dataKategori.Id, &dataKategori.Nama_kategori)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		results = append(results, dataKategori)
	}

	// Cek apakah ada kesalahan saat mengiterasi melalui hasil
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func AddMenu(c *gin.Context) {
	var menu models.AddMenu

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&menu); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lakukan sesuatu dengan data menu, misalnya simpan ke database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	epochTime := time.Now().Unix()

	query := query.AddMenu()
	rows, err := db.Query(query, &menu.Makanan, &epochTime, &menu.Kategori, &menu.Makanan, &menu.Harga, 1)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	// // Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Menu added successfully", "menu": menu})
}

func EditMenu(c *gin.Context) {
	var menu models.EditMenu

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&menu); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lakukan sesuatu dengan data menu, misalnya simpan ke database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.EditMenu()
	rows, err := db.Query(query, &menu.Harga, &menu.IdMenu)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	// // Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Menu edited", "menu": menu})
}

func HapusMenu(c *gin.Context) {
	var id_menu string

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&id_menu); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lakukan sesuatu dengan data menu, misalnya simpan ke database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.HapusMenu()
	rows, err := db.Query(query, &id_menu)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	// // Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Menu removed"})
}

func InStockMenu(c *gin.Context) {
	var id_menu string

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&id_menu); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lakukan sesuatu dengan data menu, misalnya simpan ke database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := query.InStockMenu()
	rows, err := db.Query(query, &id_menu)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()

	// // Kirim respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Menu in stock"})
}

func FilterMenu(c *gin.Context) {
	var id_kategori string

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&id_kategori); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// KONEK DATABASE
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var results []models.DataMenu

	if len(id_kategori) > 0 {
		query := query.FilterMenu(id_kategori)
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Err", err.Error())
		}
		defer rows.Close()

		// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
		for rows.Next() {

			var dataMenu models.DataMenu

			// Scan data dari baris ke variabel-variabel
			err := rows.Scan(&dataMenu.Id, &dataMenu.Nama, &dataMenu.CreateDate, &dataMenu.IdKategori, &dataMenu.NamaKategori, &dataMenu.Image, &dataMenu.Status, &dataMenu.Harga)
			if err != nil {
				fmt.Println("Error scanning row:", err.Error())
				return
			}

			// Tambahkan peta ke dalam slice
			results = append(results, dataMenu)
		}
	} else {
		query := query.GetMenu()
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Err", err.Error())
		}
		defer rows.Close()

		// Loop melalui setiap baris hasil dan tambahkan data ke dalam slice
		for rows.Next() {

			var dataMenu models.DataMenu

			// Scan data dari baris ke variabel-variabel
			err := rows.Scan(&dataMenu.Id, &dataMenu.Nama, &dataMenu.CreateDate, &dataMenu.IdKategori, &dataMenu.NamaKategori, &dataMenu.Image, &dataMenu.Status, &dataMenu.Harga)
			if err != nil {
				fmt.Println("Error scanning row:", err.Error())
				return
			}

			// Tambahkan peta ke dalam slice
			results = append(results, dataMenu)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filter Kategori Succesfully", "menu": results})
}

func AddTransaksi(c *gin.Context) {
	// code
	var dataTransaksi []models.DataTransaksi

	// BindJSON akan mencoba untuk mengurai JSON yang dikirim dalam body request
	if err := c.BindJSON(&dataTransaksi); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var atasNama string
	var totalHarga int
	var jumlah int

	for _, transaksi := range dataTransaksi {
		// Ambil field yang diinginkan
		atasNama = transaksi.AtasNama
		jumlah += transaksi.Jumlah
		hargaSatuan, err := strconv.Atoi(transaksi.HargaSatuan)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jumlah format"})
			return
		}

		// Lakukan operasi yang diperlukan, misalnya menghitung total harga
		totalHarga = jumlah * hargaSatuan
	}

	// Lakukan sesuatu dengan data menu, misalnya simpan ke database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	epochTime := time.Now().Unix()

	// Query INSERT
	result, err := db.Exec("INSERT INTO transaksi (atas_nama, sub_total, create_date, is_cancel) VALUES (?, ?, ?, ?)", atasNama, totalHarga, epochTime, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Dapatkan ID yang baru dimasukkan
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
		} else {

			// loop data
			for _, detailTransaksi := range dataTransaksi {
				// insert detail
				_, err := db.Exec("INSERT INTO detail_transaksi (id_transaksi, id_menu, id_kategori, qty, create_date, remark, keadaan, harga_satuan, is_cancel) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", id, detailTransaksi.IdMenu, detailTransaksi.IdKategori, detailTransaksi.Jumlah, epochTime, detailTransaksi.Remark, detailTransaksi.Keadaan, detailTransaksi.HargaSatuan, 0)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			// Kirim respons berhasil
			c.JSON(http.StatusOK, gin.H{"message": "Menu added successfully", "transaksi": dataTransaksi})
		}
	}
}

func GetDataTransaksi(c *gin.Context) {
	var date string
	if err := c.BindJSON(&date); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// define struct
	var dataTable []models.DataTableTransaksi
	// ngequery
	query := query.GetDataTransaksi(date)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer rows.Close()
	// scaning for table
	for rows.Next() {
		var dataRow models.DataTableTransaksi
		err := rows.Scan(&dataRow.IdTransaksi, &dataRow.IdDetailTransaksi, &dataRow.AtasNama, &dataRow.NamaKategori, &dataRow.Makanan, &dataRow.IsCancel, &dataRow.SubTotal, &dataRow.CreateDate, &dataRow.Qty, &dataRow.Remark, &dataRow.HargaSatuan)
		if err != nil {
			fmt.Println("Error scanning row:", err.Error())
			return
		}

		// Tambahkan peta ke dalam slice
		dataTable = append(dataTable, dataRow)
	}

	// define for chart
	var dataChart []models.DataChartTransaksi
	// ngequery
	queryChart := `SELECT D.nama_kategori as 'name', count(B.qty) as 'y'
	FROM transaksi A
	LEFT JOIN detail_transaksi B ON A.id = B.id_transaksi
	LEFT JOIN kategori_menu D ON D.id = B.id_kategori
	WHERE FROM_UNIXTIME(A.create_date, '%Y/%m/%d') = '` + date + `'
	GROUP BY D.nama_kategori`
	chart, err := db.Query(queryChart)
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer chart.Close()
	// scanning for chart
	for chart.Next() {
		var dataDiagram models.DataChartTransaksi

		// for chart
		errScanChat := chart.Scan(&dataDiagram.Name, &dataDiagram.Y)
		if errScanChat != nil {
			fmt.Println("Error scanning row:", errScanChat.Error())
			return
		}

		dataChart = append(dataChart, dataDiagram)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get Data Transaksi Succesfully", "data_table": dataTable, "data_chart": dataChart})
}

func GetDashboard(c *gin.Context) {
	var date string
	if err := c.BindJSON(&date); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var results []models.DataCountDashboard
	// ngequery
	query := query.GetDataDashboard(date)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer rows.Close()
	// Memindai hasil query
	for rows.Next() {
		var dataCount models.DataCountDashboard
		err := rows.Scan(&dataCount.TotalSales, &dataCount.NumOfOrders, &dataCount.TotalMenu, &dataCount.TotalKategori)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memindai baris"})
			return
		}
		results = append(results, dataCount)
	}
	rows.Close()

	// Mengembalikan hasil sebagai JSON
	c.JSON(http.StatusOK, results)
}

func GetChart(c *gin.Context) {
	var date string
	if err := c.BindJSON(&date); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Parsing string ke dalam format time.Time
	layout := "2006/01/02"
	dateStr, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	// Mengambil tahunnya saja
	year := dateStr.Year()
	// konversi string
	yearStr := strconv.Itoa(year)
	// define struct
	var results models.DataChart
	// konek database
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// GET DATA DAILY CHART
	queryDaily := query.GetDailyChart(yearStr)
	rows, err := db.Query(queryDaily)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer rows.Close()
	// Memindai hasil query
	for rows.Next() {
		var daily models.Daily
		err := rows.Scan(&daily.Date, &daily.Rupiah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memindai baris"})
			return
		}
		results.Daily = append(results.Daily, daily)
	}

	// GET DATA MONTHLY CHART
	querymonth := query.GetMonthlyChart(yearStr)
	rowsMonth, err := db.Query(querymonth)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer rowsMonth.Close()

	// Memindai hasil query bulanan
	for rowsMonth.Next() {
		var monthly models.Monthly // Ganti dengan struktur yang sesuai jika ada
		err := rowsMonth.Scan(&monthly.Date, &monthly.Rupiah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memindai baris bulanan"})
			return
		}
		results.Monthly = append(results.Monthly, monthly) // Pastikan `Monthly` adalah slice yang sesuai
	}

	// GET DATA YEARLY CHART
	querYear := query.GetYearlyChart(yearStr)
	rowsYear, err := db.Query(querYear)
	if err != nil {
		fmt.Println("Err", err.Error())
		return
	}
	defer rowsYear.Close()

	// Memindai hasil query bulanan
	for rowsYear.Next() {
		var yearly models.Yearly
		err := rowsYear.Scan(&yearly.Date, &yearly.Rupiah)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memindai baris bulanan"})
			return
		}
		results.Yearly = append(results.Yearly, yearly)
	}

	// Kembali dengan hasil
	c.JSON(http.StatusOK, results)
}

func CancelTransaksi(c *gin.Context) {
	var id_detail_transaksi string
	if err := c.BindJSON(&id_detail_transaksi); err != nil {
		// Jika terjadi kesalahan dalam menguraikan JSON, kirim respons error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	queryCancel := query.CancelTransaksi(id_detail_transaksi)
	rows, err := db.Query(queryCancel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Menu added successfully"})
	}
	defer rows.Close()

}

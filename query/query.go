package query

func GetMenu() string {
	return `SELECT 
	A.id, A.nama, A.create_date, B.id as 'id_kategori', B.nama_kategori,  A.image, A.status, A.harga
	FROM menu A
	LEFT JOIN kategori_menu B ON B.id = A.id_kategori
	ORDER BY A.status DESC, A.create_date DESC`
}

func GetMenuOrder() string {
	return `SELECT 
	A.id, A.nama, A.create_date, B.id as 'id_kategori', B.nama_kategori,  A.image, A.status, A.harga
	FROM menu A
	LEFT JOIN kategori_menu B ON B.id = A.id_kategori
	WHERE A.status = 1
	ORDER BY A.status DESC, A.create_date DESC`
}

func GetKategori() string {
	return `SELECT id, nama_kategori FROM kategori_menu where is_active = 1`
}

func AddMenu() string {
	return `INSERT INTO menu(nama, create_date, id_kategori, image, harga, status) VALUES (?, ?, ?, ?, ?, ?)`
}

func EditMenu() string {
	return `UPDATE menu SET harga = ? WHERE id = ?`
}

func HapusMenu() string {
	return `UPDATE menu SET status=0 WHERE id = ?`
}
func InStockMenu() string {
	return `UPDATE menu SET status = 1 WHERE id = ?`
}

func FilterMenu(id_kategori string) string {
	return `SELECT 
	A.id, A.nama, A.create_date, B.id as 'id_kategori', B.nama_kategori,  A.image, A.status, A.harga
	FROM menu A
	LEFT JOIN kategori_menu B ON B.id = A.id_kategori
	WHERE A.id_kategori = ` + id_kategori + ` AND A.status = 1
	ORDER BY A.status DESC, A.create_date DESC`
}

func GetDataTransaksi(date string) string {
	return `SELECT A.id, B.id as 'id_detail', A.atas_nama, D.nama_kategori, C.nama, B.is_cancel, (B.qty * B.harga_satuan) as 'sub_total', FROM_UNIXTIME(A.create_date, '%Y/%m/%d') as 'create_date',
	B.qty, B.remark, B.harga_satuan
	FROM transaksi A
	LEFT JOIN detail_transaksi B ON A.id = B.id_transaksi
	LEFT JOIN menu C ON C.id = B.id_menu
	LEFT JOIN kategori_menu D ON D.id = B.id_kategori
	WHERE FROM_UNIXTIME(A.create_date, '%Y/%m/%d') = '` + date + `'`
}

func GetDataChart(date string) string {
	return `SELECT D.nama_kategori as 'name', count(B.qty) as 'y'
	FROM transaksi A
	LEFT JOIN detail_transaksi B ON A.id = B.id_transaksi
	LEFT JOIN kategori_menu D ON D.id = B.id_kategori
    WHERE FROM_UNIXTIME(A.create_date, '%Y/%m/%d') = ` + date + `
    GROUP BY D.nama_kategori`
}

func GetDataDashboard(date string) string {
	return `SELECT 
    A.total_sales,
    A.number_of_orders,
    B.total_menu,
    C.total_kategori
FROM (
    SELECT 
        0 as id_menu,
        1 as id_kategori,
        SUM(harga_satuan * qty) as total_sales,
        COUNT(id) as number_of_orders
    FROM detail_transaksi 
    WHERE DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y/%m/%d') = '` + date + `'
) A
LEFT JOIN (
    SELECT 
        0 as id_menu, 
        COUNT(id) as total_menu 
    FROM menu 
    WHERE status = 1
) B ON B.id_menu = A.id_menu
LEFT JOIN (
    SELECT 
        1 as id_kategori, 
        COUNT(id) as total_kategori 
    FROM kategori_menu 
    WHERE is_active = 1
) C ON C.id_kategori = A.id_kategori`
}

func GetDailyChart(date string) string {
	return `SELECT 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y/%m/%d') as 'date', 
	SUM(qty * harga_satuan)  as 'rupiah' 
	FROM detail_transaksi 
	WHERE DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y') = '` + date + `' 
	GROUP BY 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y/%m/%d')`
}

func GetMonthlyChart(date string) string {
	return `SELECT 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y/%m') as 'date', 
	SUM(qty * harga_satuan)  as 'rupiah' 
	FROM detail_transaksi 
	WHERE DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y') = '` + date + `' 
	GROUP BY 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y/%m')`
}

func GetYearlyChart(date string) string {
	return `SELECT 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y') as 'date', 
	SUM(qty * harga_satuan)  as 'rupiah' 
	FROM detail_transaksi 
	WHERE DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y') = '` + date + `' 
	GROUP BY 
	DATE_FORMAT(FROM_UNIXTIME(create_date), '%Y')`
}

func CancelTransaksi(id_detail_transaksi string) string {
	return `UPDATE detail_transaksi SET is_cancel=1 WHERE id = ` + id_detail_transaksi + ``
}

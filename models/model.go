package models

type DataMenu struct {
	Id           string `json:"id"`
	Nama         string `json:"nama"`
	CreateDate   string `json:"create_date"`
	IdKategori   string `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
	Image        string `json:"image"`
	Status       string `json:"status"`
	Harga        string `json:"harga"`
}

type DataKategori struct {
	Id            string `json:"id"`
	Nama_kategori string `json:"nama_kategori"`
}

type AddMenu struct {
	Makanan  string `json:"makanan"`
	Kategori string `json:"kategori"`
	Harga    string `json:"harga"`
	// Image    string `json:"image"`
}

type EditMenu struct {
	IdMenu string `json:"id_menu"`
	Harga  string `json:"harga"`
	// Image    string `json:"image"`
}

type DataTransaksi struct {
	AtasNama    string `json:"atas_nama"`
	Jumlah      int    `json:"count"`
	HargaSatuan string `json:"harga"`
	IdMenu      string `json:"id"`
	IdKategori  string `json:"id_kategori"`
	NamaMakanan string `json:"nama"`
	Remark      string `json:"remark"`
	Keadaan     string `json:"keadaan"` // hot or ice
}

type DataTableTransaksi struct {
	IdTransaksi       string `json:"id"`
	IdDetailTransaksi string `json:"id_detail"`
	AtasNama          string `json:"atas_nama"`
	NamaKategori      string `json:"nama_kategori"`
	Makanan           string `json:"nama"`
	IsCancel          string `json:"is_cancel"`
	SubTotal          string `json:"sub_total"`
	CreateDate        string `json:"create_date"`
	Qty               int    `json:"qty"`
	Remark            string `json:"remark"`
	HargaSatuan       string `json:"harga_satuan"`
}

type DataChartTransaksi struct {
	Name string `json:"name"`
	Y    int    `json:"y"`
}

type DataCountDashboard struct {
	TotalSales    int `json:"total_sales"`
	NumOfOrders   int `json:"number_of_orders"`
	TotalMenu     int `json:"total_menu"`
	TotalKategori int `json:"total_kategori"`
}

type DataChart struct {
	Daily   []Daily   `json:"daily"`
	Monthly []Monthly `json:"monthly"`
	Yearly  []Yearly  `json:"yearly"`
}

type Daily struct {
	Date   string `json:"date"`
	Rupiah int    `json:"rupiah"`
}
type Monthly struct {
	Date   string `json:"date"`
	Rupiah int    `json:"rupiah"`
}
type Yearly struct {
	Date   string `json:"date"`
	Rupiah int    `json:"rupiah"`
}

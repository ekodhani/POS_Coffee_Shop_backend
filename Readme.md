# Simple POS Coffee Shop Backend

## Overview
The Coffee Shop Backend API allows clients to interact with the coffee shop's database to manage products and orders. This API responds with JSON data.

## Base URL
```url
http://localhost:8081
```
## Endpoint
#### 1. Menu
- **URL**: `/menu`
- **Method**: `POST`
- **Description**: Get all data menu.

#### 2.getMenuOrder
- **URL**: `/getMenuOrder`
- **Method**: `POST`
- **Description**: Get all data for order page.

#### 3.kategori
- **URL**: `/kategori`
- **Method**: `POST`
- **Description**: Get all data kategori menu.

#### 4.addMenu
- **URL**: `/addMenu`
- **Method**: `POST`
- **Description**: Insert data menu.
- **Request Body**:
  ```json
  {
    "id" : "string"
	"nama" : "string"
	"create_date" : "string"
	"id_kategori" : "string"
	"nama_kategori" : "string"
	"image" : "string"
	"status" : "string"
	"harga" : "string"
  }

#### 5.editMenu
- **URL**: `/editMenu`
- **Method**: `POST`
- **Description**: Update data menu.
- **Request Body**:
  ```json
  {
    "id_menu" : "string"
	"harga" : "string"
  }

#### 6.filterMenu
- **URL**: `/filterMenu`
- **Method**: `POST`
- **Description**: Filter data menu by kategori.
- **Request Body**:
  ```json
  {
    "id_kategori"
  }

#### 7.deleteMenu
- **URL**: `/deleteMenu`
- **Method**: `POST`
- **Description**: Out of stock menu.
- **Request Body**:
  ```json
  {
    "id_menu"
  }

#### 8.inStockMenu
- **URL**: `/inStockMenu`
- **Method**: `POST`
- **Description**: Set in stock menu.
- **Request Body**:
  ```json
  {
    "id_menu"
  }

#### 9.addTransaksi
- **URL**: `/addTransaksi`
- **Method**: `POST`
- **Description**: Add order.
- **Request Body**:
  ```json
  [{
    "atas_nama" : "string"
	"count" : "int"
	"harga" : "string"
	"id" : "string"
	"id_kategori" : "string"
	"nama" : "string"
	"remark" : "string"
	"keadaan"` : "string"
  },...]

#### 10.getHistoryTransaksi
- **URL**: `/getHistoryTransaksi`
- **Method**: `POST`
- **Description**: Get all data transaksi by Date (YYYY/MM/DD).
- **Request Body**:
  ```json
  {
    "date"
  }

#### 11.getDashboard
- **URL**: `/getDashboard`
- **Method**: `POST`
- **Description**: Get count summary dashboard by date (YYYY/MM/DD).
- **Request Body**:
  ```json
  {
    "date"
  }

#### 12.getDataChart
- **URL**: `/getDataChart`
- **Method**: `POST`
- **Description**: Get data chart for dashboard page by date (YYYY/MM/DD).
- **Request Body**:
  ```json
  {
    "date"
  }

#### 13.cancelTransaksi
- **URL**: `/cancelTransaksi`
- **Method**: `POST`
- **Description**: Cancel transaksi by Id detail transaksi.
- **Request Body**:
  ```json
  {
    "id_detail_transaksi"
  }
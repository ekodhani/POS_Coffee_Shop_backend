-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 09 Jun 2024 pada 05.04
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pos_flow_coffee`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_transaksi`
--

CREATE TABLE `detail_transaksi` (
  `id` int(11) NOT NULL,
  `id_transaksi` int(11) NOT NULL,
  `id_menu` int(11) NOT NULL,
  `id_kategori` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `create_date` int(11) NOT NULL,
  `remark` varchar(255) NOT NULL,
  `keadaan` varchar(255) NOT NULL COMMENT 'ice or hot',
  `harga_satuan` int(100) NOT NULL,
  `is_cancel` int(2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `detail_transaksi`
--

INSERT INTO `detail_transaksi` (`id`, `id_transaksi`, `id_menu`, `id_kategori`, `qty`, `create_date`, `remark`, `keadaan`, `harga_satuan`, `is_cancel`) VALUES
(1, 1, 17, 2, 2, 1716641477, '', '', 15000, 0),
(2, 1, 8, 3, 2, 1716641477, '', 'panas', 3000, 0),
(3, 2, 14, 1, 1, 1716641537, '', 'panas', 25000, 0),
(4, 2, 6, 5, 1, 1716641537, '', '', 15000, 0),
(5, 2, 14, 1, 1, 1716641537, '', 'dingin', 25000, 0),
(6, 3, 2, 1, 1, 1716641568, '', 'dingin', 25000, 0),
(7, 3, 15, 4, 1, 1716641568, '', 'dingin', 20000, 0),
(8, 4, 1, 1, 1, 1716641594, '', 'panas', 30000, 0),
(9, 4, 8, 3, 1, 1716641594, '', 'panas', 3000, 0),
(10, 5, 17, 2, 2, 1716730488, '', '', 15000, 0),
(11, 5, 15, 4, 2, 1716730488, '', 'dingin', 20000, 0),
(12, 6, 16, 2, 1, 1716730561, '', '', 25000, 0),
(13, 6, 5, 2, 1, 1716730561, '', '', 30000, 0),
(14, 6, 8, 3, 2, 1716730561, '', 'dingin', 3000, 0),
(15, 7, 17, 2, 1, 1717314235, '', '', 15000, 0),
(16, 7, 16, 2, 1, 1717314235, '', '', 25000, 0),
(17, 8, 17, 2, 1, 1717834835, '', '', 15000, 0),
(18, 8, 15, 4, 1, 1717834835, '', 'dingin', 20000, 0),
(19, 15, 17, 2, 1, 1717843372, '', '', 15000, 0),
(20, 15, 16, 2, 1, 1717843372, '', '', 25000, 0),
(21, 16, 16, 2, 1, 1717843405, '', '', 25000, 0),
(22, 16, 15, 4, 1, 1717843405, '', '', 20000, 0),
(23, 17, 15, 4, 1, 1717843453, '', '', 20000, 0),
(24, 17, 8, 3, 1, 1717843453, '', '', 3000, 0),
(25, 18, 1, 1, 1, 1717843479, '', 'panas', 30000, 0),
(26, 18, 16, 2, 1, 1717843479, '', '', 25000, 0),
(29, 20, 14, 1, 1, 1717895419, '', 'panas', 25000, 1),
(30, 21, 16, 2, 1, 1717895610, '', '', 25000, 0),
(31, 22, 17, 2, 1, 1717897958, '', '', 15000, 0),
(32, 22, 16, 2, 1, 1717897958, '', '', 25000, 0),
(33, 22, 15, 4, 2, 1717897958, '', 'dingin', 20000, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `kategori_menu`
--

CREATE TABLE `kategori_menu` (
  `id` int(11) NOT NULL,
  `nama_kategori` varchar(100) NOT NULL,
  `create_date` int(11) NOT NULL,
  `is_active` int(11) NOT NULL,
  `update_date` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `kategori_menu`
--

INSERT INTO `kategori_menu` (`id`, `nama_kategori`, `create_date`, `is_active`, `update_date`) VALUES
(1, 'Coffee', 1696081171, 1, 1696081171),
(2, 'Food', 1696081171, 1, 1696081171),
(3, 'Tea', 1696081171, 1, 1696081171),
(4, 'Drinks', 1696081171, 1, 1696081171),
(5, 'Snack', 1696081171, 1, 1696081171),
(6, 'Cake', 1696081171, 1, 1696081171);

-- --------------------------------------------------------

--
-- Struktur dari tabel `menu`
--

CREATE TABLE `menu` (
  `id` int(11) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `create_date` int(11) NOT NULL,
  `id_kategori` int(11) NOT NULL,
  `image` varchar(255) NOT NULL,
  `status` int(3) NOT NULL COMMENT '1 = available, 2 = out of stock',
  `harga` int(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `menu`
--

INSERT INTO `menu` (`id`, `nama`, `create_date`, `id_kategori`, `image`, `status`, `harga`) VALUES
(1, 'Latte', 1714284066, 1, '', 1, 30000),
(2, 'Americano', 1714284066, 1, '', 1, 25000),
(3, 'Espresso', 1714284066, 1, '', 1, 25000),
(4, 'Nasi kebuli', 1714284066, 2, '', 1, 30000),
(5, 'Kwetiau', 1714284066, 2, '', 1, 30000),
(6, 'French fries', 1714284066, 5, '', 1, 15000),
(7, 'Mix Platter', 1714284066, 5, '', 1, 25000),
(8, 'Teh solo', 1714284066, 3, '', 1, 3000),
(9, 'Es timun', 1714284066, 4, '', 1, 10000),
(10, 'Air Mineral', 1714284066, 4, '', 1, 4000),
(11, 'Caramel Latte', 1715438832, 1, 'Caramel Latte', 0, 1000),
(12, 'Machiatto', 1715439143, 1, 'Machiatto', 0, 1000),
(13, 'Capucinno', 1715439315, 1, 'Capucinno', 0, 1000),
(14, 'Salted Caramel Latte', 1715440110, 1, 'Salted Caramel Latte', 1, 25000),
(15, 'Bir Pletok', 1716480682, 4, 'Bir Pletok', 1, 20000),
(16, 'Nasi goreng', 1716480779, 2, 'Nasi goreng', 1, 25000),
(17, 'Nasi Uduk', 1716480811, 2, 'Nasi Uduk', 1, 15000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `id` int(11) NOT NULL,
  `atas_nama` varchar(100) NOT NULL,
  `is_cancel` int(2) NOT NULL,
  `sub_total` int(100) NOT NULL,
  `create_date` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transaksi`
--

INSERT INTO `transaksi` (`id`, `atas_nama`, `is_cancel`, `sub_total`, `create_date`) VALUES
(1, 'Siapa', 0, 12000, 1716641477),
(2, 'Itu dulu', 0, 75000, 1716641537),
(3, 'bebas', 0, 40000, 1716641568),
(4, 'jogja', 0, 6000, 1716641594),
(5, 'Rama', 0, 80000, 1716730488),
(6, 'Joko', 0, 12000, 1716730561),
(7, 'budi', 0, 50000, 1717314235),
(8, 'budi', 0, 40000, 1717834835),
(15, 'iyya', 0, 50000, 1717843372),
(16, 'oke', 0, 40000, 1717843405),
(17, 'oke', 0, 6000, 1717843453),
(18, 'ko', 0, 50000, 1717843479),
(20, 'tes', 0, 25000, 1717895419),
(21, 'dsadd', 0, 25000, 1717895610),
(22, 'hayo', 0, 80000, 1717897958);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `detail_transaksi`
--
ALTER TABLE `detail_transaksi`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `kategori_menu`
--
ALTER TABLE `kategori_menu`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `menu`
--
ALTER TABLE `menu`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `detail_transaksi`
--
ALTER TABLE `detail_transaksi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT untuk tabel `kategori_menu`
--
ALTER TABLE `kategori_menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `menu`
--
ALTER TABLE `menu`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

package seeder

// SESUAI DENGAN DATABASE POGRESSQL

import (
	"log"
	"web-investasi/config"
)

// / function untuk membuat semua daftar table
func StartCreateTable() {
	// daftar semu table  || Sesuaikan dengan yang di model
	var tableQuery = map[string]string{
		"users": `
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,               -- ID auto increment (tipe integer)
		name VARCHAR(100) NOT NULL,          -- Nama user (maks 100 karakter)
		password VARCHAR(255) NOT NULL,      -- Password terenkripsi (hash)
		email VARCHAR(150) UNIQUE NOT NULL,  -- Email harus unik
		saldo DECIMAL(15, 2)                 -- Jumlah saldo ,  di mana 15 adalah jumlah total digit dan 2 adalah jumlah digit di belakang koma
		created_at TIMESTAMP DEFAULT NOW(),  -- Waktu pembuatan data
		updated_at TIMESTAMP DEFAULT NOW()   -- Waktu update data terakhir
		);
		`,
		"admin": `CREATE TABLE IF NOT EXISTS admin (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		role VARCHAR(50) DEFAULT 'superadmin'
		);`,
		"products": `CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,              -- ID unik auto-increment
		name VARCHAR(100) NOT NULL,         -- Nama produk
		idr NUMERIC(12,2) NOT NULL,         -- Harga dalam Rupiah (float)
		source TEXT,                        -- path dari gambar nya , boleh kosong
		description TEXT,                   -- Deskripsi produk
		status VARCHAR(50) DEFAULT 'active',-- Status produk (data: active, inactive)
		created_at TIMESTAMP DEFAULT NOW(), -- Tanggal dibuat
		updated_at TIMESTAMP DEFAULT NOW()  -- Tanggal terakhir diubah
		);
		`,
	}

	for name, query := range tableQuery {
		_, err := config.DB.Exec(query)
		if err != nil {
			log.Printf("[x] gagal membuat table %s  , error : %v ", name, err)
		}

	}

}

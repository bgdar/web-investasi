package seeder

// SESUAI DENGAN DATABASE POGRESSQL

import (
	"log"
	"web-investasi/config"
)

/// daftar semu table  || Sesuaikan dengan yang di model 
var TableQuery = map[string]string{
	"users":`
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,               -- ID auto increment (tipe integer)
	name VARCHAR(100) NOT NULL,          -- Nama user (maks 100 karakter)
	password VARCHAR(255) NOT NULL,      -- Password terenkripsi (hash)
	email VARCHAR(150) UNIQUE NOT NULL,  -- Email harus unik
	created_at TIMESTAMP DEFAULT NOW(),  -- Waktu pembuatan data
	updated_at TIMESTAMP DEFAULT NOW()   -- Waktu update data terakhir
);
`,
	"admin": 
	`CREATE TABLE IF NOT EXISTS admins (
        id SERIAL PRIMARY KEY,
        username VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role VARCHAR(50) DEFAULT 'superadmin'
    );`,
	"products":
	`CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,              -- ID unik auto-increment
    name VARCHAR(100) NOT NULL,         -- Nama produk
    idr NUMERIC(12,2) NOT NULL,         -- Harga dalam Rupiah (float)
    description TEXT,                   -- Deskripsi produk
    status VARCHAR(50) DEFAULT 'active',-- Status produk (misal: active, inactive)
    created_at TIMESTAMP DEFAULT NOW(), -- Tanggal dibuat
    updated_at TIMESTAMP DEFAULT NOW()  -- Tanggal terakhir diubah
);
`,

}

/// function untuk membuat semua daftar table 
func StartCreateTable()  {

	for name , query := range TableQuery {
		_ , err := config.DB.Exec(query)
		if err != nil {
			log.Printf("[x] gagal membuat table %s  , error : %v ",name,err)
		}
		
	}
		
}

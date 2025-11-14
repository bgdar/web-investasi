package seeder

import (
	"log"
	"web-investasi/config"
)

/// semua example data yang akan di include ke table table yang ada

// / isi data table dengan data dymmy
// / ! panggil sekali aja , biar gak  double datanya
func FillDataTable() {

	dataDummy := map[string]string{
		"users": `INSERT INTO users (name, password, email)
VALUES
('dar',    'dar', 'dar@ed.com'),
('Hannah Lee',       'example_hash', 'hannah@example.com');`,
		"products": `
		INSERT INTO products (name, idr, source, description, status)
VALUES
('AlphaBot X1', 2499000.00, 'https://images.unsplash.com/photo-1603791440384-56cd371ee9a7', 
 'Robot edukasi dasar dengan sensor inframerah dan kemampuan mengikuti garis.', 'active'),

('MechaArm 2000', 8799000.00, 'https://images.unsplash.com/photo-1581093588401-22a5b36c56b0', 
 'Lengan robotik presisi tinggi untuk proyek industri ringan dan riset.', 'active'),

('RoboDog Mini', 4599000.00, 'https://images.unsplash.com/photo-1579586337278-3befd40fd17a', 
 'Robot anjing kecil dengan gerakan realistis dan kontrol melalui aplikasi.', 'active'),

('ServiBot Home', 6999000.00, 'https://images.unsplash.com/photo-1593642533144-3d62bd1f46b9', 
 'Robot asisten rumah tangga dengan pengenalan suara dan navigasi otomatis.', 'active'),

('CleanBot Pro', 3599000.00, 'https://images.unsplash.com/photo-1626785774573-4b799315fca1', 
 'Robot pembersih lantai dengan fitur deteksi kotoran dan pemetaan ruangan.', 'active'),

('GuardX Sentinel', 9999000.00, 'https://images.unsplash.com/photo-1606813908859-3d25e4826a9a', 
 'Robot keamanan otonom dengan kamera malam dan sistem alarm otomatis.', 'active'),

('EduBot Junior', 1899000.00, 'https://images.unsplash.com/photo-1581090465774-5c1b1e33de69', 
 'Robot pembelajaran pemrograman untuk anak-anak usia 10 tahun ke atas.', 'active'),

('AeroBot Explorer', 12599000.00, 'https://images.unsplash.com/photo-1504384308090-c894fdcc538d', 
 'Robot terbang canggih dengan sensor jarak dan mode penerbangan otomatis.', 'active'),

('AquaDroid 7', 10799000.00, 'https://images.unsplash.com/photo-1613985543328-c6f3c8b8d4c6', 
 'Robot bawah air untuk eksplorasi dan riset lingkungan laut.', 'active');

		`,
		"admin": `INSERT INTO  admin (name, password, email)
VALUES
('bgdar',     'example_hash', 'akunzero975@gmail.com'),
('dar',       'example_hash', 'dar@daraja.com');`,
	}

	for key, value := range dataDummy {
		_, err := config.DB.Exec(value)
		if err != nil {
			log.Println("[x] gagal insert data dummy untuk table :", key)
		} else {
			log.Println("✓ data dummy  berhasil di tambah")
		}
	}

}

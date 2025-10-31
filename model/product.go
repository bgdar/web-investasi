package model

import (
	"database/sql"
	"log"
	"web-investasi/config"
)

// format table Product
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	IDR         float32 `json:"idr"`
	Description string  `json:"description"`
	Status string `json:"status"`
}

// / tambah product
// / tapi ini gak saya gunakan sekarang mungkin karena sudah saya siapainn data di DB pogress nya
func GetProductId(id int64) (*Product, error) {
	query := `SELECT name, email ,description FROM users WHERE id = $1;`

	// Jalankan query
	row := config.DB.QueryRow(query, id)

	// Variabel untuk menampung hasil
	var product Product
	// Scan hasil query ke variabel
	err := row.Scan(&product.Name, &product.IDR, product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[!] Data dengan ID", id, "tidak ditemukan")
			return nil, nil
		}
		log.Println("[x] Error saat mengambil data dengan ID", id, ":", err)
		return nil, nil
	}
	return &product, nil

}

// / kembalikan 5 product
func GetProductTop5() {

}

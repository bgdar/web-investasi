package model

import (
	"database/sql"
	"log"
	"time"
	"web-investasi/config"
)

// format table Product
type Product struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	IDR         float32        `json:"idr"`
	Source      sql.NullString `json:"source"` // tipe yang memperbolehkan null
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"Created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// / tambah product
func AddProduct(name string, idr float32, source string, description string) {
	input := Product{
		IDR:  idr,
		Name: name,
		Source: sql.NullString{ // handle untuk string sqlnull
			String: name,
			Valid:  name != "",
		},
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
	}
	query := `INSERT INTO  products (name, idr,source,description )
VALUES ( $1 , $2 , $3 , $4);`
	err := config.DB.QueryRow(query, input.Name, input.IDR, input.Source, input.Description)
	if err != nil {
		log.Println("[x] gagal menambah data product :", err)
		return
	}
}

// / dapatkan semua data product
// / return array of Product
func GetAllProdct() ([]Product, error) {
	var products []Product

	query := "SELECT * from products"
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Panicln("[x] gagal membaca semua product :", err)
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.IDR, &product.Source, &product.Description, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("[x] gagal mendapatkan data product tertentu :", err)
			return nil, nil
		}
		products = append(products, product)
	}
	return products, nil
}

// / tapi ini gak saya gunakan sekarang mungkin karena sudah saya siapainn data di DB pogress nya
func GetProductByID(id int64) (*Product, error) {
	query := `SELECT name, idr, description FROM products WHERE id = $1;`

	// Jalankan query
	row := config.DB.QueryRow(query, id)

	// Variabel untuk menampung hasil
	var product Product

	// Scan hasil query ke variabel
	err := row.Scan(&product.Name, &product.IDR, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[!] Data dengan ID", id, "tidak ditemukan")
			return nil, err // return error agar bisa dideteksi caller
		}
		log.Println("[x] Error saat mengambil data dengan ID", id, ":", err)
		return nil, err
	}

	return &product, nil
}

// / kembalikan 5 product
func GetProductTop5() {

}

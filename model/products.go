package model

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"time"
	"web-investasi/config"
)

// format table Product
type Products struct {
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
func AddProducts(name string, idr float32, source string, description string) {
	input := Products{
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
func GetAllProdcts() ([]Products, error) {
	var products []Products

	query := "SELECT * from products"
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Panicln("[x] gagal membaca semua product :", err)
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ID, &product.Name, &product.IDR, &product.Source, &product.Description, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println("[x](GetAllProdct) gagal mendapatkan data product tertentu :", err)
			return nil, nil
		}
		products = append(products, product)
	}
	return products, nil
}

// / tapi ini gak saya gunakan sekarang mungkin karena sudah saya siapainn data di DB pogress nya
func GetProductByIDs(id int64) (*Products, error) {
	query := `SELECT name, idr, description FROM products WHERE id = $1;`

	// Jalankan query
	row := config.DB.QueryRow(query, id)

	// Variabel untuk menampung hasil
	var product Products

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
func GetProductTop5s() {

}

// / ambil 1 data denagn katagory data apapun di table
// / T = type  yang akan di cari
func GetOneDataProducts[T any](field string, value any) (T, error) {
	var result T

	// sanitize: hanya boleh huruf/underscore
	valid := regexp.MustCompile(`^[a-zA-Z_]+$`)
	if !valid.MatchString(field) {
		return result, fmt.Errorf("invalid column name")
	}

	query := fmt.Sprintf("SELECT %s FROM products WHERE %s = $1 LIMIT 1", field, field)

	err := config.DB.QueryRow(query, value).Scan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

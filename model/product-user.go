package model

// MODEL UNTUK handle product product berdsarkan nam user
import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"time"
	"web-investasi/config"
)

type ProductUser struct {
	ID           int64        `json:"id"` // type lain untuk db : `db:"id" json:"id"`
	Name         string       `json:"name"`
	TotalProduct int16        `json:"total_product"`
	Expired      time.Time    `json:"expired"`
	DailyReward  sql.NullTime `json:"daily_reward"`
	CreatedAt    time.Time    `json:"Created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// / buat table product denagn katagory user tertentu , by model Product
func CreateProductForUser(table_name string) {

	// table_name = bersihkan string dari caracter @ ,"/ dll " , begitu juga di saat membuat  table
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	clean_table_name := reg.ReplaceAllString(table_name, "")

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS product_%s (
		id SERIAL PRIMARY KEY,              -- ID unik auto-increment
		name VARCHAR(100) NOT NULL,         -- Nama produk
		total_product INT DEFAULT 1          -- jumlah product default 1
 		expired  TIMESTAMPTZ DEFAULT NOT NULL,      -- batas akhir product 
		daily_reward TIME DEFAULT '12:00:00'  -- defaul Waktu kapan bisa di ambil reward ( defaul jam 12.00)
		created_at TIMESTAMP DEFAULT NOW(), -- Tanggal dibuat
		updated_at TIMESTAMP DEFAULT NOW()  -- Tanggal terakhir diubah
		FOREIGN KEY (product_id) REFERENCES products(id)
		);`, table_name)

	_, err := config.DB.Exec(query)
	if err != nil {
		log.Println("[x] gagal membuat table dengan nama : ", clean_table_name)
	}
}

/// [V]Ini function penting untuk update Table Product Dan di kirim ke user
/// Guankan untuk DI kirim ke User saat merequest , untuk Update table producs tampa mengubah table nya , hanya datanya
// func AddProducsByUserProdct()(*Product,error){
//
// 	var product []Product
//
//
// 	return nil, nil // natik saja
//
// }

// / daily_reward : dapatkan jika tidak ada maka 24jam
func AddProductForUser(table_name string, name_product string, total_product int16, daily_reward time.Time) {

	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	clean_table_name := reg.ReplaceAllString(table_name, "")

	input := ProductUser{
		Name:         name_product,
		TotalProduct: total_product,
		DailyReward: sql.NullTime{
			Time:  daily_reward,
			Valid: !daily_reward.IsZero(),
		},
	}
	query := fmt.Sprintf("INSERT INTO  %s (name, total_product ,expired ,daily_reward ) VALUES ( $1 , $2 , $3 , $4);", clean_table_name)
	err := config.DB.QueryRow(query, input.Name, input.TotalProduct, input.Expired, input.DailyReward)
	if err != nil {
		log.Printf("[x] gagal menambah  table %s , error : %v:", table_name, err)
		return
	}

}

// / !! gunakan jika memang user menghapus Accont
func DropTableForUser(table_name string) bool {
	query := fmt.Sprintf("DROP TABLE %s", table_name)
	_, err := config.DB.Exec(query)
	if err != nil {
		return false
	}
	return true
}

func GetAllProductForUser(table_name string) ([]ProductUser, error) {
	var userProducts []ProductUser

	// table_name = bersihkan string dari caracter @ ,"/ dll " , begitu juga di saat membuat databasenya
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	clean_table_name := reg.ReplaceAllString(table_name, "")

	query := fmt.Sprintf("SElECT * FROM %s", clean_table_name)
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("[x] gagal membaca product atas nama :", table_name, ":", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userProduct ProductUser
		err := rows.Scan(&userProduct.ID, &userProduct.Name, &userProduct.DailyReward, &userProduct.TotalProduct, &userProduct.Expired, &userProduct.CreatedAt, &userProduct.UpdatedAt)
		if err != nil {
			log.Println("[x  gagal mendapat data product user tertentu :", err)
			return nil, nil
		}
		userProducts = append(userProducts, userProduct)
	}
	return userProducts, nil

}

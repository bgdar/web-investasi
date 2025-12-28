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

type Product struct {
	ID int64 `json:"id"` // type lain untuk db : `db:"id" json:"id"`

	UserID int64 `db:"user_id" json:"user_id"`

	ProductID int64 `db:"product_id" json:"product_id"`

	Name         string       `json:"name"`
	TotalProduct int16        `json:"total_product"`
	Expired      time.Time    `json:"expired"`
	DailyReward  sql.NullTime `json:"daily_reward"`
	CreatedAt    time.Time    `json:"Created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// / daily_reward : dapatkan jika tidak ada maka 24jam

func AddProduct(nameProduct string, totalProduct int16, dailyReward time.Time, username string) {
	// Dapatkan product_id
	productID, err := GetOneDataProducts[int64]("id", nameProduct)
	if err != nil {
		log.Println("[x](AddProductForUser) gagal mendapat product_id untuk:", nameProduct, err)
		return
	}
	// Dapatkan user_id
	userID, err := GetOneDataUser[int64]("id", username)
	if err != nil {
		log.Println("[x](AddProductForUser) gagal mendapat user_id untuk:", username, err)
		return
	}

	// // sanitize table name
	// reg := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	// cleanTable := reg.ReplaceAllString(tableName, "")

	// jika daily_reward = zero, set default 24 jam
	if dailyReward.IsZero() {
		dailyReward = time.Now().Add(24 * time.Hour)
	}

	input := Product{
		Name:         nameProduct,
		TotalProduct: totalProduct,
		DailyReward: sql.NullTime{
			Time:  dailyReward,
			Valid: true,
		},
		ProductID: productID,
		UserID:    userID,
		Expired:   time.Now().Add(24 * time.Hour), // default 24 jam
	}

	query := `
	INSERT INTO %s (name, total_product, expired, daily_reward, user_id, product_id)
	VALUES ($1, $2, $3, $4, $5, $6); product`

	_, err = config.DB.Exec(query,
		input.Name,
		input.TotalProduct,
		input.Expired,
		input.DailyReward,
		input.UserID,
		input.ProductID,
	)

	if err != nil {
		log.Printf("[x] (AddProductForUser) gagal menambah tabel ID %d, error: %v", productID, err)
		return
	}

	log.Println("[✓] berhasil menambah product untuk user:", username)
}

// / !! gunakan jika memang user menghapus Accont
func DropProduct(table_name string) bool {
	query := fmt.Sprintf("DROP TABLE %s", table_name)
	_, err := config.DB.Exec(query)
	if err != nil {
		return false
	}
	return true
}

func GetAllProduct(table_name string) ([]Product, error) {
	var userProducts []Product

	// table_name = bersihkan string dari caracter @ ,"/ dll " , begitu juga di saat membuat databasenya

	reg := regexp.MustCompile(`[^a-zA-Z0-9_]+`)

	clean_table_name := reg.ReplaceAllString(table_name, "")

	query := fmt.Sprintf("SElECT * FROM product_%s", clean_table_name)
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("[x](GetAllProductForUser) gagal membaca product atas nama :", table_name, ":", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userProduct Product
		err := rows.Scan(&userProduct.ID, &userProduct.Name, &userProduct.DailyReward, &userProduct.TotalProduct, &userProduct.Expired, &userProduct.CreatedAt, &userProduct.UpdatedAt)
		if err != nil {
			log.Println("[x ] (GetAllProductForUser) gagal mendapat data product user tertentu :", err)
			return nil, nil
		}
		userProducts = append(userProducts, userProduct)
	}
	return userProducts, nil

}

package model

import (
	"database/sql"
	"log"
	"web-investasi/config"
)

// json agar bentuk format jsonya agar tidak mengikuti default bawan golang
type User struct { // struct ini utnuk desain colom di DB nya 
	ID int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func AddUser(name, password, email string) {
	// Buat struct user
	input := User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	// Query untuk insert dan mengembalikan id user yang baru
	query := "INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id"

	// Gunakan Scan untuk membaca id yang dihasilkan oleh RETURNING
	err := config.DB.QueryRow(query, input.Name, input.Password, input.Email).Scan(&input.ID)
	if err != nil {
		log.Println("[x] Gagal menambah user baru:", err)
		return
	}

}

func GetUserByID(id int64) (*User, error) {
	query := "SELECT id, name, password, email FROM users WHERE id = $1"

	row := config.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[!] User dengan ID %d tidak ditemukan\n", id)
			return nil, nil
		}
		log.Printf("[x] Gagal mengambil user dengan ID %d: %v\n", id, err)
		return nil, err
	}
	return &user, nil
}



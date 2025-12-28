package model

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"web-investasi/config"
)

// json agar bentuk format jsonya agar tidak mengikuti default bawan golang
type User struct { // struct ini utnuk desain colom di DB nya
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Saldo    string `json:"saldo"`
}

func AddUser(name, password, email string) error {
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
		return nil
	}
	return nil
}

// / hapus user
func DelUser(id uint64) {

	query := "DELETE FROM user WHERE id = $1"
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Printf("[x] gagal mendelete user dengan id %d", id)
	}

}

func GetUserByName(name string) (*User, error) {
	query := "SELECT  id , name , password , email FROM users WHERE name = $1"

	row := config.DB.QueryRow(query, name)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[!] User dengan Name %s tidak ditemukan\n", name)
			return nil, nil
		}
		log.Printf("[x] Gagal mengambil user dengan ID %s: %v\n", name, err)
		return nil, err
	}
	return &user, nil

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

// cek user apakah ada di table
func IsUserExists(name string, password string) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 
            FROM users 
            WHERE name = $1 AND password = $2
        )
    `
	var exists bool
	err := config.DB.QueryRow(query, name, password).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// / ambil 1 data denagn katagory data apapun di table
// / T = type  yang akan di cari
func GetOneDataUser[T any](field string, value any) (T, error) {
	var result T

	// sanitize: hanya boleh huruf/underscore
	valid := regexp.MustCompile(`^[a-zA-Z_]+$`)
	if !valid.MatchString(field) {
		return result, fmt.Errorf("invalid column name")
	}

	query := fmt.Sprintf("SELECT %s FROM users WHERE %s = $1 LIMIT 1", field, field)

	err := config.DB.QueryRow(query, value).Scan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

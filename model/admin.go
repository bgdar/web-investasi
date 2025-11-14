package model

import (
	"log"
	"web-investasi/config"
)

type Admin struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func AddAdmin(name string, password string, email string) {

	input := Admin{
		Name:     name,
		Password: password,
		Email:    email,
	}

	query := "INSERT INTO admin (name, password, email) VALUES ($1, $2, $3) RETURNING id"

	// Gunakan Scan untuk membaca id yang dihasilkan oleh RETURNING
	err := config.DB.QueryRow(query, input.Name, input.Password, input.Email).Scan(&input.ID)
	if err != nil {
		log.Println("[x] Gagal menambah admin baru:", err)
		return
	}
}

package config

import (
	"database/sql"
	"fmt"
	"log"
	  _ "github.com/lib/pq" // <- wajib, agar driver PostgreSQL terdaftar di database/sql
)

var DB *sql.DB 

/// function untuk handle koneksi ke database (pogresSql)
func ConnectDatabase()  {

	host := "localhost"
	port := 5432
  user := "postgres" // default login login ke pogresSql
	password := "bebas" // karena sedang menggunakan superuser ( postgres)
	dbname := "web_investasi"

	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db , err := sql.Open("postgres",psqlInfo)
	if err != nil  {
		log.Fatalln("[x] gagal membuka koneksi ke database :", err)
		
	}
	// uji koneksi 
	err = db.Ping(); if err != nil {
		log.Fatalf("[x] Database tidak bisa diakses: %v", err)
	}

	// simpan koneksi  
	DB = db
	
}







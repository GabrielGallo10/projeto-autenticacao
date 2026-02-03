package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var db *sql.DB

func connectDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database/users.db")
	
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao testar conexão: ", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
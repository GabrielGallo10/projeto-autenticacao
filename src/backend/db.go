package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var db *sql.DB

// Método para conectar ao banco de dados SQLite3
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

	createTable()
	log.Println("Banco de dados conectado com sucesso!")
}

// Método para criação da tabela do banco de dados, necessário pois o arquivo users.db é ignorado no git
func createTable() {
	query := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT UNIQUE NOT NULL, password TEXT NOT NULL);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar tabela users: ", err)
	}
}
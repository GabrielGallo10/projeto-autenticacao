package main

import (
	"log"
	"net/http"
)

func main() {
	connectDB()
	
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	log.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
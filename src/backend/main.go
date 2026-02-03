package main

import (
	"log"
	"net/http"
)

func main() {
	connectDB()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API rodando!"))
	})
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	log.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
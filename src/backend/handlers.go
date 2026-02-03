package main

import (
	"encoding/json"
    "net/http"
)

type Usuario struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}

	hashedPassword := hashPassword(user.Password)
	_, err = db.Exec("INSERT INTO users(name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Erro ao salvar usuário: usuário já cadastrado", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário cadastrado com sucesso!"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
        return
	}

	row := db.QueryRow("SELECT password FROM users WHERE email = ?", user.Email)
	var storedHash string
	err = row.Scan(&storedHash)
	if err != nil {
		http.Error(w, "Email ou senha inválidos", http.StatusUnauthorized)
        return
	}

	if !checkPassword(storedHash, user.Password) {
		http.Error(w, "Email ou senha inválidos", http.StatusUnauthorized)
        return
	}

	w.Write([]byte("Login realizado com sucesso!"))
}
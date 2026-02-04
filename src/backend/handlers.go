package main

import (
	"encoding/json"
    "net/http"
)

// Structs
type RegisterUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Endpoint de registro
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}

	var user RegisterUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Dados inválidos!", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Campo nome é obrigatório!", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Campo email é obrigatório!", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Campo senha é obrigatório!", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Senha deve ter pelo menos 6 caracteres!", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Erro interno!", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users(name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Usuário já cadastrado!", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário registrado com sucesso!"))
}

// Endpoint de login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido!", http.StatusMethodNotAllowed)
		return
	}

	var user LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Dados inválidos!", http.StatusBadRequest)
        return
	}

	if user.Email == "" {
		http.Error(w, "Campo email é obrigatório!", http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		http.Error(w, "Campo senha é obrigatório!", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT password FROM users WHERE email = ?", user.Email)
	var storedHash string
	err = row.Scan(&storedHash)
	if err != nil || !checkPassword(storedHash, user.Password){
		http.Error(w, "Email ou senha incorretos!", http.StatusUnauthorized)
        return
	}

	w.Write([]byte("Login realizado com sucesso!"))
}


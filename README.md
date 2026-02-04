# Projeto de Autenticação em Go

## Descrição

Este projeto implementa um **backend de autenticação** para aplicações web, fornecendo endpoints para **registro** e **login** de usuários. O objetivo é criar uma base segura de gerenciamento de usuários, com senhas hashadas, validação de dados e comunicação simples com frontend.

---

## Funcionalidades

* Cadastro de usuários com validação de:

  * Nome, Senha e Email obrigatório
  * Email com formato válido
  * Senha mínima de 6 caracteres
* Login de usuários com verificação de senha
* Senhas armazenadas de forma segura usando **bcrypt**

---

## Tecnologias

* **Backend:** Go
* **Frontend:** HTML, CSS e JS
* **Banco de dados:** SQLite
* **Bibliotecas:**

  * `github.com/mattn/go-sqlite3` → driver SQLite
  * `golang.org/x/crypto/bcrypt` → para hash de senhas

---

## Estrutura do Projeto

```
src
├── backend
│   ├── main.go        # Inicializa o servidor, conecta ao banco de dados e registra endpoints
│   ├── db.go          # Conexão com banco de dados SQLite
│   ├── handlers.go    # Endpoints de login e registro e validações básicas
│   └── auth.go        # Funções de autenticação, hash de senha
│
├── database
│   └── users.db       # Banco de dados SQLite (não versionado)
│
├── frontend
│   ├── js
│   │   ├── login.js       # Lógica de login e comunicação com /login
│   │   └── register.js    # Lógica de cadastro e comunicação com /register
│   │
│   ├── pages
│   │   ├── register.html  # Página de cadastro de usuários
│   │   └── dashboard.html # Página principal após login
│   │
│   ├── style
│   │   └── style.css      # Estilos globais do frontend
│   │
│   └── index.html         # Página inicial / login
│
├── .gitignore          # Ignora arquivos desnecessários 
├── go.mod              # Módulos do Go (dependências)
├── go.sum              
└── README.md           
```

---

## Como rodar o projeto

1. **Clone o repositório**

2. **Instale dependências do Go**

3. **Compile e rode o servidor**

* O backend ficará disponível em: `http://localhost:8080`

4. **Frontend**:

* Abra o HTML no navegador ou use um **servidor local**

---

## Endpoints da API

| Método | Endpoint  | Descrição                  |
| ------ | --------- | -------------------------- |
| POST   | /register | Cadastro de novo usuário   |
| POST   | /login    | Login de usuário existente |

* **Body JSON esperado para registro:**

```json
{
  "name": "Seu Nome",
  "email": "email@exemplo.com",
  "password": "123456"
}
```

* **Body JSON esperado para login:**

```json
{
  "email": "email@exemplo.com",
  "password": "123456"
}
```

---

## Observações

* Arquivo `users.db` **não deve ser versionado**, adicionei ao `.gitignore`.
* Senhas são armazenadas **somente de forma hash** (não armazenar texto puro).


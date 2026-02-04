# Projeto de Autenticação em Go(Golang)

## Descrição

É um **sistema completo de autenticação web**, oferecendo registro e login de usuários com **segurança e feedback visual moderno**.  
O projeto combina:

- Backend em **Go** com autenticação segura (senhas hashadas)  
- Frontend responsivo com **HTML, CSS e JavaScript**  
- Feedback de erros **inline e via toast** para uma melhor experiência do usuário

---

## Funcionalidades

### Cadastro de Usuário
- Validações no frontend e backend:
  - Nome obrigatório  
  - Email obrigatório
  - Senha obrigatória com mínimo de 6 caracteres  
- Mensagens de erro **inline** abaixo dos campos  
- Notificação **toast** para erros globais (ex.: “Usuário já cadastrado”)  
- Senhas **hashadas com bcrypt** antes de salvar no banco  

### Login de Usuário
- Verificação de email e senha  
- Mensagens de erro inline e via toast:
  - “Email ou senha incorretos”  
  - Campos obrigatórios destacados  
- Redirecionamento para **dashboard** após login bem-sucedido  

### Frontend Moderno
- HTML, CSS e JS
- Toasts para mensagens globais  
- Validação em tempo real e feedback visual  
- Layout responsivo  

---

## Tecnologias

- **Backend:** Go  
- **Frontend:** HTML, CSS, JavaScript
- **Banco de dados:** SQLite  
- **Bibliotecas Go:**
  - `github.com/mattn/go-sqlite3` – driver SQLite  
  - `golang.org/x/crypto/bcrypt` – hash de senhas  
- **Bibliotecas Frontend:**
  - FontAwesome – ícones  
  - Google Fonts – tipografia Montserrat  

---

## Estrutura do Projeto

```
src
├── backend
│ ├── main.go # Inicializa servidor, conecta ao DB e registra endpoints
│ ├── db.go # Conexão SQLite e inicialização de tabelas
│ ├── handlers.go # Endpoints /register e /login com validação
│ └── auth.go # Funções de autenticação e hash de senha
│
├── database
│ └── users.db # Banco SQLite (não versionado)
│
├── frontend
│ ├── js
│ │ ├── register.js # Lógica de cadastro, fetch e validações
│ │ ├── login.js # Lógica de login, fetch e feedback de erros
│ │ └── functions.js # Funções utilitárias (toast, validação inline)
│ │
│ ├── pages
│ │ ├── register.html # Página de cadastro
│ │ └── dashboard.html # Página após login
│ │
│ ├── style
│ │ ├── reset.css # Reset CSS
│ │ ├── common.css # Estilos globais
│ │ └── register.css # Estilos específicos do registro
│ │
│ └── index.html # Página inicial / login
│
├── .gitignore # Ignora DB e arquivos desnecessários
├── go.mod # Módulos do Go
├── go.sum
README.md         
```

---

## Como rodar o projeto

1. **Clone o repositório**

* git clone <URL_DO_REPOSITÓRIO>

* cd <PASTA_DO_PROJETO>

2. **Instale dependências do Go**

* go mod tidy

3. **Compile e rode o servidor**

* cd src

* go run backend/main.go backend/db.go backend/hanlers.go backend/auth.go

* O backend ficará disponível em: `http://localhost:8080`

4. **Frontend**:

* Use um **servidor local**, por exemplo o Live Server. Pois para que os módulos JS funcionem (import/export), o HTML deve ser servido via HTTP, não diretamente com file://.

---

## Endpoints da API

### Cadastro de Usuário (`POST /register`)

**URL:**  
`POST http://localhost:8080/register`

**Body JSON esperado:**
```json
{
  "name": "Seu Nome",
  "email": "email@exemplo.com",
  "password": "123456"
}
```

**Respostas possíveis:**
| Status | Mensagem                              | Descrição                   |
| ------ | ------------------------------------- | --------------------------- |
| 200    | Cadastro realizado com sucesso        | Usuário criado corretamente |
| 400    | Campo nome é obrigatório              | Nome não informado          |
| 400    | Campo email é obrigatório             | Email não informado         |
| 400    | Campo senha é obrigatório             | Senha não informada         |
| 400    | Senha deve ter no mínimo 6 caracteres | Senha muito curta           |
| 409    | Usuário já cadastrado                 | Email já existe no banco    |

### Login de Usuário (`POST /login`)

**URL:**  
`POST http://localhost:8080/login`

**Body JSON esperado:**
```json
{
  "email": "email@exemplo.com",
  "password": "123456"
}
```

**Respostas possíveis:**
| Status | Mensagem                              | Descrição                               |
| ------ | ------------------------------------- | --------------------------------------- |
| 200    | Login realizado com sucesso           | Login realizado com sucesso             |
| 400    | Campo email é obrigatório             | Email não informado                     |
| 400    | Campo senha é obrigatório             | Senha não informada                     |
| 401    | Email ou senha incorretos             | Email inexistente ou senha incorreta    |

---

## Observações

* Arquivo `users.db` **não deve ser versionado**, adicionei ao `.gitignore`.
* Senhas são armazenadas **somente de forma hash** (não armazenar texto puro).


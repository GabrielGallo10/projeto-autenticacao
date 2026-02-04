import { mensagemValidacao, mostrarToast } from "./functions.js";

const loginForm = document.getElementById("login-form");
const messageEmail = document.getElementById("mensagem-email");
const messageSenha = document.getElementById("mensagem-senha");
const messageDadosIncorretos = document.getElementById("mensagem-dados");

loginForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const email = document.getElementById("email");
    const password = document.getElementById("password");

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ email: email.value, password: password.value })
    })
    .then(response => response.text())
    .then(data => {
        mensagemValidacao(data, "email é obrigatório", messageEmail, email);

        mensagemValidacao(data, "senha é obrigatório", messageSenha, password);

        if (data.includes("incorretos")) {
            messageDadosIncorretos.textContent = data;
            messageDadosIncorretos.classList.add("mensagem-ativo");
            email.classList.add("input-error_ativo");
            password.classList.add("input-error_ativo");

            setTimeout(() => {
                messageDadosIncorretos.classList.remove("mensagem-ativo");
                email.classList.remove("input-error_ativo");
                password.classList.remove("input-error_ativo");
            }, 3000);
        }

        if (data.includes("não permitido") || data.includes("inválidos")) {
            mostrarToast(data);
        }

        if (data.includes("sucesso")) { 
            window.location.href = "pages/dashboard.html";
        }
    })
    .catch(error => {
        console.error(error);
        mostrarToast("Erro ao conectar com o servidor");
    });
});
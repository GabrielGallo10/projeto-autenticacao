import { mensagemValidacao, mostrarToast } from "./functions.js";

const registerForm = document.getElementById("register-form");
const messageNome = document.getElementById("mensagem-nome");
const messageEmail = document.getElementById("mensagem-email");
const messageSenha = document.getElementById("mensagem-senha");

registerForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const name = document.getElementById("name");
    const email = document.getElementById("email");
    const password = document.getElementById("password");

    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name: name.value, email: email.value, password: password.value })
    })
    .then(response => response.text())
    .then(data => {
        mensagemValidacao(data, "nome é obrigatório", messageNome, name);

        mensagemValidacao(data, "email é obrigatório", messageEmail, email);

        mensagemValidacao(data, "senha é obrigatório", messageSenha, password);

        mensagemValidacao(data, "menos 6 caracteres", messageSenha, password);

        if (data.includes("inválidos") || data.includes("interno") || data.includes("já cadastrado")) { 
            mostrarToast(data); 
        }
        

        if (data.includes("sucesso")) {
            window.location.href = "../index.html";
        }
    })
    .catch(error => {
        console.error(error);
        mostrarToast("Erro ao conectar com o servidor");
    });
});


const loginForm = document.getElementById("login-form");
const message = document.getElementById("message");

loginForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ email, password })
    })
    .then(response => response.text())
    .then(data => {
        message.textContent = data;

        if (data.includes("sucesso")) { 
            window.location.href = "pages/dashboard.html";
        }
    })
    .catch(error => {
        console.error(error);
        message.textContent = "Erro ao conectar com o servidor";
    });
});
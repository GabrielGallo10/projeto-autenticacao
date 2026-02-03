const registerForm = document.getElementById("register-form");
const message = document.getElementById("message");

registerForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name, email, password })
    })
    .then(response => response.text())
    .then(data => {
        message.textContent = data;

        if (data.includes("sucesso")) {
            window.location.href = "dashboard.html";
        }
    })
    .catch(error => {
        console.error(error);
        message.textContent = "Erro ao conectar com o servidor";
    });
});
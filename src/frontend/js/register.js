const registerForm = document.getElementById("register-form");
const message = document.getElementById("message");

registerForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    if (password.length < 6) {
        message.textContent = "Senha com menos de 6 caracteres.";
        return;
    }

    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: { 
            "Content-Type": "application/json" 
        },
        body: JSON.stringify({ 
            name: name, 
            email: email, 
            password: password
        })
    })
    .then(response => response.text())
    .then(data => {
        message.textContent = data;
        if (data.includes("sucesso")) {
            window.location.href = "dashboard.html";
        }
    })
    .catch(err => {
        console.error(err);
        message.textContent = "Erro ao cadastrar usuário!";
    });
});
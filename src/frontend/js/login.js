const loginForm = document.getElementById("login-form");
const message = document.getElementById("message");

loginForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:8080/login", {
        method: "POST",
        header: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            email: email, 
            password: password
        })
    }).then(response => {
        if (!response.ok) {
            throw new Error("Email ou senha inválidos");
        }
        return response.text();
    }).then(data => {
        message.textContent = data;
        if (data.includes("sucesso")) {
            window.location.href = "pages/dashboard.html";
        }
    }).catch(error => {
        message.textContent = error.message;
    });
});
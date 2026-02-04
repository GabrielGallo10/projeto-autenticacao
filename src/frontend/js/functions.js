const toast = document.getElementById("toast");

export function mensagemValidacao(data, string, messageContainer, inputContainer) {
    try {
        if (data.includes(string)) {
            messageContainer.textContent = data;
            messageContainer.classList.add("mensagem-ativo");
            inputContainer.classList.add("input-error_ativo");

            setTimeout(() => {
                messageContainer.classList.remove("mensagem-ativo");
                inputContainer.classList.remove("input-error_ativo");
            }, 3000);
        }
    } catch(e) {
        return console.log("Error: " + e);
    }
}

export function mostrarToast(message) {
    toast.textContent = message;
    toast.classList.add("mostrar");

    setTimeout(() => {
        toast.classList.remove("mostrar");
    }, 3000);
}
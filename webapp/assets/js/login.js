document.getElementById('login').addEventListener('submit', fazerLogin);

function fazerLogin(event) {
    event.preventDefault(); // Evita o comportamento padrão de envio do formulário

    const dataToSend = new URLSearchParams();
    dataToSend.append('email', document.getElementById('email').value);
    dataToSend.append('senha', document.getElementById('senha').value);

    fetch('/login', {
        method: 'POST',
        body: dataToSend
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erro HTTP! status: ${response.status}`);
            }
            console.log('Requisição POST bem-sucedida');
            window.location = "/home";
        })
        .catch(erro => {
            console.log('Erro na requisição:', erro);
            return erro.response;
        });
}   
document.getElementById('formulario-cadastro').addEventListener('submit', criarUsuario);

function criarUsuario(event) {
  event.preventDefault(); // Evita o comportamento padrão de envio do formulário

  const senha = document.getElementById('senha').value;
  const confirmarSenha = document.getElementById('confirmar-senha').value;

  if (senha !== confirmarSenha) {
    alert("Senha nao coincidem!");
    return;
  }

  const dataToSend = new URLSearchParams();
  dataToSend.append('nome', document.getElementById('nome').value);
  dataToSend.append('email', document.getElementById('email').value);
  dataToSend.append('nick', document.getElementById('nick').value);
  dataToSend.append('senha', document.getElementById('senha').value);

  fetch('/usuarios', {
    method: 'POST',
    body: dataToSend
  })
  .then(response => {
    if (!response.ok) {
      throw new Error(`Erro HTTP! status: ${response.status}`);
    }
    console.log('Requisição POST bem-sucedida');
    return response.json(); // parse the response body as JSON
  })
  .catch(erro => {
    console.log('Erro na requisição:', erro);
    return erro.response;
  });
}
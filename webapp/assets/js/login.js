$('#login').on('submit',fazerLogin);

function fazerLogin(evento){
    evento.preventDefault();

    $.ajax({
        url:"/login",
        method:"POST",
        data: {
            email:$('#email').val(),
            senha:$('#senha').val(),
        }
    }).done(function(){
        console.log("Teste")
        //window.location = "/home"    
    }).fail(function(erro){
        console.log(erro)
        alert("Usuario ou senha invalidos")
    });
}
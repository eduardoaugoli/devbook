$('#login').on('submit',fazerLogin);

function fazerLogin(){
    EventCounts.preventDefault();

    $.ajax({
        url:"/login",
        method:"POST",
        data: {
            email:$('#email').val(),
            senha:$('#senha').val(),
        }
    }).done(function () {
        window.location = "/home";    
    }).fail(function(erro){
        alert("Usuario ou senha invalidos")
    });
}
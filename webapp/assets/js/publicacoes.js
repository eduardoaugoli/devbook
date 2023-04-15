$(`#nova-publicacao`).on(`submit`,criarPublicacao);

function criarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        type: "POST",
        url: "/publicacao",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        },  
    }).done(function(){
        alert("Publicacao Criada com sucesso!")
    }).fail(function(){
        alert("Erro ao criar a publicacao")
    });

}
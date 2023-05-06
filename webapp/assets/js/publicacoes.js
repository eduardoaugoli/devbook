$(`#nova-publicacao`).on(`submit`,criarPublicacao);
$(document).on('click','.curtir-publicacao',curtirPublicacao);
$(document).on('click','.descurtir-publicacao',descurtirPublicacao);

function criarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        type: "POST",
        url: "/publicacoes",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        },  
    }).done(function(){
        window.location = "/home";
  
    }).fail(function(){
        alert("Erro ao criar a publicacao");
    });

}


function curtirPublicacao(evento) { 
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled',true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"

    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
        
        elementoClicado.addClass(`descurtir-publicacao`);
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass(`curtir-publicacao`);


    }).fail(function(){
        alert("Erro ao curtir");
    }).always(function(){
        elementoClicado.prop('disabled',false);
    })

}

function descurtirPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled',true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"

    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);
        
        elementoClicado.removeClass(`descurtir-publicacao`);
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass(`curtir-publicacao`);


    }).fail(function(){
        alert("Erro ao curtir");
    }).always(function(){
        elementoClicado.prop('disabled',false);
    })

}
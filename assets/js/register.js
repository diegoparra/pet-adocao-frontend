$('#novo-usuario').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
           nome: $('#nome').val(),
           email: $('#email').val(),
           senha: $('#senha').val(),
           perfil: $('#perfil').val(),
        }
    }).done(function() {
      Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
      window.location = "/admin/cadastrar-usuario";
    }).fail(function() {
      Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    });
}

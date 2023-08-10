$("#login").on("submit", doLogin);

function doLogin(evento) {
  evento.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $("#email").val(),
      senha: $("#senha").val(),
    },
  })
    .done(function () {
      window.location = "/home/admin";
    })
    .fail(function () {
      Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    });
}

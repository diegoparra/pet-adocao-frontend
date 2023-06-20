$("#novo-pet").on("submit", createPet);
$("#editar-pet").on("submit", editarPet);

function createPet(evento) {
  evento.preventDefault();

  var jform = new FormData();
  jform.append("nome", $("#nome").val());
  jform.append("telefone", $("#telefone").val());
  jform.append("especie", $("#especie").val());
  jform.append("genero", $("#genero").val());
  jform.append("porte", $("#porte").val());
  jform.append("idade", $("#idade").val());
  jform.append("vacinado", $("#vacinado").val());
  jform.append("castrado", $("#castrado").val());
  jform.append("descricao", $("#descricao").val());
  jform.append("arquivo", $("#arquivo").get(0).files[0]); // Here's the important bit

  $.ajax({
    url: "/pet/cadastrar",
    type: "POST",
    data: jform,
    dataType: "json",
    mimeType: "multipart/form-data", // this too
    contentType: false,
    cache: false,
    processData: false,
  })
    .done(function () {
      Swal.fire("Sucesso!", "Pet cadastrado com sucesso", "success").then(
        function () {
          window.location = "/";
        }
      );
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao cadastrar Pet", "error");
    });
}

function editarPet(evento) {
  evento.preventDefault();
  $.ajax({
    url: "/pet/editar-animal/" + $("#id").val(),
    method: "PUT",
    data: {
      nome: $("#nome").val(),
      telefone: $("#telefone").val(),
      especie: $("#especie").val(),
      genero: $("#genero").val(),
      porte: $("#porte").val(),
      idade: $("#idade").val(),
      vacinado: $("#vacinado").val(),
      castrado: $("#castrado").val(),
      descricao: $("#descricao").val(),
      adotado: $("#adotado").val(),
    },
  })
    .done(function () {
      Swal.fire("Sucesso!", "Pet atualizado com sucesso!", "success").then(
        function () {
          window.location = "/users/admin";
        }
      );
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao atualizar o Pet!", "error");
    });
}

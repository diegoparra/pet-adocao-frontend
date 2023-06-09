$('#atualizar-senha').on('submit', atualizarSenha);
$('#editar-profile').on('submit', editProfile);
$('#editar-foto').on('submit', editPhoto);


function atualizarSenha(evento) {
  evento.preventDefault();

  if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
    Swal.fire("Ops...", "As senhas não coincidem!", "warning");
    return;
  }

  $.ajax({
    url: "/change-password",
    method: "POST",
    data: {
      atual: $('#senha-atual').val(),
      nova: $('#nova-senha').val()
    }
  }).done(function() {
    Swal.fire("Sucesso!", "A senha foi atualizada com sucesso!", "success")
      .then(function() {
        window.location = "/profile";
      })
  }).fail(function() {
    Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
  });
}

function editProfile(evento) {
  evento.preventDefault();

  $.ajax({
    url: "/edit-profile",
    method: "PUT",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      terapeuta: $('#show-terapeuta').val(),
      facebook: $('#facebook').val(),
      instagram: $('#instagram').val(),
      telefone: $('#telefone').val(),
    }
  }).done(function() {
    Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
      .then(function() {
        window.location = "/profile";
      });
  }).fail(function() {
    Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
  });
}


function editPhoto(evento) {
	evento.preventDefault();
 
	var jform = new FormData();
	jform.append('file', $('#file').get(0).files[0]); // Here's the important bit
 
	$.ajax({
 
	  url: "/edit-photo",
	  type: "PUT",
	  data: jform,
	  dataType: 'json',
	  mimeType: 'multipart/form-data', // this too
	  contentType: false,
	  cache: false,
	  processData: false,
	}).done(function() {
	  Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success")
		 .then(function() {
			window.location = "/profile";
		 })
 
	}).fail(function() {
	  Swal.fire("Ops...", "Erro ao atualizar o usuário!", "error");
	})
 }
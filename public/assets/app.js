// Pide confirmación antes de enviar formularios marcados con data-confirm.
document.addEventListener("submit", function (event) {
  var form = event.target;
  if (form.dataset.confirm && !window.confirm(form.dataset.confirm)) {
    event.preventDefault();
  }
});

// Doble clic sobre el título de una tarea para editarla.
document.addEventListener("dblclick", function (event) {
  var title = event.target.closest("[data-edit-url]");
  if (title) {
    window.location.href = title.dataset.editUrl;
  }
});

// Esc cancela la edición de una tarea.
document.addEventListener("keydown", function (event) {
  if (event.key === "Escape" && event.target.dataset.cancelUrl) {
    window.location.href = event.target.dataset.cancelUrl;
  }
});

document.addEventListener("DOMContentLoaded", function () {
  // Coloca el cursor al final del título que se está editando.
  var editInput = document.getElementById("edit-task-title");
  if (editInput) {
    editInput.focus();
    var end = editInput.value.length;
    editInput.setSelectionRange(end, end);
  }

  // Los avisos desaparecen solos al cabo de unos segundos.
  document.querySelectorAll(".flash").forEach(function (flash) {
    setTimeout(function () {
      flash.classList.add("flash-hidden");
    }, 3500);
  });
});

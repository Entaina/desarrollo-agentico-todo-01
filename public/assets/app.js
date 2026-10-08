// Pide confirmación antes de enviar formularios marcados con data-confirm.
document.addEventListener("submit", function (event) {
  var form = event.target;
  if (form.dataset.confirm && !window.confirm(form.dataset.confirm)) {
    event.preventDefault();
  }
});

// Muestra u oculta el formulario de edición bajo la tarea.
document.addEventListener("click", function (event) {
  var button = event.target.closest("[data-edit-toggle]");
  if (!button) {
    return;
  }
  var form = document.getElementById(button.dataset.editToggle);
  var open = form.hidden;
  form.hidden = !open;
  button.setAttribute("aria-expanded", String(open));
  if (open) {
    form.querySelector("input[type=text]").focus();
  }
});

// Pide confirmación antes de enviar formularios marcados con data-confirm.
document.addEventListener("submit", function (event) {
  var form = event.target;
  if (form.dataset.confirm && !window.confirm(form.dataset.confirm)) {
    event.preventDefault();
  }
});

// Muestra u oculta el formulario de edición de una tarea.
function setEditOpen(formId, open) {
  var form = document.getElementById(formId);
  var toggle = document.querySelector('[data-edit-toggle="' + formId + '"]');
  if (!form || !toggle) {
    return;
  }
  form.hidden = !open;
  toggle.setAttribute("aria-expanded", open ? "true" : "false");
  if (open) {
    var input = form.querySelector('input[name="title"]');
    input.focus();
    input.select();
  } else {
    toggle.focus();
  }
}

document.addEventListener("click", function (event) {
  var toggle = event.target.closest("[data-edit-toggle]");
  if (toggle) {
    var formId = toggle.dataset.editToggle;
    setEditOpen(formId, document.getElementById(formId).hidden);
    return;
  }

  var cancel = event.target.closest("[data-edit-cancel]");
  if (cancel) {
    setEditOpen(cancel.dataset.editCancel, false);
  }
});

// Escape cierra el formulario de edición abierto.
document.addEventListener("keydown", function (event) {
  if (event.key !== "Escape") {
    return;
  }
  var form = event.target.closest(".task-edit");
  if (form) {
    setEditOpen(form.id, false);
  }
});

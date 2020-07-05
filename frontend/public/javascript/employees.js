document.addEventListener("DOMContentLoaded", function () {
  renderEmployees();

  let add_button = document.querySelector("#add-employee-button");
  let add_modal = document.querySelector("#add-employee-modal");
  M.Modal.init(add_modal, {});
  add_button.addEventListener("click", function (e) {
    var instance = M.Modal.getInstance(add_modal);
    instance.open();
  });
});

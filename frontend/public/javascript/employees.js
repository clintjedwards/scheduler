function populatePositionSelector() {
  let positions = JSON.parse(localStorage.getItem("positions"));

  let selector = document.getElementById("position-selector");
  let html = `<option value="" disabled selected> Choose Positions </option>`;

  for (let [id, position] of Object.entries(positions)) {
    html += `<option value="${sanitizeHTML(id)}">${sanitizeHTML(
      position.primary_name
    )}`;
    if (position.secondary_name !== "") {
      html += ` (${sanitizeHTML(position.secondary_name)})`;
    }
    html += "</option>";
  }

  selector.innerHTML = html;

  let position_selector = document.querySelector("#position-selector");
  M.FormSelect.init(position_selector, {
    classes: "text-primary",
  });
}

document.addEventListener("DOMContentLoaded", function () {
  renderEmployees();

  let datepicker = document.querySelectorAll(".datepicker");
  M.Datepicker.init(datepicker, {
    format: "mm-dd-yyyy",
    setDefaultDate: true,
  });

  let add_button = document.querySelector("#add-employee-button");
  let add_modal = document.querySelector("#add-employee-modal");
  M.Modal.init(add_modal, {});
  add_button.addEventListener("click", function (e) {
    var instance = M.Modal.getInstance(add_modal);
    instance.open();
  });

  populatePositionSelector();
});

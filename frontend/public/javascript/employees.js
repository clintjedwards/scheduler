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

function addUnavailabilityField() {
  var new_unavail_field = document.createElement("div");
  new_unavail_field.classList.add("input-field");
  new_unavail_field.classList.add("suffix");

  var new_input_field = document.createElement("input");
  new_input_field.id = "icon_prefix";
  new_input_field.type = "text";

  var new_icon_field = document.createElement("i");
  new_icon_field.classList.add("material-icons");
  new_icon_field.classList.add("clickable");
  new_icon_field.innerText = "clear";
  new_icon_field.addEventListener("click", function (e) {
    removeUnavailabilityField(e.target);
  });

  new_unavail_field.appendChild(new_input_field);
  new_unavail_field.appendChild(new_icon_field);

  let list = document.getElementById("unavailability-list");
  list.appendChild(new_unavail_field);
}

function removeUnavailabilityField(e) {
  e.parentNode.remove();
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

  let add_unavailability_button = document.querySelector("#add-unavailability");
  add_unavailability_button.addEventListener("click", function (e) {
    addUnavailabilityField();
  });

  populatePositionSelector();
});

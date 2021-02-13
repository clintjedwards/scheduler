import SchedulerClient from "./scheduler_client.js";

function populatePositions() {
  let client = new SchedulerClient();

  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get("id");

  client.listPositions().then((positions) => {
    let positions_element = document.getElementById("positions");
    for (const [id, position] of Object.entries(positions)) {
      let text = position.primary_name;

      if (position.secondary_name.length !== 0) {
        text = text + " |" + position.secondary_name;
      }

      let option = new Option(text, id);

      positions_element.appendChild(option);
    }
  });
}

function toggle_collapse(element) {
  if (element.style.display === "block") {
    element.style.display = "none";
  } else {
    element.style.display = "block";
  }
}

function help_collapse() {
  let help_button = document.getElementById("show-help");
  help_button.addEventListener("click", function() {
    let help_text = document.getElementById("help-text");
    toggle_collapse(help_text);
  });
}

function add_unavail_field(count) {
  let times = document.getElementById("unavail_times");

  let new_field = document.createElement("input");
  new_field.setAttribute("type", "text");
  new_field.setAttribute("autocomplete", "off");
  new_field.setAttribute("id", "time-" + count);
  new_field.setAttribute(
    "class",
    "mb-2 text-2xl font-light border border-transparent mt-1 text-center shadow-sm border-gray-500 rounded-md"
  );

  times.appendChild(new_field);
  return count + 1;
}

function remove_unavail_field(count) {
  let times = document.getElementById("unavail_times");

  times.removeChild(times.lastChild);
  return count - 1;
}

function add_unavail_button_actions() {
  var count = 0;

  let add_button = document.getElementById("add-button");
  let remove_button = document.getElementById("remove-button");

  add_button.addEventListener("click", function() {
    count = add_unavail_field(count);
  });

  remove_button.addEventListener("click", function() {
    count = remove_unavail_field(count);
  });
}

document.addEventListener("DOMContentLoaded", function() {
  populatePositions();
  help_collapse();
  add_unavail_button_actions();
});

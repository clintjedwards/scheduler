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
        text = text + " | " + position.secondary_name;
      }

      let option = new Option(text, id);

      positions_element.appendChild(option);
    }
  });
}

document.addEventListener("DOMContentLoaded", function() {
  populatePositions();
});

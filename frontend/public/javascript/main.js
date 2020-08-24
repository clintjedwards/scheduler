function populatePositionList() {
  client.listPositions().then((data) => {
    localStorage.setItem("positions", JSON.stringify(data));
  });
}

function populateScheduleList() {
  client.listSchedules().then((data) => {
    localStorage.setItem("schedules", JSON.stringify(data.Schedules));
    localStorage.setItem("schedules_order", JSON.stringify(data.Order));
  });
}

function renderPositions() {
  client.listPositions().then((data) => {
    let positions = data;
    localStorage.setItem("positions", JSON.stringify(data));

    let content = document.getElementById("positions-content-body");

    let innerHTML = "";
    innerHTML += "<ul class='collection'>";

    for (const [id, position] of Object.entries(positions)) {
      innerHTML += `<li class="collection-item">
        <h5>${sanitizeHTML(position.primary_name)}</h5>
        <p class="grey-text text-darken-1">${sanitizeHTML(
          position.secondary_name
        )}</p>
        <p>${sanitizeHTML(position.description)}</p>
        </li>`;
    }
    innerHTML += "</ul>";

    content.innerHTML = innerHTML;
  });
}

document.addEventListener("DOMContentLoaded", function () {
  populateSystemInfo();
  populateEmployeeList();
  populatePositionList();
  populateScheduleList();

  var elems = document.querySelectorAll(".sidenav");
  M.Sidenav.init(elems, { menuWidth: 300 });
});

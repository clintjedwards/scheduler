function renderSchedules() {
  client.listSchedules().then((data) => {
    state.schedules = data.Schedules;
    state.schedules_order = data.Order;

    let content = document.getElementById("schedules-content-body");

    let innerHTML = "";
    innerHTML += "<ul id='schedules-collection' class='collection'>";

    for (let id of state.schedules_order) {
      innerHTML += `<li id="${id}" class="collection-item">
            <h6>${state.schedules[id].start} - ${state.schedules[id].end}</h6>
            </li>`;
    }
    innerHTML += "</ul>";

    content.innerHTML = innerHTML;

    document.querySelectorAll(".collection-item").forEach((item) => {
      item.addEventListener("click", function (e) {
        console.log(e.currentTarget.id);
      });
    });
  });
}

function renderSchedule(id) {
  client.getSchedule(id).then((schedule) => {
    let content = document.getElementById("view-schedule-modal-content");
    let innerHTML = "";
    innerHTML += `<h4>Schedule ${schedule.start} - ${schedule.end}</h4>
    <h6 class="grey-text text-darken-1">Created: ${humanizedBuildTime(
      schedule.created
    )} (${humanizedRelativeBuildTime(schedule.created)})</h6>
    `;
    content.innerHTML = innerHTML;
  });
}

document.addEventListener("DOMContentLoaded", function () {
  renderSchedules();
  var elem = document.querySelector(".modal");
  M.Modal.init(elem, {});
  var instance = M.Modal.getInstance(elem);
  instance.open();
  renderSchedule("C3PWd");
});

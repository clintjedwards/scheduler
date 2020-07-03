function renderSchedules() {
  client.listSchedules().then((data) => {
    if (data.Schedules === null || data.Order === null) {
      return;
    }
    state.schedules = data.Schedules;
    state.schedules_order = data.Order;

    let content = document.getElementById("schedules-content-body");

    let innerHTML = "";
    innerHTML += "<ul id='schedules-collection' class='collection'>";

    for (let id of state.schedules_order) {
      innerHTML += `<li id="${id}" class="collection-item">
            <h6>${humanizedDate(state.schedules[id].start)} - ${humanizedDate(
        state.schedules[id].end
      )}</h6>
            </li>`;
    }
    innerHTML += "</ul>";

    content.innerHTML = innerHTML;

    document.querySelectorAll(".collection-item").forEach((item) => {
      item.addEventListener("click", function (e) {
        var elem = document.querySelector(".modal");
        var instance = M.Modal.getInstance(elem);
        instance.open();
        renderSchedule(e.currentTarget.id);
      });
    });
  });
}

function renderSchedule(id) {
  client.getSchedule(id).then((schedule) => {
    let heading = document.getElementById("view-schedule-modal-heading");
    let innerHTML = "";
    innerHTML += `<div id="schedule-title"><h4>Schedule ${humanizedDate(
      schedule.start
    )} - ${humanizedDate(schedule.end)}</h4>
    <h6 class="grey-text text-darken-1">Created: ${humanizedBuildTime(
      schedule.created
    )} (${humanizedRelativeBuildTime(schedule.created)})</h6></div><br><br>
    `;

    innerHTML += generateHeadings(schedule.time_table);
    heading.innerHTML = innerHTML;

    let content = document.getElementById("view-schedule-modal-content");
    innerHTML = "";
    innerHTML += generateCalendar(schedule.time_table);
    content.innerHTML = innerHTML;
  });
}

function humanizedDate(date) {
  let momentObj = moment(date, "MM-DD-YYYY");
  let humanDate = momentObj.format("MMM Do");
  return humanDate;
}

document.addEventListener("DOMContentLoaded", function () {
  renderSchedules();
  var elem = document.querySelector(".modal");
  M.Modal.init(elem, {});
});

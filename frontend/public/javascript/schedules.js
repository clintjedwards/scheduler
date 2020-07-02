function renderSchedules() {
  client.listSchedules().then((data) => {
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
        console.log(e.currentTarget.id);
      });
    });
  });
}

function renderSchedule(id) {
  client.getSchedule(id).then((schedule) => {
    let content = document.getElementById("view-schedule-modal-content");
    let innerHTML = "";
    innerHTML += `<h4>Schedule ${humanizedDate(
      schedule.start
    )} - ${humanizedDate(schedule.end)}</h4>
    <h6 class="grey-text text-darken-1">Created: ${humanizedBuildTime(
      schedule.created
    )} (${humanizedRelativeBuildTime(schedule.created)})</h6>
    `;
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
  var instance = M.Modal.getInstance(elem);
  instance.open();
  renderSchedule("C3PWd");
});
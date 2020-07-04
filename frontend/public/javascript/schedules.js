// renderSchedules draws a list of schedules that have been generated
// and enables modal interaction for each
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

// renderSchedule draws a full schedule timetable
function renderSchedule(id) {
  client.getSchedule(id).then((schedule) => {
    // Schedule heading contains the dates for the schedule.
    // The heading needs to be outside of the modal content because modal-content
    // uses transform which prevents us from setting the heading to position fixed.
    let modal_header = document.getElementById("view-schedule-modal-heading");
    let innerHTML = `<div id="schedule-title"><h4>Schedule ${humanizedDate(
      schedule.start
    )} - ${humanizedDate(schedule.end)}</h4>
    <h6 class="text-secondary">Created: ${humanizedBuildTime(
      schedule.created
    )} (${humanizedRelativeBuildTime(schedule.created)})</h6></div><br><br>
    `;

    innerHTML += generateHeadings(schedule.time_table);
    modal_header.innerHTML = innerHTML;

    // Schedule content contains the times and scheduled employees for those times
    let modal_content = document.getElementById("view-schedule-modal-content");
    innerHTML = generateCalendar(schedule.time_table);
    modal_content.innerHTML = innerHTML;
  });
}

// humanizedDate takes a date in format 06-19-1990 and returns humanized date
// in format Jun 19
function humanizedDate(date) {
  let momentObj = moment(date, "MM-DD-YYYY");
  let humanDate = momentObj.format("MMMM Do");
  return humanDate;
}

document.addEventListener("DOMContentLoaded", function () {
  populateEmployeeList();
  populatePositionList();
  renderSchedules();
  var elem = document.querySelector(".modal");
  M.Modal.init(elem, {});
});

const client = new SchedulerClient();

employees = [];
schedules = [];
schedules_order = [];

function humanizedBuildTime(time) {
  let human_time = moment(moment.unix(time)).format("L");
  return human_time;
}

function humanizedRelativeBuildTime(time) {
  let human_time = moment(moment.unix(time)).fromNow();
  return human_time;
}

function populateSystemInfo() {
  systemInfo = client.getSystemInfo().then((info) => {
    let footer_text = `Version v${info.semver} | ${humanizedBuildTime(
      info.build_time
    )} (${humanizedRelativeBuildTime(info.build_time)}) | ${info.commit}`;

    if (info.debug_enabled) {
      footer_text += " | Debug Enabled";
    }

    let footer = document.getElementById("footer-text");
    footer.innerText = footer_text;
  });
}

function populateEmployeeList() {
  client.listEmployees().then((data) => {
    employees = data;
  });
}
function populatePositionList() {
  client.listPositions().then((data) => {
    positions = data;
  });
}
function populateScheduleList() {
  client.listSchedules().then((data) => {
    schedules = data.Schedules;
    schedules_order = data.Order;
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

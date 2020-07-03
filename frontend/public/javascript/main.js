const client = new SchedulerClient();

let state = {};

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
    state.employees = data;
  });
}
function populatePositionList() {
  client.listPositions().then((data) => {
    state.positions = data;
  });
}
function populateScheduleList() {
  client.listSchedules().then((data) => {
    state.schedules = data.Schedules;
    state.schedules_order = data.Order;
  });
}

function renderEmployees() {
  client.listEmployees().then((data) => {
    state.employees = data;

    let content = document.getElementById("employees-content-body");

    let innerHTML = "";
    innerHTML += "<ul class='collection'>";

    for (const [id, employee] of Object.entries(state.employees)) {
      innerHTML += `<li class="collection-item">
      <span class="title">${employee.name}</span>
      </li>`;
    }
    innerHTML += "</ul>";

    content.innerHTML = innerHTML;
  });
}

function renderPositions() {
  client.listPositions().then((data) => {
    state.positions = data;

    let content = document.getElementById("positions-content-body");

    let innerHTML = "";
    innerHTML += "<ul class='collection'>";

    for (const [id, position] of Object.entries(state.positions)) {
      innerHTML += `<li class="collection-item">
        <h5>${position.primary_name}</h5>
        <p class="grey-text text-darken-1">${position.secondary_name}</p>
        <p>${position.description}</p>
        </li>`;
    }
    innerHTML += "</ul>";

    content.innerHTML = innerHTML;
  });
}

document.addEventListener("DOMContentLoaded", function () {
  populateSystemInfo();

  var elems = document.querySelectorAll(".sidenav");
  M.Sidenav.init(elems, { menuWidth: 300 });
});

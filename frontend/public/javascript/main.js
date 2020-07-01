function humanizedBuildTime(time) {
  let human_time = moment(moment.unix(time)).format("L");
  return human_time;
}

function humanizedRelativeBuildTime(time) {
  let human_time = moment(moment.unix(time)).fromNow();
  return human_time;
}

function populateSystemInfo() {
  client = new SchedulerClient();
  systemInfo = client.getSystemInfo().then((info) => {
    let footer_text = `Version v${info.semver} | ${humanizedBuildTime(
      info.build_time
    )} (${humanizedRelativeBuildTime(info.build_time)}) | ${info.commit}`;

    if (info.debug) {
      footer_text += " | Debug Enabled";
    }

    let footer = document.getElementById("footer-text");
    footer.innerText = footer_text;
  });
}

function populateEmployees() {}

document.addEventListener("DOMContentLoaded", function () {
  populateSystemInfo();
});

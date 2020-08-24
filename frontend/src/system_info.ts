var utils = require("./utils");
var client = require("./scheduler_client");

var system_info = {
  populate: function () {
    systemInfo = client.getSystemInfo().then((info) => {
      let footer_text = `Version v${sanitizeHTML(
        info.semver
      )} | ${utils.humanizedBuildTime(
        info.build_time
      )} (${utils.humanizedRelativeBuildTime(
        info.build_time
      )}) | ${utils.sanitizeHTML(info.commit)}`;

      if (info.debug_enabled) {
        footer_text += " | Debug Enabled";
      }

      let footer = document.getElementById("footer-text");
      footer.innerText = footer_text;
    });
  },
};

module.exports = system_info;

<script>
  import { format, formatDistance, fromUnixTime } from "date-fns";
  import { client } from "../client.js";

  let systemInfo;

  function buildTimeDate(unixTime) {
    let date = format(fromUnixTime(unixTime), "MMMM do, y");
    return date;
  }

  function humanizedRelativeBuildTime(time) {
    let human_time = formatDistance(fromUnixTime(time), new Date(), {
      addSuffix: true,
    });
    return human_time;
  }

  client.getSystemInfo().then((info) => {
    let footer_text = `Version v${info.semver} | ${buildTimeDate(
      info.build_time
    )} (${humanizedRelativeBuildTime(info.build_time)}) | ${info.commit}`;

    if (info.debug) {
      footer_text += " | Debug Enabled";
    }

    systemInfo = footer_text;
  });
</script>

<footer class="mx-auto w-full mt-16 h-10 text-center">
  <hr class="border-t border-gray-500 mb-5" />
  <div class="text-gray-500 font-light">{systemInfo}</div>
</footer>

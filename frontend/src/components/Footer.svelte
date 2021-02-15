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

<footer>{systemInfo}</footer>

<style>
  footer {
    width: 100%;
    margin-top: 4em;
    height: 2.5rem;
    text-align: center;
    font-weight: 300;
    color: #6c757d;
  }
</style>

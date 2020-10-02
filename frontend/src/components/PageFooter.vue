<template>
  <footer>
    <div class="container">
      Version v{{ $store.state.app_info.semver }} | {{ humanizedBuildTime }} ({{
        humanizedRelativeBuildTime
      }}) |
      {{ $store.state.app_info.commit }}
      <template v-if="$store.state.app_info.debug_enabled"
        >| Debug Enabled</template
      >
    </div>
  </footer>
</template>

<script>
import * as moment from "moment";

export default {
  name: "PageFooter",
  computed: {
    humanizedBuildTime: function() {
      return moment(moment.unix(this.$store.state.app_info.build_time)).format(
        "L"
      );
      return build_time;
    },
    humanizedRelativeBuildTime: function() {
      return moment(
        moment.unix(this.$store.state.app_info.build_time)
      ).fromNow();
    },
  },
};
</script>

<style scoped>
.container {
  width: auto;
  max-width: 680px;
}
</style>

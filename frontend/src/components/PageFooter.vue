<template>
  <footer class="footer">
    <div class="page-footer-text fluid-container bg-dark text-secondary">
      Version v{{ $store.state.appInfo.semver }} | {{ humanizedBuildTime }} ({{
      humanizedRelativeBuildTime
      }}) | {{ $store.state.appInfo.commit }}
      <template
        v-if="$store.state.appInfo.debug_enabled"
      >| Debug Enabled</template>
    </div>
  </footer>
</template>

<script lang="ts">
import Vue from "vue";
import * as moment from "moment";

export default Vue.extend({
  data: function() {
    return {};
  },
  computed: {
    humanizedBuildTime: function() {
      let build_time = moment(
        moment.unix(this.$store.state.appInfo.build_time)
      ).format("L");

      return build_time;
    },
    humanizedRelativeBuildTime: function() {
      let build_time = moment(
        moment.unix(this.$store.state.appInfo.build_time)
      ).fromNow();

      return build_time;
    }
  }
});
</script>

<style scoped>
.footer {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 20px;
}

.page-footer-text {
  /* color: #e0e0e0; */
  padding-left: 5px;
  padding-right: 5px;
}
</style>

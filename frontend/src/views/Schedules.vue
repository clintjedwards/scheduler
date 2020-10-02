<template>
  <div>
    <h1>Schedules</h1>
    <b-list-group class="list">
      <b-list-group-item
        v-for="schedule_id in $store.state.schedules.order"
        :key="schedule_id"
        >{{ humanizeDate($store.state.schedules.schedules[schedule_id].start) }}
        -
        {{
          humanizeDate($store.state.schedules.schedules[schedule_id].end)
        }}</b-list-group-item
      >
    </b-list-group>
  </div>
</template>

<script>
import * as moment from "moment";

export default {
  components: {},
  data: function() {
    return {};
  },
  methods: {
    humanizeDate: function(date) {
      return moment(date, "MM/DD/YYYY").format("dddd, MMMM Do");
    },
  },
};
</script>

<style scoped>
h1 {
  padding-top: 60px;
}
.list {
  padding-top: 80px;
}
</style>

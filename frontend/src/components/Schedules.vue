<template>
  <div class="mt-5 mr-5 ml-5">
    <b-container>
      <b-row cols="1">
        <b-col cols="10">
          <h5 class="text-secondary">View and configure generated schedules.</h5>
        </b-col>
        <b-col cols="2">
          <button type="button" class="btn btn-dark">Generate Schedule</button>
        </b-col>
      </b-row>
    </b-container>
    <div class="content">
      <b-list-group>
        <b-list-group-item
          v-for="id in schedules.order"
          :key="id"
          class="flex-column align-items-start"
        >
          <div class="d-flex w-100 justify-content-between">
            <h5 class="mb-1">{{schedules.schedules[id].start}} - {{schedules.schedules[id].end}}</h5>
          </div>

          <p
            class="mb-1"
          >Donec id elit non mi porta gravida at eget metus. Maecenas sed diam eget risus varius blandit.</p>

          <small>Donec id elit non mi porta.</small>
        </b-list-group-item>
      </b-list-group>
    </div>
    <!-- <router-view></router-view> -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Schedules, Schedule } from "../scheduler_client";

export default Vue.extend({
  components: {},
  data: function() {
    return {
      schedules: {
        schedules: {} as Map<string, Schedule>,
        order: [] as string[]
      }
    };
  },
  mounted() {
    let schedules: Map<string, Schedule> = this.$store.state.schedules
      .Schedules;
    let order: string[] = this.$store.state.schedules.Order;

    this.schedules.schedules = schedules;
    this.schedules.order = order;
  },
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "setSchedules") {
        let schedules: Map<string, Schedule> = this.$store.state.schedules
          .Schedules;
        let order: string[] = this.$store.state.schedules.Order;

        this.schedules.schedules = schedules;
        this.schedules.order = order;
      }
    });
  }
});
</script>

<style scoped>
.content {
  margin-top: 5em;
}
</style>

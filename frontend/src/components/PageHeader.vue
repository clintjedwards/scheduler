<template>
  <div>
    <b-navbar class="justify-content-center" type="dark" variant="dark">
      <b-navbar-nav align="center">
        <b-navbar-brand href="/">
          <img src="/images/schedule_icon_light.svg" width="24" height="24" />
        </b-navbar-brand>
        <b-nav-item href="/employees">Employees</b-nav-item>
        <b-nav-item href="/positions">Positions</b-nav-item>
        <b-nav-item href="#">Schedules</b-nav-item>
        <b-nav-item href="#">Placeholder</b-nav-item>
        <b-nav-item href="#">Placeholder</b-nav-item>
      </b-navbar-nav>
    </b-navbar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { SchedulerClientWrapper } from "../scheduler_client_wrapper";

let client: SchedulerClientWrapper;
client = new SchedulerClientWrapper();

export default Vue.extend({
  methods: {}
});
</script>

<style scoped></style>

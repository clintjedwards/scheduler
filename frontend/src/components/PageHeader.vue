<template>
  <div>
    <b-navbar class="justify-content-center" type="dark" variant="dark">
      <b-navbar-nav align="center">
        <b-navbar-brand href="/">
          <img src="/images/schedule_icon_light.svg" width="24" height="24" />
        </b-navbar-brand>
        <b-nav-item to="/employees">Employees</b-nav-item>
        <b-nav-item to="/positions">Positions</b-nav-item>
        <b-nav-item to="/schedules">Schedules</b-nav-item>
        <b-nav-item href="#">Placeholder</b-nav-item>
        <b-nav-item href="#">Placeholder</b-nav-item>
      </b-navbar-nav>
    </b-navbar>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped></style>

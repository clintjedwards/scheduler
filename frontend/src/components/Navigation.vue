<template>
  <div>
    <div class="px-3 py-2">
      <h2>Scheduler</h2>
      <hr />
      <nav class="mb-3">
        <b-nav vertical>
          <b-nav-item tag="router-link" to="/" active>
            <b-icon icon="house"></b-icon> Home
          </b-nav-item>
          <b-nav-item tag="router-link" to="/employees">
            <b-icon icon="person-badge"></b-icon> Employees
          </b-nav-item>
          <b-nav-item tag="router-link" to="/positions">
            <b-icon icon="card-list"></b-icon> Positions
          </b-nav-item>
          <b-nav-item tag="router-link" to="/schedules">
            <b-icon icon="calendar2-day"></b-icon> Schedules
          </b-nav-item>
        </b-nav>
      </nav>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      visible: true,
      close: true,
      esc: true,
      change: true,
    };
  },
};
</script>

<style>
nav {
  padding-top: 80px;
}
</style>

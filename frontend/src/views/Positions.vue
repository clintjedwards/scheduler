<template>
  <div>
    <h1>Positions</h1>
    <b-list-group class="list">
      <b-list-group-item
        v-for="position in $store.state.positions"
        :key="position.id"
        >{{ position.primary_name }}
        <br />
        <span class="text-muted">{{
          position.secondary_name
        }}</span></b-list-group-item
      >
    </b-list-group>
  </div>
</template>

<script>
export default {
  components: {},
  data: function() {
    return {};
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

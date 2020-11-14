<template>
  <div>
    <h1>Employees</h1>
    <div id="add_employee_btn">
      <b-button variant="outline-primary" v-on:click="showAddEmployeeModal()">
        <b-icon icon="person-plus"></b-icon>
        Add Employee</b-button
      >
    </div>
    <b-list-group class="list" style="cursor: pointer;" flush>
      <b-list-group-item
        v-for="employee in $store.state.employees"
        :key="employee.id"
        :to="employee.id"
        append
        >{{ employee.name }}</b-list-group-item
      >
    </b-list-group>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  components: {},
  data: function() {
    return {};
  },
  methods: {
    showAddEmployeeModal: function() {
      this.$router.push({ name: "addEmployeeModal" });
    },
  },
};
</script>

<style scoped>
h1 {
  padding-top: 60px;
}
.list {
  padding-top: 100px;
}

#add_employee_btn {
  padding-top: 40px;
  padding-right: 180px;
  float: right;
}
</style>

<template>
  <div class="mt-5 mr-5 ml-5">
    <h5
      class="text-secondary"
    >Configure employee data, unavailability, eligible positions, and more.</h5>
    <div class="float-right pr-5">
      <button type="button" class="btn btn-dark">Add Employee</button>
    </div>
    <div class="content">
      <ul class="list-unstyled">
        <b-media v-for="employee in employees" :key="employee.id" tag="li">
          <template v-slot:aside>
            <b-img blank blank-color="#abc" width="64" alt="placeholder"></b-img>
          </template>
          <h5 class="mt-0 mb-1">{{employee.name}}</h5>
          <p class="mb-0">{{employee.notes}}</p>
          <br />
        </b-media>
      </ul>
    </div>
    <!-- <router-view></router-view> -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Employees, Employee } from "../scheduler_client";

export default Vue.extend({
  components: {},
  data: function() {
    return {
      employees: [] as Employee[]
    };
  },
  mounted() {
    let employees: { [key: string]: Employee } = this.$store.state.employees;
  },
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "setEmployees") {
        let employees: { [key: string]: Employee } = this.$store.state
          .employees;
        let newEmployeeList: Employee[] = [];
        for (const [key, value] of Object.entries(employees)) {
          newEmployeeList.push(value);
        }
        this.employees = newEmployeeList;
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

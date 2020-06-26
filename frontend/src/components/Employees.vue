<template>
  <div class="mt-5 mr-5 ml-5">
    <b-container>
      <b-row cols="1">
        <b-col cols="10">
          <h5
            class="text-secondary"
          >Configure employee data, unavailability, eligible positions, and more.</h5>
        </b-col>
        <b-col cols="2">
          <button type="button" class="btn btn-dark">Add Employee</button>
        </b-col>
      </b-row>
    </b-container>
    <div class="content">
      <b-list-group>
        <b-list-group-item
          v-for="employee in employees"
          :key="employee.id"
          class="flex-column align-items-start"
        >
          <h5 class="mt-0 mb-1">{{employee.name}}</h5>
          <p class="mb-0">{{employee.notes}}</p>
        </b-list-group-item>
      </b-list-group>
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
    let newEmployeeList: Employee[] = [];
    for (const [key, value] of Object.entries(employees)) {
      newEmployeeList.push(value);
    }
    this.employees = newEmployeeList;
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

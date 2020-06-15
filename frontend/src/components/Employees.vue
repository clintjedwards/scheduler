<template>
  <div class="mt-5 mr-5 ml-5">
    <h5 class="text-secondary">
      Configure employee data, unavailability, eligible positions, and more.
    </h5>
    <div class="float-right pr-5">
      <button type="button" class="btn btn-dark">Add Employee</button>
    </div>
    <div class="mt-5">
      <ul class="list-unstyled">
        <b-media v-for="employee in employeeList" :key="employee.id" tag="li">
          <template v-slot:aside>
            <b-img
              blank
              blank-color="#abc"
              width="64"
              alt="placeholder"
            ></b-img>
          </template>
          <h5 class="mt-0 mb-1">List-based media object</h5>
          <p class="mb-0">
            Cras sit amet nibh libero, in gravida nulla. Nulla vel metus
            scelerisque ante sollicitudin. Cras purus odio, vestibulum in
            vulputate at, tempus viverra turpis. Fusce condimentum nunc ac nisi
            vulputate fringilla. Donec lacinia congue felis in faucibus.
          </p>
        </b-media>
      </ul>
    </div>
    <!-- <router-view></router-view> -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Employee } from "../scheduler_message_pb";

export default Vue.extend({
  components: {},
  data: function() {
    return {
      employeeList: [] as Employee.AsObject[],
    };
  },
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "setEmployees") {
        this.employeeMapToList();
      }
    });
  },
  methods: {
    employeeMapToList: function() {
      let employeeDataMap: { [key: string]: Employee } = this.$store.state
        .employeeData;
      let newEmployeeList: Employee.AsObject[] = [];

      for (const [key, value] of Object.entries(employeeDataMap)) {
        newEmployeeList.push(value.toObject());
      }
      this.employeeList = newEmployeeList;
    },
  },
});
</script>

<style scoped></style>

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
        <b-media v-for="employee in employeeList" :key="employee.id" tag="li">
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
import { Schedules, Schedule } from "../scheduler_client";

// export default Vue.extend({
//   components: {},
//   data: function() {
//     return {
//       positions: [] as Position[]
//     };
//   },
//   mounted() {
//     this.positions = this.$store.state.positions;
//   },
//   created() {
//     this.$store.subscribe((mutation, state) => {
//       if (mutation.type === "setPositions") {
//         let positions: { [key: string]: Position } = this.$store.state
//           .positions;
//         let newPositionList: Position[] = [];
//         for (const [key, value] of Object.entries(positions)) {
//           newPositionList.push(value);
//         }
//         this.positions = newPositionList;
//       }
//     });
//   }
// });
</script>

<style scoped>
.content {
  margin-top: 5em;
}
</style>

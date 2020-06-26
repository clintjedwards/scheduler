<template>
  <div class="mt-5 mr-5 ml-5">
    <b-container>
      <b-row cols="1">
        <b-col cols="10">
          <h5 class="text-secondary">Configure employee positions.</h5>
        </b-col>
        <b-col cols="2">
          <button type="button" class="btn btn-dark">Add Position</button>
        </b-col>
      </b-row>
    </b-container>
    <div class="content">
      <b-list-group>
        <b-list-group-item
          v-for="position in positions"
          :key="position.id"
          class="flex-column align-items-start"
        >
          <div class="d-flex w-100 justify-content-between"></div>
          <h5 class="mt-0 mb-1">{{position.primary_name}}</h5>
          <h6 class="mt-0 mb-1 pl-3 text-secondary">{{position.secondary_name}}</h6>
          <small>{{position.description}}</small>
        </b-list-group-item>
      </b-list-group>
    </div>
    <!-- <router-view></router-view> -->
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Positions, Position } from "../scheduler_client";

export default Vue.extend({
  components: {},
  data: function() {
    return {
      positions: [] as Position[]
    };
  },
  mounted() {
    let positions: { [key: string]: Position } = this.$store.state.positions;
    let newPositionList: Position[] = [];
    for (const [key, value] of Object.entries(positions)) {
      newPositionList.push(value);
    }
    this.positions = newPositionList;
  },
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "setPositions") {
        let positions: { [key: string]: Position } = this.$store.state
          .positions;
        let newPositionList: Position[] = [];
        for (const [key, value] of Object.entries(positions)) {
          newPositionList.push(value);
        }
        this.positions = newPositionList;
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

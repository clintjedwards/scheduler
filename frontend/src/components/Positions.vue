<template>
  <div class="mt-5 mr-5 ml-5">
    <h5 class="text-secondary">Configure schedule positions</h5>
    <div class="float-right pr-5">
      <button type="button" class="btn btn-dark">Add Position</button>
    </div>
    <div class="content">
      <ul class="list-unstyled">
        <b-media v-for="position in positions" :key="position.id" tag="li">
          <template v-slot:aside>
            <b-img blank blank-color="#abc" width="64" alt="placeholder"></b-img>
          </template>
          <h5 class="mt-0 mb-1">{{position.primary_name}}</h5>
          <h6 class="mt-0 mb-1 pl-3 text-secondary">{{position.secondary_name}}</h6>
          <p class="mb-0">{{position.description}}</p>
          <br />
        </b-media>
      </ul>
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
    this.positions = this.$store.state.positions;
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

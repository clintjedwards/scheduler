<template>
  <div class="mt-5 mr-5 ml-5">
    <h5 class="text-secondary">Configure schedule positions</h5>
    <div class="float-right pr-5">
      <button type="button" class="btn btn-dark">Add Position</button>
    </div>
    <div class="content">
      <ul class="list-unstyled">
        <b-media v-for="position in positionList" :key="position.id" tag="li">
          <template v-slot:aside>
            <b-img blank blank-color="#abc" width="64" alt="placeholder"></b-img>
          </template>
          <h5 class="mt-0 mb-1">{{position.primaryName}}</h5>
          <h6 class="mt-0 mb-1 pl-3 text-secondary">{{position.secondaryName}}</h6>
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
import { Position } from "../scheduler_message_pb";

export default Vue.extend({
  components: {},
  data: function() {
    return {
      positionList: [] as Position.AsObject[]
    };
  },
  mounted() {
    this.positionMapToList();
  },
  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "setPositions") {
        this.positionMapToList();
      }
    });
  },
  methods: {
    positionMapToList: function() {
      let positions: { [key: string]: Position } = this.$store.state.positions;
      let newPositionList: Position.AsObject[] = [];

      for (const [key, value] of Object.entries(positions)) {
        newPositionList.push(value.toObject());
      }
      this.positionList = newPositionList;
    }
  }
});
</script>

<style scoped>
.content {
  margin-top: 5em;
}
</style>

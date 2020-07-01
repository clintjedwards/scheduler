<template>
  <b-modal id="manage_schedule_modal" v-model="showModal">
    <div class="modal-content">
      <h1>This is an example schedule</h1>
    </div>
  </b-modal>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data: function() {
    return {
      showModal: true
    };
  },
  watch: {
    showModal(val) {
      if (!val) {
        this.$router.push({ name: "schedules" });
      }
    }
  }
});
</script>
<style scoped>
</style>

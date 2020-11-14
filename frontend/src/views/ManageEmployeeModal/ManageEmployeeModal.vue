<template>
  <div>
    <form ref="manage_employee_form" @submit.stop.prevent="handleSubmit">
      <b-modal
        id="manage_employee_modal"
        size="xl"
        :title="employee.id"
        title-class="text-muted"
        v-model="showModal"
        @ok="handleSubmit"
        @close="handleClose"
        @cancel="handleClose"
        no-close-on-backdrop
        no-close-on-esc
        ok-title="Employee"
      >
        <p class="text-right text-muted">
          Started {{ humanizeDate }} |
          <b-badge
            variant="success"
            v-if="employee.status === 'active'"
            class="center-badge"
            >Active</b-badge
          >
          <b-badge v-else variant="danger" class="center-badge"
            >Inactive</b-badge
          >
        </p>
        <h1 class="display-3 text-center text-capitalize font-weight-lighter">
          {{ employee.name }}
        </h1>
        <hr />

        <b-container id="employee_details">
          <b-row>
            <b-col>
              <!-- Positions -->
              <h3 class="font-weight-light">Positions</h3>
              <b-list-group
                id="positions_group"
                v-for="(position, index) in positions"
                v-bind:key="`position-${index}`"
                flush
              >
                <b-list-group-item>
                  <p>
                    {{ position.primary_name }} |
                    <small class="text-muted">
                      {{ position.secondary_name }}</small
                    >
                  </p>
                </b-list-group-item>
              </b-list-group>
            </b-col>
            <b-col>
              <!-- Unavailablities -->
              <h3 class="font-weight-light">Unavailabilities</h3>
              <b-list-group
                v-for="(time, index) in employee.unavailabilities"
                v-bind:key="`time-${index}`"
              >
                <b-list-group-item>
                  {{ time }}
                </b-list-group-item>
              </b-list-group>
            </b-col>
          </b-row>
        </b-container>

        <!-- Notes -->
        <h3 id="employee_notes_header" class="font-weight-light">Notes</h3>
        <p>{{ employee.notes }}</p>
      </b-modal>
    </form>
  </div>
</template>

<script>
import SchedulerClient from "../../scheduler_client";
let client = new SchedulerClient();
import * as moment from "moment";

export default {
  components: {},
  async mounted() {
    let employee = await client.getEmployee(this.$route.params.id);
    if (employee === undefined) {
      return;
    }
    this.employee = employee;
  },
  computed: {
    positions() {
      let positions = [];
      for (const [id, position] of Object.entries(this.employee.positions)) {
        positions.push(this.$store.state.positions[id]);
      }
      return positions;
    },
    humanizeDate() {
      return moment(this.employee.start_date, "YYYY-MM-DD").format(
        "MMMM Do, YYYY"
      );
    },
    validateName() {
      if (this.name.length == 0) {
        return null;
      }
      if (this.name.length >= 4) {
        return true;
      }

      return false;
    },
    invalidFeedback() {
      if (this.name.length > 0) {
        return "Enter at least 4 characters.";
      }
      return "Please enter a name.";
    },
  },
  methods: {
    addUnavailTime: function() {
      this.unavailabilities.push({ time: "", state: null });
    },
    removeUnavailTime: function(index) {
      this.unavailabilities.splice(index, 1);
    },
    validateUnavail(index) {
      const regex = /(((\d+,)+\d+|(\d+(-)\d+)|\d+|\*) ?){6}/g;
      if (regex.test(this.unavailabilities[index].time)) {
        this.unavailabilities[index].state = true;
        return true;
      }
      this.unavailabilities[index].state = false;
      return false;
    },
    toggleHelp: function() {
      if (this.showHelp) {
        this.showHelp = false;
        var element = document.getElementById("help_button");
        element.classList.remove("btn-outline-primary");
        element.classList.add("btn-primary");
        return;
      }
      this.showHelp = true;
      var element = document.getElementById("help_button");
      element.classList.remove("btn-primary");
      element.classList.add("btn-outline-primary");
      return;
    },
    handleClose: function() {
      this.$router.push({ name: "Employees" });
    },
    handleSubmit: function(e) {
      e.preventDefault();
      let unavailabilities = [];
      for (const [index, value] of this.unavailabilities.entries()) {
        let valid = this.validateUnavail(index);
        if (!valid) {
          return;
        }
        unavailabilities.push(value.time);
      }

      let payload = {
        name: this.name,
        notes: this.notes,
        start_date: this.start_date,
        unavailabilities: unavailabilities,
        positions: this.selected_positions,
      };

      client
        .createEmployee(payload)
        .then(() => {
          client
            .listEmployees()
            .then((employees) => {
              this.$store.commit("updateEmployees", employees);
              this.$router.push({ name: "Employees" });
            })
            .catch(() => {
              console.log("could not load employees");
              // report error to user here and don't exit modal
            });
        })
        .catch(() => {
          console.log("could not create employee");
        });
    },
  },
  data() {
    return {
      formMode: "",
      showModal: true,
      showHelp: false,
      employee: {
        name: "",
        notes: "",
        start_date: "",
        unavailabilities: [],
        positions: [],
      },
      // unavailabilties is a list of unavail objects
      // an unavail object is the text value of the unavailability and its current form state
      // represented like such {time: string, state: boolean}
      // We store the states in the unavailabilities list because then when we update the state
      // it causes vue to re-render and show the error properly
      unavailabilities: [],
    };
  },
};
</script>

<style scoped>
.center-badge {
  vertical-align: middle;
  margin-top: -0.1em;
}

#employee_details {
  padding-top: 100px;
}

.list-group-item > p {
  margin-bottom: 0px;
}
.list-group-item {
  padding-top: 8px;
  padding-bottom: 8px;
  border-bottom: 1px solid #808080;
}
#employee_notes_header {
  padding-top: 30px;
}
</style>

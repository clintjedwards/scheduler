<template>
  <div>
    <form ref="add_employee_modal" @submit.stop.prevent="handleSubmit">
      <b-modal
        id="add_employee_modal"
        size="xl"
        title="Add Employee"
        v-model="showModal"
        @ok="handleSubmit"
        @close="handleClose"
        @cancel="handleClose"
        no-close-on-backdrop
        no-close-on-esc
        ok-title="Add Employee"
      >
        <div>
          <!--Full name-->
          <b-form-group
            id="fieldset_name"
            label="Full Name"
            label-for="input_name"
            label-cols-sm="4"
            label-cols-lg="3"
            valid-feedback="Valid!"
            :invalid-feedback="invalidFeedback"
          >
            <b-form-input
              id="input-name"
              v-model="name"
              :state="validateName"
              debounce="500"
              trim
            ></b-form-input>
          </b-form-group>

          <!--Start Date-->
          <b-form-group
            id="fieldset_start_date"
            label="Start Date"
            label-for="input_start_date"
            label-cols-sm="4"
            label-cols-lg="3"
          >
            <b-form-datepicker
              id="input_start_date"
              v-model="start_date"
            ></b-form-datepicker>
          </b-form-group>

          <!--Positions-->
          <b-form-group
            id="fieldset_positions"
            label="Positions"
            label-for="input_positions"
            label-cols-sm="4"
            label-cols-lg="3"
            description="Choose positions that this employee is eligible for."
          >
            <b-form-select
              id="input_positions"
              v-model="selected_positions"
              :options="positions"
              multiple
            ></b-form-select>
          </b-form-group>

          <!--Notes-->
          <b-form-group
            id="fieldset_notes"
            label="Notes"
            label-for="input_notes"
            label-cols-sm="4"
            label-cols-lg="3"
          >
            <b-form-textarea
              v-model="notes"
              id="input_notes"
              size="sm"
              trim
            ></b-form-textarea>
          </b-form-group>

          <!--Unavailable Times-->
          <div>
            <b-container>
              <b-row>
                <b-col>
                  <label for="input_unavailabilities">Unavailable Times</label>
                  <b-button
                    size="sm"
                    variant="primary"
                    v-on:click="addUnavailTime"
                  >
                    <b-icon icon="calendar-plus"></b-icon>
                    Add Unavailable Period
                  </b-button>
                  <p class="text-muted">
                    Set the time ranges when this employee is not available
                  </p>
                  <pre class="text-muted">
Field          Allowed values
-----          --------------
Minute         0-59
Hour           0-23
Day of month   1-31
Month          1-12
Day of week    0-7
Year           1970-2100</pre
                  >
                  <div id="dates">
                    <b-input-group
                      class="input_unavailabilities"
                      v-for="(time, index) in unavailabilities"
                      v-bind:key="`time-${index}`"
                    >
                      <b-form-input
                        placeholder="Minute Hour DoM Month DoW Year"
                        v-model="unavailabilities[index].time"
                        :state="unavailabilities[index].state"
                        trim
                      >
                      </b-form-input>
                      <b-input-group-append>
                        <b-button
                          v-on:click="removeUnavailTime(index)"
                          variant="outline-danger"
                        >
                          <b-icon icon="trash"></b-icon>
                        </b-button>
                      </b-input-group-append>
                      <b-form-invalid-feedback>
                        Incorrect format; Click "show help" for formatting help.
                      </b-form-invalid-feedback>
                    </b-input-group>
                  </div>
                </b-col>
                <b-col>
                  <div id="help_button_wrap">
                    <b-button
                      id="help_button"
                      size="sm"
                      v-on:click="toggleHelp()"
                      variant="primary"
                    >
                      <b-icon icon="question-square"></b-icon>
                      Show help
                    </b-button>
                  </div>
                  <div id="help_text" v-if="showHelp">
                    <p>
                      To set the time periods in which an employee is not
                      available, we write the time period in a special format.
                      This format defines the time in which the employee should
                      not work for each individual unit of time. We start from
                      the smallest unit of time and continue to the largest. You
                      can view the units of time in the example table to the
                      left.
                    </p>
                    <p>
                      You are allowed to use commas(representing a list),
                      dashes(representing a range), a star(meaning any value),
                      or just a single number to convey what times an employee
                      is not available.
                    </p>
                    <p>
                      As an example, let's go through how you would write an
                      employee is unavailable from
                      <code>12/25/2020 to 12/30/2020</code>.
                    </p>
                    <ol>
                      <li>
                        The first time period we define is minute. Since for
                        this example we don't want the employee to be
                        unavailable for entire days, we can put a
                        <code>*</code>. This signifies that, for this time
                        period, all minutes are unschedulable.
                      </li>
                      <li>
                        The second is the hour time period. It follows that we
                        also don't want to specify hours. So we put another
                        <code>*</code>.
                      </li>
                      <li>
                        The third is day of month or date. For this we could
                        choose to use a list like so:
                        <code> 25,26,27,28,29,30</code>. But thats a lot of
                        writing. Instead, we use a range: <code>25-30</code>.
                      </li>
                      <li>
                        The third is Day of month. For this we could use a list
                        like so: <code> 25,26,27,28,29,30</code>. But thats a
                        lot of writing. Instead we use a range:
                        <code>25-30</code>.
                      </li>
                      <li>
                        The fourth is Days of the week. For this we don't mind
                        which days the 25-30 fall on. So as usual we leave a
                        <code>*</code>.
                      </li>
                      <li>
                        Lastly, we need to define the year or else this employee
                        gets Christmas off every year. So we just use a single
                        value: <code>2020</code>.
                      </li>
                    </ol>
                    <p>
                      Altogether the full expression becomes:
                      <code>* * 25-30 12 * 2020</code>
                    </p>
                    <pre class="text-muted">
More Examples:

Unavailable Time                Input
-----------------               -----
12/25/2020 to 12/30/2020        * * 25-30 12 * 2020
Never available                 * * * * * *
Every year on the first of Jan  * * 1 1 * *
All of Jun, July, Aug for 2020  * * * 6,7,8 * 2020
Never schedule for 9am          * 9 * * * *
</pre
                    >
                  </div>
                </b-col>
              </b-row>
            </b-container>
          </div>
        </div>
      </b-modal>
    </form>
  </div>
</template>

<script>
import SchedulerClient from "../scheduler_client";
let client = new SchedulerClient();
import * as moment from "moment";

export default {
  components: {},
  computed: {
    validateName() {
      if (this.name.length == 0) {
        return null;
      }
      if (this.name.length >= 4) {
        return true;
      }

      return false;
    },
    positions() {
      let positionsMap = this.$store.state.positions;
      let positions = [];
      for (const [key, value] of Object.entries(positionsMap)) {
        let text = value.primary_name;
        if (value.secondary_name) {
          text = text + " (" + value.secondary_name + ")";
        }

        positions.push({
          value: key,
          text: text,
        });
      }

      return positions;
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
      showModal: true,
      showHelp: false,
      name: "",
      start_date: "",
      notes: "",
      // unavailabilties is a list of unavail objects
      // an unavail object is the text value of the unavailability and its current form state
      // represented like such {time: string, state: boolean}
      // We store the states in the unavailabilities list because then when we update the state
      // it causes vue to re-render and show the error properly
      unavailabilities: [],
      nameState: null,
      selected_positions: [],
    };
  },
};
</script>

<style scoped>
label {
  padding-right: 30px;
}

#dates {
  padding-top: 10px;
  padding-bottom: 10px;
}

#help_text {
  padding-top: 15px;
}

#help_button_wrap {
  padding-top: 5px;
}

.container {
  padding-left: 0px;
}

.input_unavailabilities {
  padding-top: 3px;
  padding-bottom: 3px;
}
</style>

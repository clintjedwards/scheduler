<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import { ProgramsStore } from "../store.js";
  import Button from "./Button.svelte";
  import ManageScheduleForm from "./FormComponents/ManageScheduleForm.svelte";

  let new_schedule = {
    start: "",
    end: "",
    program: ProgramsStore[0],
  };

  let addSchedule = () => {
    new_schedule.program = $ProgramsStore[new_schedule.program];

    client
      .addSchedule(new_schedule)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/schedules", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div>
  <h1 class="font-heading text-4xl text-orange mb-10">Add New Schedule</h1>

  <ManageScheduleForm {new_schedule}>
    <div id="submit" class="full" on:click={addSchedule}>
      <Button>Add Schedule</Button>
    </div>
  </ManageScheduleForm>
</div>

<style>
  #submit {
    margin-top: 20px;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

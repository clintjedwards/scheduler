<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ViewScheduleForm from "./FormComponents/ViewScheduleForm.svelte";

  export let id;

  let schedule;

  client
    .getSchedule(id)
    .then((response) => {
      schedule = response;
    })
    .catch((error) => {
      console.error(error);
    });

  let deleteSchedule = () => {
    client
      .deleteSchedule(id)
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

  let saveSchedule = () => {
    client
      .updateSchedule(id, schedule)
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

<div id="main">
  <div id="actions" class="flex justify-between mb-10 text-xl">
    <Button color="red" on:click={deleteSchedule}>Delete</Button>
  </div>

  <div id="content">
    <ViewScheduleForm {schedule} />
  </div>
</div>

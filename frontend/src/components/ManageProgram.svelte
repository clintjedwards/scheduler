<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ViewProgramForm from "./FormComponents/ViewProgramForm.svelte";

  export let id;

  let program;

  client
    .getProgram(id)
    .then((response) => {
      program = response;
    })
    .catch((error) => {
      console.error(error);
    });

  let deleteProgram = () => {
    client
      .deleteProgram(id)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/programs", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };

  let saveProgram = () => {
    client
      .updateProgram(id, program)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/programs", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div id="main">
  <div id="actions" class="flex justify-between mb-10 text-xl">
    <Button color="red" on:click={deleteProgram}>Delete</Button>
  </div>

  <div id="content">
    <ViewProgramForm {program} />
  </div>
</div>

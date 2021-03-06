<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ViewPositionForm from "./FormComponents/ViewPositionForm.svelte";

  export let id;

  let position;

  client
    .getPosition(id)
    .then((response) => {
      position = response;
    })
    .catch((error) => {
      console.error(error);
    });

  let deletePosition = () => {
    client
      .deletePosition(id)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/positions", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };

  let savePosition = () => {
    client
      .updatePosition(id, position)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/positions", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div id="main">
  <div id="actions" class="flex justify-between mb-10 text-xl">
    <Button type="danger" on:click={deletePosition}>Delete</Button>
  </div>

  <div id="content">
    <ViewPositionForm {position} />
  </div>
</div>

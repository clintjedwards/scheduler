<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ManagePositionForm from "./FormComponents/ManagePositionForm.svelte";

  let new_position = {
    primary_name: "",
    secondary_name: "",
    description: "",
    metadata: {},
  };

  let addPosition = () => {
    client
      .addPosition(new_position)
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

<div>
  <h1 class="font-heading text-4xl text-orange mb-10">Add New Position</h1>

  <ManagePositionForm {new_position}>
    <div id="submit" class="full mt-10 text-xl" on:click={addPosition}>
      <Button>Add Position</Button>
    </div>
  </ManagePositionForm>
</div>

<style>
  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ManagePositionForm from "./FormComponents/ManagePositionForm.svelte";

  let new_schedule = {
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
  <h1 class="font-heading text-4xl text-orange mb-10">Add New Schedule</h1>

  <ManagePositionForm {new_schedule}>
    <div id="submit" class="full" on:click={addPosition}>
      <Button>Add Position</Button>
    </div>
  </ManagePositionForm>
</div>

<style>
  #submit {
    margin-top: 20px;
  }

  h1 {
    color: #ff3e00;
    font-size: 2.5em;
    text-align: left;
    margin: 0 0 30px 0;
    font-weight: 200;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

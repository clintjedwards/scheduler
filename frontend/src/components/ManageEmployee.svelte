<script>
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ViewEmployeeForm from "./FormComponents/ViewEmployeeForm.svelte";

  export let id;

  let employee;
  let mode = "view";

  client
    .getEmployee(id)
    .then((response) => {
      employee = response;
    })
    .catch((error) => {
      console.error(error);
    });

  function switchViewMode() {
    mode = "view";
  }
  function switchEditMode() {
    mode = "edit";
  }
</script>

<div>
  {#if mode === "view"}
    <ViewEmployeeForm {employee} />
  {:else}
    <div>edit mode</div>
  {/if}

  <Button>Delete</Button>
  <Button>Clear</Button>
  <Button on:click={switchViewMode}>View</Button>
  <Button on:click={switchEditMode}>Edit</Button>
</div>

<style></style>

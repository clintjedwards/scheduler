<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ManageEmployeeForm from "./FormComponents/ManageEmployeeForm.svelte";
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

  let deleteEmployee = () => {
    client
      .deleteEmployee(id)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/employees", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div id="main">
  <div id="actions">
    <div>
      {#if mode === "edit"}
        <Button type="danger" on:click={deleteEmployee}>Delete</Button>
      {/if}
    </div>
    <div>
      {#if mode === "edit"}
        <Button on:click={switchViewMode}>View</Button>
      {/if}
      {#if mode === "view"}
        <Button on:click={switchEditMode}>Edit</Button>
      {/if}
      {#if mode === "edit"}
        <Button>Save</Button>
      {/if}
    </div>
  </div>

  <div id="content">
    {#if mode === "view"}
      <ViewEmployeeForm {employee} />
    {:else}
      <ManageEmployeeForm new_employee={employee} />
    {/if}
  </div>
</div>

<style>
  #actions {
    display: flex;
    justify-content: space-between;
    margin-bottom: 40px;
  }
</style>

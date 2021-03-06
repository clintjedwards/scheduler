<script>
  import { link } from "svelte-routing";
  import { client } from "../client.js";
  import { EmployeesStore } from "../store.js";
  import Button from "./Button.svelte";

  client.listEmployees().then((employees) => {
    EmployeesStore.update(() => {
      return employees;
    });
  });
</script>

<employees>
  <div id="actions" class="text-right">
    <a href="/employees/add" use:link><Button>Add Employee</Button></a>
  </div>

  <ul class="text-left w-3/4 mx-auto">
    {#each Object.entries($EmployeesStore) as [id, employee] (id)}
      <a href="/employees/{id}"
        ><li
          class="font-lg p-4 border-b border-gray-500 cursor-pointer hover:bg-orange hover:bg-opacity-25"
        >
          {employee.name}
        </li></a
      >
    {/each}
  </ul>
</employees>

<style>
</style>

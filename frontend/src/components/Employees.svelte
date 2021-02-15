<script>
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
  <div id="actions">
    <Button>Add Employee</Button>
  </div>

  <ul>
    {#each Object.entries($EmployeesStore) as [id, employee] (id)}
      <li>{employee.name}</li>
    {/each}
  </ul>
</employees>

<style>
  #actions {
    text-align: right;
  }
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    font-size: 1.5em;
    font-weight: 300;
    padding-bottom: 1em;
    padding-top: 1em;
    padding-left: 1em;
    border-bottom: 1px solid #6c757d;
    cursor: pointer;
  }

  li:hover {
    background-color: #dfe6e9;
  }
</style>

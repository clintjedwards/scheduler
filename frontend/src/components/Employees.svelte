<script>
  import { client } from "../client.js";
  import { EmployeesStore } from "../store.js";

  client.listEmployees().then((employees) => {
    EmployeesStore.update(() => {
      return employees;
    });
  });
</script>

<employees>
  <ul>
    {#each Object.entries($EmployeesStore) as [id, employee] (id)}
      <li>{employee.name}</li>
    {/each}
  </ul>
</employees>

<style>
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    margin-top: 0.5em;
    margin-bottom: 0.5em;
  }
</style>

<script>
  import { client } from "../client.js";
  import { PositionsStore } from "../store.js";

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });

  let new_employee = {
    name: "",
    notes: "",
    start_date: "",
    unavailabilities: [],
    positions: [],
  };
</script>

<div>
  <h1>Add New Employee</h1>

  <label for="name">Name:</label>
  <input type="text" id="name" bind:value={new_employee.name} />

  <label for="start_date">Start Date:</label>
  <input type="date" id="start_date" bind:value={new_employee.start_date} />

  <h3>Unavail Times</h3>
  {#each new_employee.unavailabilities as unavailability (unavailability)}
    <input type="text" bind:value={unavailability} />
  {/each}
  <button>Add Another</button>

  <label for="positions">Positions:</label>
  <select id="positions">
    {#each Object.entries($PositionsStore) as [id, position] (id)}
      <option value={id}>
        {position.primary_name}
        {#if position.secondary_name}
          | {position.secondary_name}
        {/if}
      </option>
    {/each}
  </select>

  <label for="notes">Notes:</label>
  <textarea id="notes" bind:value={new_employee.notes} />
</div>

<style>
</style>

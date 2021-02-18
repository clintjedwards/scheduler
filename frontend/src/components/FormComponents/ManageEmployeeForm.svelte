<script>
  import { client } from "../../client.js";
  import { PositionsStore } from "../../store.js";
  import UnavailHelp from "./UnavailHelp.svelte";

  export let new_employee = {};

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });

  let showHelp = false;

  let toggleHelp = () => {
    if (showHelp == true) {
      showHelp = false;
    } else {
      showHelp = true;
    }
  };

  let addUnavailInput = () => {
    new_employee.unavailabilities = [...new_employee.unavailabilities, ""];
  };

  let removeUnavailInput = (index) => {
    new_employee.unavailabilities.splice(index, 1);
    new_employee = new_employee;
  };
</script>

<div class="grid-container">
  <div class="full">
    <label for="name">Name:</label>
    <input
      type="text"
      id="name"
      bind:value={new_employee.name}
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="start_date">Start Date:</label>
    <input type="date" id="start_date" bind:value={new_employee.start_date} />
  </div>
  <h3 class="full">Unavailable Times</h3>
  <div>
    <p>Set the time ranges when this employee is not available.</p>
    <pre>
Field          Allowed values
-----          --------------
Minute         0-59
Hour           0-23
Day of month   1-31
Month          1-12
Day of week    0-7
Year           1970-2100
      </pre>
    {#each new_employee.unavailabilities as unavailability, index (index)}
      <div class="time">
        <span class="remove_icon" on:click={() => removeUnavailInput(index)}>
          <img src="/images/trash-alt-regular.svg" alt="remove" />
        </span>
        <input class="unavail_input" type="text" bind:value={unavailability} />
      </div>
    {/each}
    <button id="addButton" on:click={addUnavailInput}>Add Time</button>
  </div>
  <div style="justify-self: center">
    <button id="helpButton" on:click={toggleHelp}>Show Help</button>
    {#if showHelp}
      <UnavailHelp {showHelp} />
    {/if}
  </div>
  <div class="full">
    <label for="positions">Positions:</label>
    <select id="positions" bind:value={new_employee.positions} multiple>
      {#each Object.entries($PositionsStore) as [id, position] (id)}
        <option value={id}>
          {position.primary_name}
          {#if position.secondary_name}
            | {position.secondary_name}
          {/if}
        </option>
      {/each}
    </select>
  </div>
  <div class="full">
    <label for="notes">Notes:</label>
    <textarea id="notes" bind:value={new_employee.notes} />
  </div>
  <slot />
</div>

<style>
  .grid-container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 2em;
    justify-items: start;
  }

  #addButton {
    margin-top: 1em;
  }

  .time {
    display: inline-block;
    position: relative;
  }

  .unavail_input {
    font-size: 1.5em;
    text-align: center;
    margin-top: 2px;
    margin-bottom: 2px;
  }

  .time .remove_icon {
    position: absolute;
    top: 10px;
    left: auto;
    right: 10px;
    cursor: pointer;
  }

  button {
    border: 1px solid #6c757d;
    cursor: pointer;
    padding: 8px 12px;
    font-size: 1em;
    box-shadow: none;
    color: #ff3e00;
    background: white;
  }

  input,
  label {
    display: block;
  }

  label {
    margin-bottom: 2px;
    color: #6c757d;
    text-align: left;
  }

  pre {
    text-align: left;
  }

  h3 {
    margin-top: 0;
    margin-bottom: 0;
    color: #6c757d;
    text-align: left;
    font-weight: 400;
  }

  input#name,
  input#start_date {
    font-weight: 300;
    text-indent: 5%;
  }

  input {
    font-size: 1.5em;
  }

  input#name {
    font-size: 2.5em;
    width: 100%;
  }

  #positions {
    font-size: 1em;
    width: 80%;
    height: 100%;
  }

  #notes {
    width: 80%;
    height: 100%;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

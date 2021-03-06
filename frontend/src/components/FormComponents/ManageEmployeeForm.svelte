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
    <label for="name" class="block text-sm-heading text-xl">Name:</label>
    <input
      type="text"
      id="name"
      class="text-2xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50"
      bind:value={new_employee.name}
      placeholder="Employee Name"
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="start_date" class="block text-sm-heading text-xl"
      >Start Date:
    </label>
    <input
      type="date"
      id="start_date"
      class="text-2xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50"
      placeholder="MM/DD/YYYY"
      bind:value={new_employee.start_date}
    />
  </div>
  <h3 class="full block text-sm-heading text-xl">Unavailable Times</h3>
  <div>
    <p>Set the time ranges when this employee is not available.</p>
    <pre
      class="text-xs mt-4 mb-4">
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
        <input
          type="text"
          class="text-body form-input mt-1 block
          shadow focus:outline-none focus:ring-2 focus:ring-orange
          focus:ring-opacity-50"
          bind:value={unavailability}
          placeholder="* * * * * *"
        />
        <img
          src="/images/trash-icon.svg"
          alt="remove"
          class="remove_icon font-light"
          on:click={() => removeUnavailInput(index)}
        />
      </div>
    {/each}
    <div class="inline-block mr-2 mt-2">
      <button
        type="button"
        id="addButton"
        class="focus:outline-none text-white text-sm py-1.5 px-3
        rounded-md bg-blue-500 hover:bg-blue-600 hover:shadow-lg flex items-center"
        on:click={addUnavailInput}
      >
        <svg
          class="w-4 h-4 mr-2"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6v6m0 0v6m0-6h6m-6 0H6"
          />
        </svg>
        Add Time
      </button>
    </div>
  </div>
  <div style="justify-self: center">
    <button
      type="button"
      id="helpButton"
      class="focus:outline-none text-white text-sm py-1.5 px-3
    rounded-md bg-blue-500 hover:bg-blue-600 hover:shadow-lg flex items-center"
      on:click={toggleHelp}
    >
      <svg
        class="w-4 h-4 mr-2"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
      Show Help
    </button>
    {#if showHelp}
      <UnavailHelp {showHelp} />
    {/if}
  </div>
  <div class="full">
    <label for="positions" class="block text-sm-heading text-xl">
      Positions:
    </label>
    {#if Object.keys($PositionsStore).length > 0}
      <select
        id="positions"
        class="w-full"
        bind:value={new_employee.positions}
        multiple
      >
        {#each Object.entries($PositionsStore) as [id, position] (id)}
          <option value={id}>
            {position.primary_name}
            {#if position.secondary_name}
              | {position.secondary_name}
            {/if}
          </option>
        {/each}
      </select>
    {:else}
      <div>No positions available</div>
    {/if}
  </div>
  <div class="full">
    <label for="notes" class="block text-sm-heading text-xl">Notes:</label>
    <textarea
      id="notes"
      class="text-xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50 w-full"
      placeholder="Employee notes"
      bind:value={new_employee.notes}
    />
  </div>
  <slot />
</div>

<style>
  .grid-container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 2em;
    column-gap: 5em;
    justify-items: start;
  }

  div .time {
    display: flex;
  }

  .time input {
    flex: 1;
    font-size: 1.5em;
    text-align: center;
    margin-top: 2px;
    margin-bottom: 2px;
  }

  .time .remove_icon {
    flex: 0 0 30px;
    height: 30px;
    margin-top: auto;
    margin-bottom: auto;
  }

  textarea {
    resize: vertical;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
    text-align: left;
  }
</style>

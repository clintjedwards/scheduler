<script>
  import { client } from "../../client.js";
  import { ProgramsStore } from "../../store.js";

  export let new_schedule = {};

  client.listPrograms().then((programs) => {
    ProgramsStore.update(() => {
      return programs;
    });
  });
</script>

<div>
  <div class="mb-10 flex space-x-10 items-center">
    <label for="start_date" class="block text-sm-heading text-xl">
      Start Date:
    </label>
    <input
      type="date"
      id="start_date"
      class="text-2xl text-body form-input mt-1 block w-full
        shadow focus:outline-none focus:ring-2 focus:ring-orange
        focus:ring-opacity-50"
      placeholder="MM/DD/YYYY"
      bind:value={new_schedule.start}
    />
    <label for="end_date" class="block text-sm-heading text-xl">
      End Date:
    </label>
    <input
      type="date"
      id="end_date"
      class="text-2xl text-body form-input mt-1 block w-full
          shadow focus:outline-none focus:ring-2 focus:ring-orange
          focus:ring-opacity-50"
      placeholder="MM/DD/YYYY"
      bind:value={new_schedule.end}
    />

    <div>
      <label for="program_select" class="text-sm-heading text-xl">
        Program:
      </label>
      {#if Object.keys($ProgramsStore).length > 0}
        <select id="program_select" bind:value={new_schedule.program}>
          {#each Object.entries($ProgramsStore) as [id, program] (id)}
            <option value={id}>
              {program.name}
            </option>
          {/each}
        </select>
      {:else}
        <div>No programs available</div>
      {/if}
    </div>
  </div>
  <slot />
</div>

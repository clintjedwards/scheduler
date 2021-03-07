<script>
  import { client } from "../../client.js";
  import { PositionsStore } from "../../store.js";

  export let new_program = {};

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });

  let addShift = (day) => {
    new_program[day] = [
      ...new_program[day],
      {
        start: "",
        end: "",
        program_id: "",
      },
    ];
  };

  let removeShift = (day) => {
    new_program[day].pop();
    new_program[day] = new_program[day];
  };
</script>

<div>
  <div class="mb-10">
    <label for="name" class="block text-sm-heading text-xl"> Name:</label>
    <input
      type="text"
      id="name"
      class="text-2xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50"
      bind:value={new_program.name}
      placeholder="Program Name"
      autocomplete="off"
    />
  </div>
  <div>
    {#each Object.entries(new_program) as [key, value] (key)}
      <div class="mb-5">
        {#if ["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"].indexOf(key) >= 0}
          <h2 class="capitalize text-xl text-sm-heading">{key}</h2>
          {#each value as shift, index (index)}
            <div class="flex justify-evenly items-center mb-4">
              <div>
                <h3 class="text-xl">Shift {index + 1}</h3>
              </div>
              <div>
                <label for="shift-start-{index}">Start:</label>
                <input
                  type="time"
                  id="shift-start-{index}"
                  class="text-2xl text-body form-input mt-1
                shadow focus:outline-none focus:ring-2 focus:ring-orange
                focus:ring-opacity-50"
                  bind:value={shift.start}
                  autocomplete="off"
                />
              </div>
              <div>
                <label for="shift-end">End:</label>
                <input
                  type="time"
                  id="shift-end"
                  class="text-2xl text-body form-input mt-1
                shadow focus:outline-none focus:ring-2 focus:ring-orange
                focus:ring-opacity-50"
                  bind:value={shift.end}
                  autocomplete="off"
                />
              </div>
              <div>
                <label for="position" class="text-sm-heading text-xl">
                  Position:
                </label>
                {#if Object.keys($PositionsStore).length > 0}
                  <select id="position" bind:value={shift.position_id}>
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
            </div>
          {/each}
          <div class="flex">
            <button
              type="button"
              id="addShiftButton"
              class="focus:outline-none text-white text-sm py-1.5 px-3
            rounded-md bg-blue-500 hover:bg-blue-600 hover:shadow-lg
            flex items-center mr-4"
              on:click={() => addShift(key)}
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
              Add Shift
            </button>
            {#if value.length !== 0}
              <button
                type="button"
                id="removeShiftButton"
                class="focus:outline-none text-white text-sm py-1.5 px-3
          rounded-md bg-blue-500 hover:bg-blue-600 hover:shadow-lg
          flex items-center"
                on:click={() => removeShift(key)}
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
                Remove Shift
              </button>
            {/if}
          </div>
        {/if}
      </div>
    {/each}
  </div>
  <slot />
</div>

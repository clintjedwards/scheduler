<script>
  import { client } from "../../client.js";
  import { PositionsStore } from "../../store.js";

  export let program = {};

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });

  let humanizeTime = (time) => {};
</script>

<div>
  {#if program.id}
    {#each Object.entries(program) as [key, value] (key)}
      {#if ["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"].indexOf(key) >= 0}
        <div class="mb-5">
          <h2 class="capitalize text-sm-heading text-2xl">{key}</h2>
          {#each value as shift, index (index)}
            <div
              class="flex ml-10 space-x-10 hover:bg-orange hover:bg-opacity-25"
            >
              <div>Shift {index + 1}</div>
              <div>{shift.start} to {shift.end}</div>
              {#if $PositionsStore[shift.position_id]}
                <div>
                  {$PositionsStore[shift.position_id].primary_name}
                  <span class="text-gray-600">
                    {$PositionsStore[shift.position_id].secondary_name}
                  </span>
                </div>
              {/if}
            </div>
          {/each}
        </div>
      {/if}
    {/each}
  {/if}
</div>

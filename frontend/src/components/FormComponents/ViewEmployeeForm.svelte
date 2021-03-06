<script>
  import { format, parse } from "date-fns";
  import { client } from "../../client.js";
  import { PositionsStore } from "../../store.js";

  function buildTimeDate(datetime) {
    let date = format(parse(datetime, "yyyy-MM-dd", new Date()), "MMMM do, y");
    return date;
  }

  export let employee = {};

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });
</script>

<div class="grid-container text-body">
  {#if employee.id}
    <div class="full-column">
      <h1 class="font-heading text-4xl text-orange mb-4">{employee.name}</h1>
      {#if employee.start_date}
        <span class="text-gray-600">
          Started {buildTimeDate(employee.start_date)}
        </span>
      {/if}
    </div>
    <div>
      <h3 class="text-sm-heading text-xl mb-4">Unavailable</h3>
      <ul class="ml-10">
        {#if employee.unavailabilities}
          {#each employee.unavailabilities as time (time)}
            <li>
              {time}
            </li>
          {/each}
        {/if}
      </ul>
    </div>
    <div>
      <h3 class="text-sm-heading text-xl mb-4">Positions</h3>
      <ul class="ml-10">
        {#each employee.positions as id (id)}
          {#if $PositionsStore[id]}
            <li class="list-disc">
              {$PositionsStore[id].primary_name} | {$PositionsStore[id]
                .secondary_name}
            </li>
          {/if}
        {/each}
      </ul>
    </div>
    <div id="notes" class="full-column">
      <p>{employee.notes}</p>
    </div>
  {/if}
</div>

<style>
  .grid-container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 2em;
    justify-items: start;
  }

  .full-column {
    grid-column: 1/3;
    width: 100%;
  }
</style>

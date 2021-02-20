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

<div class="grid-container">
  {#if employee.id}
    <div class="full-column">
      <h1>{employee.name}</h1>
      {#if employee.start_date}
        <span>Started {buildTimeDate(employee.start_date)}</span>
      {/if}
    </div>
    <div>
      <h3>Unavailable</h3>
      <ul>
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
      <h3>Positions</h3>
      <ul>
        {#each Object.entries(employee.positions) as [id, _] (id)}
          {#if $PositionsStore[id]}
            <li>
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

  h1 {
    color: #ff3e00;
    font-size: 2.5em;
    margin: 0 0 30px 0;
    font-weight: 200;
  }

  #notes {
    text-align: left;
  }

  h3 {
    margin-top: 0;
    margin-bottom: 0;
    color: #6c757d;
    text-align: left;
    font-weight: 400;
  }
</style>

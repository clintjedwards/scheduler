<script>
  import { format, parse } from "date-fns";
  import { client } from "../../client.js";
  import { EmployeesStore, PositionsStore } from "../../store.js";

  export let schedule = {};

  function humanizeDate(datetime) {
    let date = format(parse(datetime, "yyyy-MM-dd", new Date()), "MMMM do, y");
    return date;
  }

  let loadEmployee = client.listEmployees().then((employees) => {
    EmployeesStore.update(() => {
      return employees;
    });
  });

  let loadPosition = client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });
</script>

<div>
  {#await loadPosition then _}
    {#if schedule.id}
      <h1 class="font-heading text-4xl text-orange mb-4">
        {humanizeDate(schedule.start)} - {humanizeDate(schedule.end)}
      </h1>
      {#each Object.entries(schedule.time_table) as [date, shifts] (date)}
        <h2>{humanizeDate(date)}</h2>
        {#each shifts as shift (shift)}
          <div>
            {shift.start} - {shift.end}
            Position: {$PositionsStore[shift.position_id].primary_name}
            Employee: {$EmployeesStore[shift.employee_id].name}
          </div>
        {/each}
      {/each}
    {/if}
  {/await}
</div>

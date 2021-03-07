<script>
  export let program = {};
</script>

<div>
  {#if program.id}
    {#each Object.entries(program) as [key, value] (key)}
      {#if ["monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"].indexOf(key) >= 0}
        <h2>{key}</h2>
        {#each value as shift, index (index)}
          <h3>{shift.start}</h3>
          <h3>{shift.end}</h3>
          <h3>{shift.position_id}</h3>
        {/each}
      {/if}
    {/each}
  {/if}
</div>

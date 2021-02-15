<script>
  import { client } from "../client.js";
  import { SchedulesStore } from "../store.js";
  import Button from "./Button.svelte";

  client.listSchedules().then((Schedules) => {
    SchedulesStore.update(() => {
      return Schedules;
    });
  });
</script>

<schedules>
  <div id="actions">
    <Button>Add Schedule</Button>
  </div>

  <ul>
    {#if $SchedulesStore.order}
      {#each $SchedulesStore.order as id (id)}
        <li>
          {$SchedulesStore.schedules[id].start} - {$SchedulesStore.schedules[id]
            .end}
        </li>
      {/each}
    {/if}
  </ul>
</schedules>

<style>
  #actions {
    text-align: right;
  }
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    margin-top: 0.5em;
    margin-bottom: 0.5em;
  }
</style>

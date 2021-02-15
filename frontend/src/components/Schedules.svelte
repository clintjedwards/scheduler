<script>
  import { client } from "../client.js";
  import { SchedulesStore } from "../store.js";

  client.listSchedules().then((Schedules) => {
    SchedulesStore.update(() => {
      console.log(Schedules);
      return Schedules;
    });
  });
</script>

<schedules>
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
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    margin-top: 0.5em;
    margin-bottom: 0.5em;
  }
</style>

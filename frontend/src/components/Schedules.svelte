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
    font-size: 1.5em;
    font-weight: 300;
    padding-bottom: 1em;
    padding-top: 1em;
    padding-left: 1em;
    border-bottom: 1px solid #6c757d;
    cursor: pointer;
  }

  /* ul a {
    color: inherit;
    text-decoration: none;
  } */

  li:hover {
    background-color: #dfe6e9;
  }
</style>

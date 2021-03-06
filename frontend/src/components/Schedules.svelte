<script>
  import { link } from "svelte-routing";
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
  <div id="actions" class="flex justify-end text-xl mr-10">
    <a href="/schedules/add" use:link><Button>Add Schedule</Button></a>
  </div>
  <ul class="text-left w-3/4 mx-auto">
    {#if $SchedulesStore.order}
      {#each $SchedulesStore.order as id (id)}
        <a href="/schedules/{id}">
          <li
            class="font-lg p-4 border-b border-gray-500 cursor-pointer
            hover:bg-orange hover:bg-opacity-25"
          >
            {$SchedulesStore.schedules[id].start} - {$SchedulesStore.schedules[
              id
            ].end}
          </li>
        </a>
      {/each}
    {/if}
  </ul>
</schedules>

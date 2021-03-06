<script>
  import { link } from "svelte-routing";
  import { client } from "../client.js";
  import { PositionsStore } from "../store.js";
  import Button from "./Button.svelte";

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });
</script>

<positions>
  <div id="actions" class="text-right">
    <a href="/positions/add" use:link><Button>Add Position</Button></a>
  </div>

  <ul class="text-left w-3/4 mx-auto">
    {#each Object.entries($PositionsStore) as [id, position] (id)}
      <a href="/positions/{id}">
        <li
          class="font-lg p-4 border-b border-gray-500 cursor-pointer
          hover:bg-orange hover:bg-opacity-25"
        >
          <div>
            <span id="primary" class="text-xl">{position.primary_name}</span>
            {#if position.secondary_name}
              <span id="secondary" class="text-gray-700"
                >{position.secondary_name}</span
              >
            {/if}
          </div>
          <p />
        </li>
      </a>
    {/each}
  </ul>
</positions>

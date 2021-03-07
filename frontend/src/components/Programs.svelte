<script>
  import { link } from "svelte-routing";
  import { client } from "../client.js";
  import { ProgramsStore } from "../store.js";
  import Button from "./Button.svelte";

  client.listPrograms().then((programs) => {
    ProgramsStore.update(() => {
      return programs;
    });
  });
</script>

<programs>
  <div id="actions" class="flex justify-end text-xl mr-10">
    <a href="/programs/add" use:link><Button>Add Program</Button></a>
  </div>

  <ul class="text-left w-3/4 mx-auto">
    {#each Object.entries($ProgramsStore) as [id, program] (id)}
      <a href="/programs/{id}">
        <li
          class="font-lg p-4 border-b border-gray-500 cursor-pointer
          hover:bg-orange hover:bg-opacity-25"
        >
          {program.name}
        </li>
      </a>
    {/each}
  </ul>
</programs>

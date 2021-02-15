<script>
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
  <div id="actions">
    <Button>Add Position</Button>
  </div>

  <ul>
    {#each Object.entries($PositionsStore) as [id, position] (id)}
      <li>
        <div>
          <span id="primary">{position.primary_name}</span>
          {#if position.secondary_name}
            <span id="secondary">{position.secondary_name}</span>
          {/if}
        </div>
        <p />
      </li>
    {/each}
  </ul>
</positions>

<style>
  #actions {
    text-align: right;
  }
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    font-size: 1em;
    font-weight: 300;
    padding-bottom: 1em;
    padding-top: 1em;
    padding-left: 1em;
    border-bottom: 1px solid #6c757d;
    cursor: pointer;
  }

  li:hover {
    background-color: #dfe6e9;
  }

  #primary {
    min-width: 3em;
    font-size: 1.5em;
    display: inline-block;
  }

  #secondary {
    color: #4b5258;
    /* margin-left: 1em; */
  }
</style>

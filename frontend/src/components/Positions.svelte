<script>
  import { client } from "../client.js";
  import { PositionsStore } from "../store.js";

  client.listPositions().then((positions) => {
    PositionsStore.update(() => {
      return positions;
    });
  });
</script>

<positions>
  <ul>
    {#each Object.entries($PositionsStore) as [id, position] (id)}
      <li>
        {position.primary_name} <span>{position.secondary_name}</span>
      </li>
    {/each}
  </ul>
</positions>

<style>
  ul {
    list-style-type: none;
    text-align: left;
  }
  li {
    margin-top: 0.5em;
    margin-bottom: 0.5em;
  }
  span {
    color: #6c757d;
    margin-left: 1em;
  }
</style>

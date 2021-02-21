<script>
  export let position = {};
</script>

<div class="grid-container">
  {#if position.id}
    <div class="full-column">
      <h1>{position.primary_name}</h1>
    </div>
    <div class="full-column">
      <h2>{position.secondary_name}</h2>
    </div>
    <div id="description" class="full-column">
      <p>{position.description}</p>
    </div>
  {/if}
</div>

<style>
  .grid-container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 2em;
    justify-items: start;
  }

  .full-column {
    grid-column: 1/3;
    width: 100%;
  }

  h1 {
    color: #ff3e00;
    font-size: 2.5em;
    margin: 0 0 30px 0;
    font-weight: 200;
  }
</style>

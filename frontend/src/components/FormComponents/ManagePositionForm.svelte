<script>
  export let new_position = {};
</script>

<div class="grid-container">
  <div class="full">
    <label for="primary_name">Primary Name:</label>
    <input
      type="text"
      id="primary_name"
      bind:value={new_position.primary_name}
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="secondary_name">Secondary Name:</label>
    <input
      type="text"
      id="secondary_name"
      bind:value={new_position.secondary_name}
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="description">Description:</label>
    <textarea id="description" bind:value={new_position.description} />
  </div>
  <slot />
</div>

<style>
  .grid-container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    row-gap: 2em;
    column-gap: 5em;
    justify-items: start;
  }

  label {
    display: block;
    margin-bottom: 2px;
    color: #6c757d;
  }

  #secondary_name,
  #primary_name {
    font-weight: 300;
    text-indent: 5%;
    font-size: 2.5em;
    width: 100%;
  }

  input {
    font-size: 1.5em;
  }

  #description {
    width: 80%;
    height: 100%;
  }

  textarea {
    resize: vertical;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
    text-align: left;
  }
</style>

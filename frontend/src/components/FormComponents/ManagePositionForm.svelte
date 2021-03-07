<script>
  export let new_position = {};
</script>

<div class="grid-container">
  <div class="full">
    <label for="primary_name" class="block text-sm-heading text-xl">
      Primary Name:
    </label>
    <input
      type="text"
      id="primary_name"
      class="text-2xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50"
      bind:value={new_position.primary_name}
      placeholder="Retail"
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="secondary_name" class="block text-sm-heading text-xl">
      Secondary Name:
    </label>
    <input
      type="text"
      id="secondary_name"
      class="text-2xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50"
      bind:value={new_position.secondary_name}
      placeholder="Cookie Baking"
      autocomplete="off"
    />
  </div>
  <div class="full">
    <label for="description" class="block text-sm-heading text-xl">
      Description:
    </label>
    <textarea
      id="description"
      class="text-xl text-body form-input mt-1 block w-full
      shadow focus:outline-none focus:ring-2 focus:ring-orange
      focus:ring-opacity-50 w-full"
      placeholder="Position Description"
      bind:value={new_position.description}
    />
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

  .full {
    grid-column: 1/3;
    width: 100%;
    text-align: left;
  }
</style>

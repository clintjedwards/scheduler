<script>
  export let position = {};
</script>

<div>
  {#if position.id}
    <h1 class="font-heading text-4xl text-orange mb-2">
      {position.primary_name}
    </h1>
    <h2 class="font-sm-heading text-3xl mb-4 ml-6">
      {position.secondary_name}
    </h2>
    <p class="text-body">{position.description}</p>
  {/if}
</div>

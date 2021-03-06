<script>
  export let color = "blue";
</script>

<button
  type="button"
  class="focus:outline-none text-{color}-600 py-1.5 px-3
  rounded-md hover:bg-{color}-600 hover:shadow-lg hover:bg-opacity-25
  flex items-center"
  on:click
>
  <slot />
</button>

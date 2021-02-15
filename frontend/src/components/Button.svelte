<script>
  export let type = "primary";
</script>

<button class={type} on:click>
  <slot />
</button>

<style>
  button {
    border: 0;
    cursor: pointer;
    padding: 8px 12px;
    font-size: 1.2em;
    box-shadow: none;
  }
  .primary {
    color: #2574a9;
    background: white;
  }
  .success {
    background: #45c496;
    color: white;
  }
  .danger {
    color: #d91b42;
    background: white;
  }
</style>

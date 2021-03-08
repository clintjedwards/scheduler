<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ManageProgramForm from "./FormComponents/ManageProgramForm.svelte";

  let new_program = {
    name: "",
    monday: [],
    tuesday: [],
    wednesday: [],
    thursday: [],
    friday: [],
    saturday: [],
    sunday: [],
  };

  let normalizeTime = (time) => {
    return time.replace(":", "");
  };

  let addProgram = () => {
    for (const [key, shifts] of Object.entries(new_program)) {
      if (key === "name") {
        continue;
      }

      shifts.forEach((shift) => {
        shift.start = normalizeTime(shift.start);
        shift.end = normalizeTime(shift.end);
      });

      new_program[key] = shifts;
    }

    client
      .addProgram(new_program)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/programs", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div>
  <h1 class="font-heading text-4xl text-orange mb-10">Add New Program</h1>

  <ManageProgramForm {new_program}>
    <div id="submit" class="full mt-10 text-xl" on:click={addProgram}>
      <Button>Add Program</Button>
    </div>
  </ManageProgramForm>
</div>

<style>
  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

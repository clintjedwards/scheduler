<script>
  import { navigate } from "svelte-routing";
  import { client } from "../client.js";
  import Button from "./Button.svelte";
  import ManageEmployeeForm from "./FormComponents/ManageEmployeeForm.svelte";

  let new_employee = {
    name: "",
    notes: "",
    start_date: "",
    unavailabilities: [""],
    positions: [],
  };

  let addEmployee = () => {
    new_employee.unavailabilities.forEach((time, index) => {
      time = time.trim();

      if (time === "") {
        new_employee.unavailabilities.splice(index, 1);
        return;
      }

      new_employee.unavailabilities[index] = time;
    });

    client
      .addEmployee(new_employee)
      .then((response) => {
        if (!response.ok) {
          throw new Error(
            `incorrect error code returned: ${response.status} ${response.statusText}`
          );
        }
      })
      .then(() => {
        navigate("/employees", { replace: true });
      })
      .catch((error) => {
        console.error(error);
      });
  };
</script>

<div class="w-3/4 mx-auto">
  <h1 class="font-heading text-4xl text-orange mb-10">Add New Employee</h1>

  <ManageEmployeeForm {new_employee}>
    <div id="submit" class="full mt-10" on:click={addEmployee}>
      <Button>Add Employee</Button>
    </div>
  </ManageEmployeeForm>
</div>

<style>
  #submit {
    margin-top: 20px;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

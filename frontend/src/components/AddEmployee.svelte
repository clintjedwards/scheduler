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

<div>
  <h1>Add New Employee</h1>

  <ManageEmployeeForm {new_employee}>
    <div id="submit" class="full" on:click={addEmployee}>
      <Button>Add Employee</Button>
    </div>
  </ManageEmployeeForm>
</div>

<style>
  #submit {
    margin-top: 20px;
  }

  h1 {
    color: #ff3e00;
    font-size: 2.5em;
    text-align: left;
    margin: 0 0 30px 0;
    font-weight: 200;
  }

  .full {
    grid-column: 1/3;
    width: 100%;
  }
</style>

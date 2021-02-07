import SchedulerClient from "./scheduler_client.js";

function populateEmployees() {
  let client = new SchedulerClient();
  client.listEmployees().then((employees) => {
    let employee_list = document.getElementById("employee_list");

    for (const [id, employee] of Object.entries(employees)) {
      let a = document.createElement("a");
      a.setAttribute("href", "/employee.html?id=" + id);

      let li = document.createElement("li");
      let context = document.createTextNode(employee.name);

      a.appendChild(context);

      li.appendChild(a);
      li.setAttribute("class", "mb-3 mt-3");

      employee_list.appendChild(li);
    }
  });
}

document.addEventListener("DOMContentLoaded", function() {
  populateEmployees();
});

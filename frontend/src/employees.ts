var utils = require("./utils");
var client = require("./scheduler_client");

var employee = {
  populateEmployeeList: function () {
    client.listEmployees().then((data) => {
      localStorage.setItem("employees", JSON.stringify(data));
    });
  },

  renderEmployees: function () {
    client.listEmployees().then((data) => {
      let employees = data;
      localStorage.setItem("employees", JSON.stringify(data));

      let content = document.getElementById("employees-content-body");

      let innerHTML = "";
      innerHTML += "<ul class='collection'>";

      for (const [id, employee] of (<any>Object).entries(employees)) {
        innerHTML += `<li class="collection-item">
        <span class="title">${utils.sanitizeHTML(employee.name)}</span>
        </li>`;
      }
      innerHTML += "</ul>";

      content.innerHTML = innerHTML;
    });
  },
};

module.exports = employee;

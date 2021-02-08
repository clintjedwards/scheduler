import SchedulerClient from "./scheduler_client.js";

function populateEmployee() {
  let client = new SchedulerClient();

  const urlParams = new URLSearchParams(window.location.search);
  const id = urlParams.get("id");

  client.getEmployee(id).then((employee) => {
    let body = document.getElementById("employee-info-top");
    body.appendChild(generateID(employee.id));
    body.appendChild(generateName(employee.name));
    body.appendChild(generateStartDate(employee.start_date));
    generatePositions(employee.positions);
    generateUnavailabilities(employee.unavailabilities);
    generateNotes(employee.notes);
  });
}

function generateID(employeeID) {
  let p = document.createElement("p");
  p.setAttribute("class", "text-gray-500");
  let content = document.createTextNode(employeeID);
  p.appendChild(content);
  return p;
}

function generateName(employeeName) {
  let h1 = document.createElement("h1");
  h1.setAttribute("class", "text-5xl font-light mb-2");
  let name = document.createTextNode(employeeName);
  h1.appendChild(name);

  return h1;
}

function generateStartDate(employeeStartDate) {
  let p = document.createElement("p");
  p.setAttribute("class", "text-gray-500");
  let content = document.createTextNode(
    "Started " + humanizeDate(employeeStartDate)
  );
  p.appendChild(content);
  return p;
}

function humanizeDate(date) {
  return moment(date, "YYYY-MM-DD").format("MMMM Do, YYYY");
}

function generatePositions(employeePositions) {
  if (!employeePositions) {
    return;
  }

  let client = new SchedulerClient();
  client.listPositions().then(function(positions) {
    let positions_list_element = document.getElementById("positions-list");

    for (const [id, empty] of Object.entries(employeePositions)) {
      let a = document.createElement("a");
      a.setAttribute("href", "/position.html?id=" + id);

      let li = document.createElement("li");
      //TODO(clintjedwards:sanitize this content
      let text = positions[id].primary_name;

      if (positions[id].secondary_name.length !== 0) {
        text =
          text +
          " | <span class='text-gray-500'>" +
          positions[id].secondary_name +
          "</span>";
      }
      a.innerHTML = text;

      li.appendChild(a);
      li.setAttribute("class", "mb-3 mt-3");

      positions_list_element.appendChild(li);
    }
  });
}

function generateUnavailabilities(employeeUnavailabilities) {
  if (!employeeUnavailabilities) {
    return;
  }

  let unavail_list_element = document.getElementById("unavailabilities-list");

  employeeUnavailabilities.forEach(function(unavail) {
    let li = document.createElement("li");
    let text = document.createTextNode(unavail);
    li.appendChild(text);
    li.setAttribute("class", "mb-3 mt-3");

    unavail_list_element.appendChild(li);
  });
}

function generateNotes(employeeNotes) {
  let notes_element = document.getElementById("employee-info-bottom");
  let text = document.createTextNode(employeeNotes);
  notes_element.appendChild(text);
}

document.addEventListener("DOMContentLoaded", function() {
  populateEmployee();
});

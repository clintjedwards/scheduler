// generateCalendar generates the times and allocations for a specific schedule
function generateCalendar(timetable) {
  let html = `<div id="calendar">`;
  html += generateDays(timetable, false);
  html += `</div>`;
  return html;
}

// generateHeadings generates the dates that are used for headings to the timetable
function generateHeadings(timetable) {
  let html = `<div class="heading">`;

  for (const [date, shifts] of Object.entries(timetable)) {
    // convert date into pretty date and use as heading
    // take in date format: 06-19-1990 and return Jun and 19th
    momentObj = moment(date, "MM-DD-YYYY");
    humanDate = momentObj.format("Do");
    humanDay = momentObj.format("ddd");

    html += `<div class="colgroup"><div class="row"><h2>${humanDay}</h2><h2>${humanDate}</h2></div></div>`;
  }
  html += `</div>`;
  return html;
}

// generateDays iterates through the timetable datastructure by times first and then by date
// this is because to create valid html we need to generate things by rows instead of columns.
// To this end we start iterating through a static list of all times first and then
// draw the row by iterating through allocs for each date
function generateDays(timetable) {
  let html = "";
  let employees = JSON.parse(localStorage.getItem("employees"));
  let positions = JSON.parse(localStorage.getItem("positions"));

  for (const [date, shifts] of Object.entries(timetable)) {
    html += `<div class="colgroup">`;
    for (let shift of timetable[date]) {
      html += `<div class="row">
      <h6>${humanizeTime(shift.start)} - ${humanizeTime(shift.end)}</h6>
      <h5 class="text-tertiary">${sanitizeHTML(
        employees[shift.employee_id].name
      )}</h5>
        <p>${sanitizeHTML(positions[shift.position_id].primary_name)}`;
      if (positions[shift.position_id].secondary_name !== "") {
        html += `<span class="text-secondary"> (${
          positions[shift.position_id].secondary_name
        })</span>`;
      }
      html += `</p></div>`;
    }
    html += `</div>`;
  }

  return html;
}

function humanizeTime(time) {
  momentObj = moment(time, "hhmm");
  humanTime = momentObj.format("h:mm a"); // format: 1:23 pm
  return humanTime;
}

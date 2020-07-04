const timeslots = [
  "0000",
  "0030",
  "0100",
  "0130",
  "0200",
  "0230",
  "0300",
  "0330",
  "0400",
  "0430",
  "0500",
  "0530",
  "0600",
  "0630",
  "0700",
  "0730",
  "0800",
  "0830",
  "0900",
  "0930",
  "1000",
  "1030",
  "1100",
  "1130",
  "1200",
  "1230",
  "1300",
  "1330",
  "1400",
  "1430",
  "1500",
  "1530",
  "1600",
  "1630",
  "1700",
  "1730",
  "1800",
  "1830",
  "1900",
  "1930",
  "2000",
  "2030",
  "2100",
  "2130",
  "2200",
  "2230",
  "2300",
  "2330",
];

// occupiedTimeSlots returns a list of timeslots that have at least one allocation.
// This is useful in a situation where we don't want to render unused time slots.
function occupiedTimeSlots(timetable) {
  let usedTimeSlots = {};

  for (const [date, times] of Object.entries(timetable)) {
    usedTimeSlots = {};
    for (const time of timeslots) {
      if (times[time].length !== 0) {
        usedTimeSlots[time] = true;
      }
    }
  }

  return usedTimeSlots;
}

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
  html += `<div class="cell"></div>`;

  for (const [date, times] of Object.entries(timetable)) {
    // convert date into pretty date and use as heading
    // take in date format: 06-19-1990 and return Jun and 19th
    momentObj = moment(date, "MM-DD-YYYY");
    humanDate = momentObj.format("Do");
    humanDay = momentObj.format("ddd");

    html += `<div class="cell"><h2>${humanDay}</h2><h2>${humanDate}</h2></div>`;
  }
  html += `</div>`;
  return html;
}

// generateDays iterates through the timetable datastructure by times first and then by date
// this is because to create valid html we need to generate things by rows instead of columns.
// To this end we start iterating through a static list of all times first and then
// draw the row by iterating through allocs for each date
function generateDays(timetable, drawAll) {
  let occupiedSlots = occupiedTimeSlots(timetable);

  let html = "";
  for (const time of timeslots) {
    // if this timeslot contains no allocations don't draw it
    if (!drawAll && !occupiedSlots[time]) {
      continue;
    }
    html += `<div class="row">`;

    // render the time period y axis
    momentObj = moment(time, "hhmm");
    humanTime = momentObj.format("h:mm a"); // format: 1:23 pm
    html += `<div class="cell timeslot">${humanTime}</div>`;

    for (const [date, timeslots] of Object.entries(timetable)) {
      // render all allocations for time period
      let allocs = timeslots[time];
      if (allocs.length === 0) {
        html += `<div class="cell"></div>`;
        continue;
      }
      html += `<div class="cell">
            <h6>${state.employees[allocs[0].employee_id].name}</h6>
            <p>${state.positions[allocs[0].position_id].primary_name}</p>
            <p class="text-secondary">${
              state.positions[allocs[0].position_id].secondary_name
            }</p>
        </div>`;
    }
    html += `</div>`;
  }
  return html;
}

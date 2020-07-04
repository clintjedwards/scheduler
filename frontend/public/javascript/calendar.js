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

// returns a list of empty timeslots so that we can not render them if we so choose
function usedTimeSlots(timetable) {
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

function generateCalendar(timetable) {
  let html = `<div id="calendar">`;
  html += generateDays(timetable, false);
  html += `</div>`;
  return html;
}

function generateHeadings(timetable) {
  let html = `<div class="heading">`;
  html += `<div class="cell"></div>`;

  for (const [date, times] of Object.entries(timetable)) {
    momentObj = moment(date, "MM-DD-YYYY");
    humanDate = momentObj.format("Do");
    humanDay = momentObj.format("ddd");

    // convert date into pretty date and use as heading
    html += `<div class="cell"><h2>${humanDay}</h2><h2>${humanDate}</h2></div>`;
  }
  html += `</div>`;
  return html;
}

function generateDays(timetable, drawAll) {
  let html = "";
  let nonEmptySlots = usedTimeSlots(timetable);
  for (const time of timeslots) {
    if (!drawAll && !nonEmptySlots[time]) {
      continue;
    }
    html += `<div class="row">`;
    // render the time period y axis
    momentObj = moment(time, "hhmm");
    // format: 1:23 pm
    humanTime = momentObj.format("h:mm a");
    html += `<div class="cell timeslot">${humanTime}</div>`;
    for (const [date, timeslots] of Object.entries(timetable)) {
      // render all allocations for that time period
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

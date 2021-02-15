class SchedulerClient {
  url = location.protocol + "//" + location.host;

  getSystemInfo() {
    return fetch(`${this.url}/api/system/info`).then((res) => res.json());
  }

  listEmployees() {
    return fetch(`${this.url}/api/employees`).then((res) => res.json());
  }

  getEmployee(id) {
    return fetch(`${this.url}/api/employees/${id}`).then((res) => res.json());
  }

  listPositions() {
    return fetch(`${this.url}/api/positions`).then((res) => res.json());
  }

  getPosition(id) {
    return fetch(`${this.url}/api/positions/${id}`).then((res) => res.json());
  }

  listSchedules() {
    return fetch(`${this.url}/api/schedules`).then((res) => res.json());
  }

  getSchedule(id) {
    return fetch(`${this.url}/api/schedules/${id}`).then((res) => res.json());
  }
}

export let client = new SchedulerClient();

interface Employee {
  id: string;
  name: string;
  notes: string;
  start_date: string;
  status: string;
  unavailable: { [key: string]: string };
  positions: { [key: string]: boolean };
  preferences: { [key: string]: string };
  created: number;
  modified: number;
}

interface Program {
  monday: { [key: string]: Shift };
  tuesday: { [key: string]: Shift };
  wednesday: { [key: string]: Shift };
  thursday: { [key: string]: Shift };
  friday: { [key: string]: Shift };
  saturday: { [key: string]: Shift };
  sunday: { [key: string]: Shift };
}

interface Shift {
  start: string;
  end: string;
  employee: string;
}

interface Schedule {
  id: string;
  start: string;
  end: string;
  program: Program;
  preferences: { [key: string]: string };
  employee_filter: string[];
  time_table: Map<string, Map<string, Shift[]>>;
}

interface SystemInfo {
  build_time: string;
  commit: string;
  debug_enabled: boolean;
  frontend_enabled: boolean;
  semver: string;
}

interface Position {
  id: string;
  primary_name: string;
  secondary_name: string;
  description: string;
}

interface Employees {
  [key: string]: Employee;
}

interface Positions {
  [key: string]: Position;
}

interface Schedules {
  schedule: Map<string, Schedule>;
  order: string[];
}

class SchedulerClient {
  url = location.protocol + "//" + location.host;

  getSystemInfo(): Promise<SystemInfo> {
    return fetch(`${this.url}/api/system/info`).then((res) => res.json());
  }

  listEmployees(): Promise<Employees> {
    return fetch(`${this.url}/api/employees`).then((res) => res.json());
  }

  getEmployee(id: string) {
    return fetch(`${this.url}/api/employees/${id}`).then((res) => res.json());
  }

  addEmployees(employee: Employee) {
    //return fetch();
  }

  listPositions(): Promise<Positions> {
    return fetch(`${this.url}/api/positions`).then((res) => res.json());
  }

  getPosition(id: string) {
    return fetch(`${this.url}/api/positions/${id}`).then((res) => res.json());
  }

  listSchedules(): Promise<Schedules> {
    return fetch(`${this.url}/api/schedules`).then((res) => res.json());
  }

  getSchedule(id: string) {
    return fetch(`${this.url}/api/schedules/${id}`).then((res) => res.json());
  }
}

export {
  Program,
  Schedule,
  Employees,
  Employee,
  Positions,
  SchedulerClient,
  Shift,
  SystemInfo,
  Schedules,
  Position,
};

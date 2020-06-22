export interface SystemInfo {
  build_time: string;
  commit: string;
  debug_enabled: boolean;
  frontend_enabled: boolean;
  semver: string;
}

export interface Employees {
  [k: string]: Employee;
}

export interface Employee {
  id: string;
  name: string;
  notes: string;
  start_date: string;
  status: string;
  unavailable: Map<string, string>;
  positions: Map<string, boolean>;
  preferences: Map<string, string>;
  created: number;
  modified: number;
}

export class SchedulerClient {
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
}

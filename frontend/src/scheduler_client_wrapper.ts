import { Promise } from "es6-promise";
import { Employee, Position } from "./scheduler_message_pb";
import { SchedulerAPIClient } from "./Scheduler_serviceServiceClientPb";
import {
  AddEmployeeRequest,
  AddPositionRequest,
  GetSystemInfoRequest,
  ListEmployeesRequest,
  ListPositionsRequest,
} from "./scheduler_transport_pb";

export interface Employees {
  [key: string]: Employee;
}

export interface Positions {
  [key: string]: Position;
}

export interface SystemInfo {
  build_time: string;
  commit: string;
  debug_enabled: boolean;
  frontend_enabled: boolean;
  semver: string;
}

export class SchedulerClientWrapper {
  client: SchedulerAPIClient;

  constructor() {
    let url = location.protocol + "//" + location.host;
    this.client = new SchedulerAPIClient(url, null, null);
  }

  listEmployees(): Promise<Employees | undefined> {
    let listEmployeesRequest = new ListEmployeesRequest();

    return new Promise((resolve, reject) => {
      this.client.listEmployees(listEmployeesRequest, null, function(
        err,
        response
      ) {
        if (err) {
          reject(err);
        }

        let employees: Employees = {};

        response
          .getEmployeesMap()
          .forEach(function(value: Employee, key: string | number) {
            employees[key] = value;
          });

        resolve(employees);
      });
    });
  }

  addEmployee(employeeData: AddEmployeeRequest.AsObject): Promise<string> {
    return new Promise((resolve, reject) => {
      let addEmployeeRequest = new AddEmployeeRequest();
      addEmployeeRequest.setName(employeeData.name);
      addEmployeeRequest.setNotes(employeeData.notes);
      addEmployeeRequest.setStartDate(employeeData.startDate);

      let positionsMap = addEmployeeRequest.getPositionsMap();
      employeeData.positionsMap.forEach(function(value, key) {
        positionsMap.set(value[0], value[1]);
      });

      let preferencesMap = addEmployeeRequest.getPreferencesMap();
      employeeData.preferencesMap.forEach(function(value, key) {
        preferencesMap.set(value[0], value[1]);
      });

      let unavailabilityMap = addEmployeeRequest.getUnavailabilityMap();
      employeeData.unavailabilityMap.forEach(function(value, key) {
        unavailabilityMap.set(value[0], value[1]);
      });

      this.client.addEmployee(addEmployeeRequest, null, function(
        err,
        response
      ) {
        if (err) {
          reject(err);
        }
        resolve();
      });
    });
  }

  listPositions(): Promise<Positions | undefined> {
    let listPositionsRequest = new ListPositionsRequest();

    return new Promise((resolve, reject) => {
      this.client.listPositions(listPositionsRequest, null, function(
        err,
        response
      ) {
        if (err) {
          reject(err);
        }

        let positions: Positions = {};

        response
          .getPositionsMap()
          .forEach(function(value: Position, key: string | number) {
            positions[key] = value;
          });

        resolve(positions);
      });
    });
  }

  addPosition(positionData: AddPositionRequest.AsObject): Promise<string> {
    return new Promise((resolve, reject) => {
      let addPositionRequest = new AddPositionRequest();
      addPositionRequest.setPrimaryName(positionData.primaryName);
      addPositionRequest.setSecondaryName(positionData.secondaryName);
      addPositionRequest.setDescription(positionData.description);

      this.client.addPosition(addPositionRequest, null, function(
        err,
        response
      ) {
        if (err) {
          reject(err);
        }
        resolve();
      });
    });
  }

  //getSystemInfo retrieves a system information
  getSystemInfo(): Promise<SystemInfo | undefined> {
    let getSystemInfoRequest = new GetSystemInfoRequest();

    return new Promise((resolve, reject) => {
      this.client.getSystemInfo(getSystemInfoRequest, {}, function(
        err,
        response
      ) {
        if (err) {
          reject(err);
          return;
        }
        let systemInfo: SystemInfo = {
          build_time: response.getBuildTime(),
          commit: response.getCommit(),
          debug_enabled: response.getDebugEnabled(),
          frontend_enabled: response.getFrontendEnabled(),
          semver: response.getSemver(),
        };
        resolve(systemInfo);
      });
    });
  }
}

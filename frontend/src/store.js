import { writable } from "svelte/store";

const EmployeesStore = writable({});
const PositionsStore = writable({});
const SchedulesStore = writable({});

export { EmployeesStore, PositionsStore, SchedulesStore };

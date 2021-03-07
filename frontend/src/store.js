import { writable } from "svelte/store";

const EmployeesStore = writable({});
const PositionsStore = writable({});
const ProgramsStore = writable({});
const SchedulesStore = writable({});

export { EmployeesStore, PositionsStore, ProgramsStore, SchedulesStore };

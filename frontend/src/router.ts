import Vue from "vue";
import VueRouter from "vue-router";
import Employees from "./components/Employees.vue";
import NotFound from "./components/NotFound.vue";
import Positions from "./components/Positions.vue";
import Schedules from "./components/Schedules.vue";

Vue.use(VueRouter);

const routes = [
  { path: "/", redirect: "/employees" },
  {
    path: "/employees",
    name: "employees",
    component: Employees,
    // children: [
    //   {
    //     path: "add",
    //     name: "addEmployeeModal",
    //     component: AddEmployeeModal,
    //   },
    //   {
    //     path: ":id",
    //     name: "updateEmployeeModal",
    //     component: UpdateEmployeeModal,
    //   },
    // ],
  },
  {
    path: "/positions",
    name: "positions",
    component: Positions,
    // children: [
    //   {
    //     path: "add",
    //     name: "addEmployeeModal",
    //     component: AddEmployeeModal,
    //   },
    //   {
    //     path: ":id",
    //     name: "updateEmployeeModal",
    //     component: UpdateEmployeeModal,
    //   },
    // ],
  },
  {
    path: "/schedules",
    name: "schedules",
    component: Schedules,
    // children: [
    //   {
    //     path: "add",
    //     name: "addEmployeeModal",
    //     component: AddEmployeeModal,
    //   },
    //   {
    //     path: ":id",
    //     name: "updateEmployeeModal",
    //     component: UpdateEmployeeModal,
    //   },
    // ],
  },
  { path: "*", component: NotFound },
];

const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;

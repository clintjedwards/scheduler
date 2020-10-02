import Vue from "vue";
import VueRouter from "vue-router";
import AddEmployeeModal from "../views/AddEmployeeModal";
import Employees from "../views/Employees";
import Home from "../views/Home";
import Positions from "../views/Positions";
import Schedules from "../views/Schedules";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/employees",
    name: "Employees",
    component: Employees,
    children: [
      {
        path: "add",
        name: "addEmployeeModal",
        component: AddEmployeeModal,
      },
    ],
  },
  {
    path: "/positions",
    name: "Positions",
    component: Positions,
  },
  {
    path: "/schedules",
    name: "Schedules",
    component: Schedules,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;

import Vue from "vue";
import VueRouter from "vue-router";
import Employees from "./components/Employees.vue";
import NotFound from "./components/NotFound.vue";

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
  { path: "*", component: NotFound },
];

const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;

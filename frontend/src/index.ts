import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.css";
import Vue from "vue";
import PageFooter from "./components/PageFooter.vue";
import PageHeader from "./components/PageHeader.vue";
import router from "./router";
import { Employees, SchedulerClient, SystemInfo } from "./scheduler_client";
import store from "./store";

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

let client: SchedulerClient;
client = new SchedulerClient();

router.beforeEach((to, from, next) => {
  if (store.state.isInitialized) {
    next();
    return;
  }

  var employeesPromise = client.listEmployees();
  // var positionsPromise = client.listPositions();

  Promise.all([employeesPromise]).then((values) => {
    store.commit("setEmployees", values[0]);
    // store.commit("setPositions", values[1]);
    store.commit("setIsInitialized");
    next();
    return;
  });
});

const app = new Vue({
  el: "#app",
  store,
  router,
  components: {
    PageFooter,
    PageHeader,
  },
  mounted() {
    client.getSystemInfo().then((systemInfo: SystemInfo | undefined) => {
      store.commit("setSystemInfo", systemInfo);
    });

    setInterval(() => {
      client.listEmployees().then((employees: Employees | undefined) => {
        store.commit("setEmployees", employees);
      });
      // client.listPositions().then((positions: Positions | undefined) => {
      //   store.commit("setPositions", positions);
      // });
    }, 180000); //3mins
  },
});

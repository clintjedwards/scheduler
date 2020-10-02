import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "bootstrap/dist/css/bootstrap.css";
import Vue from "vue";
import Navigation from "./components/Navigation.vue";
import PageFooter from "./components/PageFooter.vue";
import router from "./router";
import SchedulerClient from "./scheduler_client";
import store from "./store";

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
Vue.config.productionTip = false;

let client = new SchedulerClient();

router.beforeEach((to, from, next) => {
  // if store is initialized go straight to next route, if not load the app
  if (store.state.isInitialized) {
    next();
    return;
  }

  var employeesPromise = client.listEmployees();
  var positionsPromise = client.listPositions();
  var schedulesPromise = client.listSchedules();

  Promise.all([employeesPromise, positionsPromise, schedulesPromise]).then(
    (values) => {
      store.commit("updateEmployees", values[0]);
      store.commit("updatePositions", values[1]);
      store.commit("updateSchedules", values[2]);
      store.commit("setIsInitialized");
      next();
      return;
    }
  );
});

new Vue({
  el: "#app",
  router,
  store,
  components: {
    Navigation,
    PageFooter,
  },
  mounted() {
    client.getAppInfo().then((app_info) => {
      if (app_info) {
        store.commit("updateAppInfo", app_info);
      }
    });

    setInterval(() => {
      client.listEmployees().then((employees) => {
        store.commit("updateEmployees", employees);
      });
      client.listPositions().then((positions) => {
        store.commit("updatePositions", positions);
      });
      client.listSchedules().then((schedules) => {
        store.commit("updateSchedules", schedules);
      });
    }, 180000); //3mins
  },
});

import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const state = {
  isInitialized: false,
  app_info: {
    semver: "0.0.0",
    build_time: "0",
    commit: "dev",
    debug_enabled: true,
  },
  employees: {},
  schedules: {
    order: [],
    schedules: {},
  },
  positions: {},
};

const mutations = {
  setIsInitialized(state) {
    state.isInitialized = true;
  },
  updateAppInfo(state, systemInfo) {
    state.app_info = systemInfo;
  },
  updateEmployees(state, employees) {
    state.employees = employees;
  },
  updatePositions(state, positions) {
    state.positions = positions;
  },
  updateSchedules(state, schedules) {
    state.schedules = schedules;
  },
};

const store = new Vuex.Store({
  state,
  mutations,
});

export default store;

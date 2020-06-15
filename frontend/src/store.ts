import Vue from "vue";
import Vuex, { MutationTree } from "vuex";
import { Employees, SystemInfo } from "./scheduler_client_wrapper";

Vue.use(Vuex);

interface RootState {
  isInitialized: boolean;
  appInfo: SystemInfo;
  employees: Employees;
}

const state: RootState = {
  isInitialized: false,
  appInfo: {
    build_time: "",
    commit: "",
    debug_enabled: true,
    frontend_enabled: false,
    semver: "",
  },
  employees: {},
};

const mutations: MutationTree<RootState> = {
  setIsInitialized(state) {
    state.isInitialized = true;
  },
  setAppInfo(state, systemInfo: SystemInfo) {
    state.appInfo = systemInfo;
  },
  setEmployees(state, employees: Employees) {
    state.employees = employees;
  },
};

const store = new Vuex.Store<RootState>({
  state,
  mutations,
});

export default store;

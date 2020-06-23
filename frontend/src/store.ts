import Vue from "vue";
import Vuex, { MutationTree } from "vuex";

Vue.use(Vuex);

interface RootState {
  isInitialized: boolean;
  systemInfo: Object;
  employees: Object;
  positions: Object;
  schedules: Object;
}

const state: RootState = {
  isInitialized: false,
  systemInfo: {},
  employees: {},
  positions: {},
  schedules: {},
};

const mutations: MutationTree<RootState> = {
  setIsInitialized(state) {
    state.isInitialized = true;
  },
  setSystemInfo(state, systemInfo: Object) {
    state.systemInfo = systemInfo;
  },
  setEmployees(state, employees: Object) {
    state.employees = employees;
  },
  setPositions(state, positions: Object) {
    state.positions = positions;
  },
  setSchedules(state, schedules: Object) {
    state.schedules = schedules;
  },
};

const store = new Vuex.Store<RootState>({
  state,
  mutations,
});

export default store;

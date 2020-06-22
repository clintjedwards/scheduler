import Vue from "vue";
import Vuex, { MutationTree } from "vuex";

Vue.use(Vuex);

interface RootState {
  isInitialized: boolean;
  systemInfo: Object;
  employees: Object;
}

const state: RootState = {
  isInitialized: false,
  systemInfo: {},
  employees: {},
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
};

const store = new Vuex.Store<RootState>({
  state,
  mutations,
});

export default store;

import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import actions from './actions'
import mutations from './mutations'
import state from './state'
import {createLogger } from 'vuex' 
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css'
const debug = process.env.NODE_ENV !== 'production';
const Logger = createLogger;

Vue.use(Vuex)
Vue.use(ElementUI)
export default new Vuex.Store({
  getters,
  actions,
  mutations,
  state,
  strict: debug,
  plugins: debug ? [Logger()] : []
})

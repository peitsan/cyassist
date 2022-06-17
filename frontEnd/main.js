import Vue from 'vue'
import App from './App.vue'
import store from '@/store/index.js'
import VueRouter from 'vue-router'
import router from './router/router.js'
Vue.config.productionTip = false
Vue.use(VueRouter);
new Vue({
  el:'#app',
  render: h => h(App),
  router,
  beforeCreate() {
		Vue.prototype.$bus = this
	},
})

import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import { routes } from './routes'

Vue.use(VueRouter)

export const router = new VueRouter({
  mode: 'history',
  routes
})



Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')

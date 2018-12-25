import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
//import VueRouter from 'vue-router'
import routes from './routes'
import BootstrapVue from 'bootstrap-vue'

Vue.use(BootstrapVue);

//Vue.use(VueRouter)

//const router = new VueRouter({routes})
Vue.prototype.$axios = axios

new Vue({
  el: '#app',
  render: h => h(App),
  routes
})

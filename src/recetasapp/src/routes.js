import Vue from 'vue'
import Router from 'vue-router'
import formulario from './formulario'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/new',
      name: 'formulario',
      component: formulario
    }
  ]
})
import Vue from 'vue'
import Router from 'vue-router'
import Replicants from '@/components/Replicants'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/replicants',
      name: 'Replicants',
      component: Replicants
    }
  ]
})

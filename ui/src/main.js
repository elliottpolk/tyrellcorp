// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import icons from './icons'

import vueresource from 'vue-resource'
import BootstrapVue from 'bootstrap-vue'
import uuid from 'vue-uuid'

import '@/assets/css/style.scss'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(icons)
Vue.use(vueresource)
Vue.use(BootstrapVue)
Vue.use(uuid)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})

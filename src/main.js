// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// vue
import Vue from 'vue'
import FastClick from 'fastclick'
// 路由配置表
import VueRouter from 'vue-router'
import App from './App'
import Home from './components/login'
import first from './components/pages/first'
import assembly from './components/pages/assembly'
import VueResource from 'vue-resource'

Vue.use(VueRouter)
Vue.use(VueResource)
// 设置路由
const routes = [{
  path: '/first',
  component: first
},
{
  path: '/',
  component: Home
},
{
  path: '/assembly',
  component: assembly
}]

const router = new VueRouter({
  mode: 'history',
  routes
})

FastClick.attach(document.body)

/* eslint-disable no-new */
new Vue({
  router,
  render: h => h(App)
}).$mount('#app-box')

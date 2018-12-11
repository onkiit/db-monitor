import Vue from 'vue'
import Router from 'vue-router'
import Monitoring from './views/Monitoring.vue'
import NotFound from './views/NotFound.vue'
import { store } from './store/index'

Vue.use(Router)

const router =  new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/monitoring',
      name: 'monitoring',
      component: Monitoring,
      meta: {
        auth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import(/* webpackChunkName: "login" */ './views/Login.vue'),
      meta: {
        auth: false
      }

    },
    {
      path: '/register',
      name: 'register',
      component: () => import(/* webpackChunkName: "register" */ './views/Register.vue'),
      meta: {
        auth: false
      }
    },
    {
      path: '*',
      name: 'notFound',
      component: () => import(/* webpackChunkName: "notFound" */ './views/NotFound.vue'),
    }
  ]
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.auth)){
    if(!store.getters['isLoggedIn']){
      next({
        path: '/login',
        query: {
          redirect: to.fullPath
        }
      })
    }
    else{
      next()
    }
  }
  else{
    next()
  }
})

export default router
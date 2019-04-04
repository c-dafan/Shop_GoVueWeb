import Vue from 'vue'
import Router from 'vue-router'
import home_head from '../view/head/home_head'
import home from '../view/body_main/home'
import cart_home from '../view/body_main/cart_home'
import login_home from '../view/body_main/login_home'
import details from '../view/body_main/details'
import function_self from '../view/Aside/function_self'
import orders from '../view/Aside/orders'
import address from '../view/Aside/address'
import focus_user from '../view/Aside/focus_user'
import info_user from '../view/Aside/info_user'
import history_user from '../view/Aside/history_user'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/home',
      name: 'home',
      components: {
        head: home_head,
        main: home
      },
    },
    {
      path: '/cart',
      name: 'cart',
      components: {
        head: home_head,
        main: cart_home
      },
    },
    {
      path: '/',
      name: 'app',
      components: {
        head: home_head,
        main: login_home,
      },
    },
    {
      path: '/detail/:gid',
      name: 'detail',
      components: {
        head: home_head,
        main: details,
      },
    }, {
      path: '/info',
      name: 'info',
      components: {
        head: home_head,
        main: function_self,
      },
      children: [
        {
          path: 'order/:inf',
          name: 'order',
          component: orders,
        },
        {
          path: 'address',
          name: 'address',
          component: address
        },
        {
          path: 'focus',
          name: 'focus',
          component: focus_user
        },
        {
          path: 'history',
          name: 'history',
          component: history_user
        },
        {
          path: 'info/:inf',
          name: 'info_user',
          component: info_user
        }
      ]
    },
  ]
})
import store from '../stores/store'

router.beforeEach((to, from, next) => {
  if (to.fullPath === '/') {
    next()
  } else {
    if (store.state.userInfo.token === 'undefined' ||
      store.state.userInfo.token === '' ||
      store.state.userInfo.token === null) {
      next('/')
    } else {
      next()
    }
  }
})

export default router

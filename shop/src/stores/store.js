import Vue from 'vue'
import Vuex from 'vuex'
import cookie from '../static/js/cookie'

Vue.use(Vuex)

const userInfo = {
  Id: cookie.getCookie('Id') || '',
  Username:cookie.getCookie('username') || '',
  token: cookie.getCookie('token') || '',
}

const state = {
  userInfo
}

export default new Vuex.Store(
  {
    state,
    mutations: {
      set_info (state) {
        state.userInfo = {
          Id: cookie.getCookie('Id'),
          token: cookie.getCookie('token'),
          Username: cookie.getCookie('username')
        }
      }
    },
    actions: {},
  }
)

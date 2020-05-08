import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isLogin: false,
    userID:"",
    userName: "",
  },
  mutations: {
    login(state, userInfo){
      state.isLogin = true
      state.userID = userInfo.userID
      state.userName = userInfo.userName
    },
    logout(state){
      state.isLogin = false
    }
  },
  actions: {
  },
  modules: {
  }
})

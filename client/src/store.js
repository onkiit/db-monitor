import Vue from 'vue'
import Vuex from 'vuex'
import http from './http'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    modal: false,
    modalTitle: "",
    modalCaller: "",
    redisHealth: {}
  },
  getters: {
    modal: state => state.modal,
    modalTitle: state => state.modalTitle,
    modalCaller: state => state.modalCaller,
    redisHealth: state => state.redisHealth
  },
  mutations: {
    setModal(state, payload){
      state.modal = payload
    },
    setModalTitle(state, payload){
      state.modalTitle = payload
    },
    setModalCaller(state, payload){
      state.modalCaller = payload
    },
    setRedisHealth(state, payload){
      state.redisHealth = payload
    }
  },
  actions: {
    setRedisHealth({commit}){
      http.get('/redis/health')
      .then(({data}) => {
        commit('setRedisHealth', data.redis_health)
      })
    }
  }
})

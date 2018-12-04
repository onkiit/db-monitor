import Vue from 'vue'
import Vuex from 'vuex'
import http from './http'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    modal: false,
    modalTitle: "",
    modalCaller: "",
    redisHealth: {},
    psqlHealth: {},
    mongoHealth: {}
  },
  getters: {
    modal: state => state.modal,
    modalTitle: state => state.modalTitle,
    modalCaller: state => state.modalCaller,
    redisHealth: state => state.redisHealth,
    psqlHealth: state => state.psqlHealth,
    mongoHealth: state => state.mongoHealth
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
    },
    setPsqlHealth(state, payload){
      state.psqlHealth = payload
    },
    setMongoHealth(state, payload){
      state.mongoHealth = payload
    }
  },
  actions: {
    setRedisHealth({commit}){
      http.get('/redis/health')
      .then(({data}) => {
        commit('setRedisHealth', data.redis_health)
      })
    },
    setPsqlHealth({commit}) {
      http.get('/postgres/health')
      .then(({data}) => {
        commit('setPsqlHealth', data.psql_health)
      })
    },
    setMongoHealth({commit}){
      http.get('/mongo/health')
      .then(({data}) => {
        commit('setMongoHealth', data.mongo_health)
      })
    }
  }
})

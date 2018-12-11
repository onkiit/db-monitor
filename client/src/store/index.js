import Vue from 'vue'
import Vuex from 'vuex'

import postgres from './modules/postgres.js'
import redis from './modules/redis.js'
import mongo from './modules/mongo.js'
import mysql from './modules/mysql.js'
import auth from './modules/auth.js'

Vue.use(Vuex)

const state = {
    modal: false,
    modal_title: "",
    modal_caller: ""
}

const getters = {
    modal: state => state.modal,
    modal_title: state => state.modal_title,
    modal_caller: state => state.modal_caller
}

const mutations = {
    setModal(state, payload){
        state.modal = payload
    },
    setModalTitle(state, payload){
        state.modal_title = payload
    },
    setModalCaller(state, payload){
        state.modal_caller = payload
    }
}

export const store = new Vuex.Store({
    state,
    getters,
    mutations,
    modules: {
        postgres: postgres,
        redis: redis,
        mongo: mongo,
        mysql: mysql,
        auth: auth
    }
})
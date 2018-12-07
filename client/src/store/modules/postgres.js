import http from '../../http'

const state = {
    version: "",
    active_client: 0,
    health: {},
    disabled: false
}

const getters = {
    disabled: state => state.disabled,
    version: state => state.version,
    active_client: state => state.active_client ,
    health: state => state.health
}

const mutations = {
    setDisabled(state, payload){
        state.disabled = payload
    },
    setVersion(state, payload){
        state.version = payload
    },
    setActiveClient(state, payload){
        state.active_client = payload
    },
    setHealth(state, payload) {
        state.health = payload
    }
}

const actions = {
    setVersion({commit}) {
        http.get('/postgres/version')
        .then(({data}) => {
            commit('setVersion', data.version)
        })
    },
    setActiveClient({commit}) {
        http.get('/postgres/client')
        .then(({data}) => {
            commit('setActiveClient', data.active_client)
        })
    },
    setHealth({commit}) {
        http.get('/postgres/health')
        .then(({data}) => {
            commit('setHealth', data.psql_health)
        })
        .catch(err => {
            if(err){
                commit('setDisabled', true)
            }
        })
    }
}

export default {
    namespaced: true,
    state,
    getters,
    mutations,
    actions
}
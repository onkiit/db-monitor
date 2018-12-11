const state = {
    isLoggedIn: false,
    user: {}
}

const getters = {
    isLoggedIn: state => state.isLoggedIn,
    user: state => state.user
}

const mutations = {
    setLoggedIn(state, payload){
        localStorage.setItem('isLoggedin', payload)
        state.isLoggedIn = payload
    },
    setUser(state, payload){
        localStorage.setItem('user', payload)
        state.user = payload
    }
}

export default {
    state,
    getters,
    mutations
}
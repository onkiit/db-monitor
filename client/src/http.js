import axios from 'axios'
import Vue from 'vue';


const http = axios.create({
    baseURL: "http://127.0.0.1:8180",
})

Vue.prototype.$http = http

export default http
import axios from 'axios'
import Vue from 'vue';

const http = axios.create({
    baseURL: "http://192.168.1.17:8180",
})

Vue.prototype.$http = http

export default http
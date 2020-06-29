// Import Vue
import Vue from 'vue';
import VueRouter from 'vue-router';

// Import Vue App, routes, store
import App from './App';
import routes from './routes';
import store from './store';
import i18n from './i18n'
import vuetify from '@/plugins/vuetify' // path to vuetify export
Vue.use(VueRouter);
Vue.use(require('vue-moment'));
const axios = require('axios').default;

import Fingerprint2 from "fingerprintjs2"
let hash = require('object-hash');

window.userId = -1;
if (window.requestIdleCallback) {
    requestIdleCallback(function () {
        Fingerprint2.get(function (components) {
            window.userId = hash(components);
        })
    })
} else {
    setTimeout(function () {
        Fingerprint2.get(function (components) {
            window.userId = hash(components);
        })
    }, 500)
}

Vue.prototype.askBackend = function (url, param) {
    console.log('ask smth!', process.env.BACK_SERVER);
    param.user_sign = window.userId;
    let config = {
        headers: {
            m: window.userId,
        }
    }
    return new Promise((resolve, reject) => {
        axios.post(`/api/${url}`, param, config)
            .then(resp => resolve(resp.data))
            .catch(error => {
                let code = +error.response.status;
                let message = ''
                switch (code) {
                    case 401:
                        this.$store.commit('setAuth', 0);
                        message = 'Unauthorized';
                        break;
                    case 409:
                        message = 'Already exist';
                        break;
                    case 400:
                        message = 'Bad request';
                        break;
                }
                if (message) {
                    this.$store.commit('setAlert', {
                        display: true,
                        text: message,
                        color: 'error',
                        delay: 5,
                    });
                }
                reject(error)
            })
    });
};

// Configure router
const router = new VueRouter({
    routes,
    linkActiveClass: 'active',
    mode: 'history'
});

console.log(process.env.BACK_SERVER);

new Vue({
    vuetify,
    i18n,
    store,
    el: '#app',
    render: h => h(App),
    router
});

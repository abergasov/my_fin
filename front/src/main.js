// Import Vue
import Vue from 'vue';
import VueRouter from 'vue-router';

// Import Vue App, routes, store
import App from './App';
import routes from './routes';
import vuetify from '@/plugins/vuetify' // path to vuetify export
Vue.use(VueRouter);


Vue.prototype.askBackend = function (param) {
    let domain = process.env.BACK_SERVER;
    console.log('ask smth!')
    console.log(domain, param);
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
    el: '#app',
    render: h => h(App),
    router
});

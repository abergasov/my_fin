import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        auth: 0,
        auth_expires: 0,
        auth_user: 0,
        dataLoading: false,
        categories: null,
        expenses: [],
        alertData: {
            display: false,
            text: '',
            color: '',
            delay: 0,
        }
    },
    mutations: {
        setLoading(state, payload) {
            state.dataLoading = payload;
        },

        setUserId(state, payload) {
            state.auth_user = +payload;
        },
        setAuthExpires(state, payload) {
            state.auth_expires = +payload;
        },
        setAuth (state, payload) {
            state.auth = payload;
            localStorage.auth = +payload;
        },
        setExpenses (state, payload) {
            state.expenses = payload;
        },
        setCategories (state, payload) {
            state.categories = payload;
        },
        setAlert (state, payload) {
            console.log('set alert');
            state.alertData = payload;
            setTimeout(() => {
                state.alertData.display = false;
                console.log('clear alert');
            }, (payload.delay || 5) * 1000)
        }
    },
    getters: {
        authGetter: state => {
            return state.auth;
        }
    }
})

export default store;

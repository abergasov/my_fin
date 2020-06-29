import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        auth: 0,
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
            }, payload.delay * 1000)
        }
    },
    getters: {
        authGetter: state => {
            return state.auth;
        }
    }
})

export default store;

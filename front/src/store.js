import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        categories: null,
        alertData: {
            display: false,
            text: '',
            color: '',
            delay: 0,
        }
    },
    mutations: {
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
    }
})

export default store;

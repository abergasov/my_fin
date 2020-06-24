import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        categories: null,
    },
    mutations: {
        setCategories (state, payload) {
            state.categories = payload;
        }
    }
})

export default store;

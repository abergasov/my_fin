import Vue from 'vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)

const opts = {
    theme: {
        themes: {
            light: {
                background: '#cccccc',
                primary: '#3f51b5',
                secondary: '#b0bec5',
                accent: '#8c9eff',
                error: '#b71c1c',
                warning: '#fb8c00',
                info: '#2196f3',
                success: '#4caf50',
            },
            dark: {
                background: '#555555',
                primary: '#3f51b5',
                secondary: '#b0bec5',
                accent: '#8c9eff',
                error: '#b71c1c',
                warning: '#fb8c00',
                info: '#2196f3',
                success: '#4caf50',
            }
        },
        dark: true,
    },
}

export default new Vuetify(opts)
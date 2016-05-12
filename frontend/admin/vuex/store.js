import Vue from 'vue';
import Vuex from 'vuex';
import songs from './modules/songs.js';

Vue.use(Vuex);

module.exports = new Vuex.Store({
    modules: {
        songs
    }
});
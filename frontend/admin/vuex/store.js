import Vue from 'vue';
import Vuex from 'vuex';
import songs from './modules/songs.js';
import singers from './modules/singers.js';

Vue.use(Vuex);

module.exports = new Vuex.Store({
    modules: {
        songs
    }
});
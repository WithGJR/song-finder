import Vue from 'vue';
import Vuex from 'vuex';
import songs from './modules/songs.js';
import singers from './modules/singers.js';
import albums from './modules/albums.js';

Vue.use(Vuex);

module.exports = new Vuex.Store({
  modules: {
    songs,
    singers,
    albums
  }
});
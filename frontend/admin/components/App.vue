<template>
  <div class="container">
  	<div class="page-header">
    	<h1 class="text-center">管理後台</h1>
      <nav-bar></nav-bar>
  	</div>
    <router-view></router-view>
  </div>
</template>

<script>
import Vue from 'vue';
import Vuex from 'vuex';
import store from '../vuex/store.js';
import NavBar from '../components/NavBar.vue';
import { fetchSongs, fetchSingers, fetchAlbums } from '../vuex/actions.js';

Vue.use(Vuex);

fetchSongs(store);
fetchSingers(store);
fetchAlbums(store);

module.exports = {
  store,
  components: {
    NavBar
  }
};
</script>

<template>
  <song-list :songs="songs"></song-list>
</template>

<script>
import SongList from './SongList.vue';
import { getSongs } from '../vuex/getters.js';
import { fetchSongs } from '../vuex/actions.js';

module.exports = {
  components: {
    SongList
  },
  vuex: {
    getters: {
      songs: getSongs
    },
    actions: {
      fetchSongs
    }
  },
  created(){
    this.fetchSongs();
  }
};
</script>

<template>
    <div>
    <div v-if="isLoading" class="text-center">
        <img src="/static/loading.gif" />
    </div>
    <div v-if="haveNoSongs" class="alert alert-warning">
        找不到歌曲
    </div>
    
    <ul class="list-group">
        <li v-for="(index, song) in songs" class="list-group-item">
            <h2>{{song.name}}</h2>
            <hr>
            <div>
                <a v-link="{path: '/songs/' + index + '/edit'}" class="btn btn-default">
                    <span class="glyphicon glyphicon-pencil"></span> 編輯
                </a>
                <a v-on:click="remove(index)" class="btn btn-danger">
                    <span class="glyphicon glyphicon-trash"></span> 刪除
                </a>
            </div>
        </li>
    </ul>
    </div>
</template>

<script>
import { getSongs } from '../vuex/getters.js';
import { fetchSongs, deleteSong } from '../vuex/actions.js';

module.exports = {
  vuex: {
    getters: {
      songs: getSongs,
      isLoading: ({ songs }) => songs.isLoading
    },
    actions: {
      fetchSongs,
      deleteSong
    }
  },
  computed: {
    haveNoSongs(){
      return !this.isLoading && this.songs.length === 0;
    }
  },
  methods: {
    remove(index){
      if (!window.confirm('確定要刪除嗎？')) {
        return;
      }
      this.deleteSong(index);
    }
  },
  created(){
    for (var i = 0; i<this.songs.length; i++) {
      console.log(i + ' :' + this.songs[i].name);
    }
  }
};
</script>

<template>
    <div>
    <div v-if="isLoading" class="text-center">
        <img src="/static/loading.gif" />
    </div>
    <div v-if="haveNoSongs" class="alert alert-warning">
        找不到歌曲
    </div>
    
    <ul class="list-group">
        <li v-for="song in songs" class="list-group-item">
            <h2>{{song.Name}}</h2>
            <hr>
            <div>
                <a v-link="{path: '/songs/' + song.ID + '/edit'}" class="btn btn-default">
                    <span class="glyphicon glyphicon-pencil"></span> 編輯
                </a>
                <a v-on:click="remove(song.ID)" class="btn btn-danger">
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
    remove(id){
      if (!window.confirm('確定要刪除嗎？')) {
        return;
      }
      this.deleteSong(id);
    }
  }
};
</script>

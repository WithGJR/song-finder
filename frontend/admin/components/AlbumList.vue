<template>
    <a v-link="{path: '/albums/new'}" class="btn btn-success btn-lg">
      <span class="glyphicon glyphicon-plus"></span> 新增專輯
    </a>
  
    <hr>

    <div>
    <div v-if="isLoading" class="text-center">
        <img src="/static/loading.gif" />
    </div>
    <div v-if="haveNoAlbums" class="alert alert-warning">
        目前還沒有新增專輯
    </div>

    <ul class="list-group">
        <li v-for="album in albums" class="list-group-item">
            <h2>{{album.Name}}</h2>
            <hr>
            <div>
                <a v-link="{path: '/albums/' + album.ID + '/edit'}" class="btn btn-default">
                    <span class="glyphicon glyphicon-pencil"></span> 編輯
                </a>
                <a v-link="{path: '/albums/' + album.ID + '/upload'}" class="btn btn-default">
                    <span class="glyphicon glyphicon-camera"></span> 上傳封面
                </a>
                <a v-on:click="remove(album.ID)" class="btn btn-danger">
                    <span class="glyphicon glyphicon-trash"></span> 刪除
                </a>
            </div>
        </li>
    </ul>
    </div>
</template>

<script>
import { getAlbums } from '../vuex/getters.js';
import { deleteAlbum } from '../vuex/actions.js';

module.exports = {
  vuex: {
    getters: {
      albums: getAlbums,
      isLoading: ({ albums }) => albums.isLoading
    },
    actions: {
      deleteAlbum
    }
  },
  computed: {
    haveNoAlbums(){
      return !this.isLoading && this.albums.length === 0;
    }
  },
  methods: {
    remove(id){
      if (!window.confirm('確定要刪除嗎？')) {
        return;
      }
      this.deleteAlbum(id);
    }
  }
};
</script>

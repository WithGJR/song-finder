<template>
    <a v-link="{path: '/albums/new'}" class="btn btn-success btn-lg">
      <span class="glyphicon glyphicon-plus"></span> 新增專輯
    </a>
  
    <hr>

    <div>
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
      albums: getAlbums
    },
    actions: {
      deleteAlbum
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

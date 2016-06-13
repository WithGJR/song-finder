<template>
    <a v-link="{path: '/singers/new'}" class="btn btn-success btn-lg">
      <span class="glyphicon glyphicon-plus"></span> 新增歌手
    </a>
  
    <hr>

    <div>
    <ul class="list-group">
        <li v-for="singer in singers" class="list-group-item">
            <h2>{{singer.Name}}</h2>
            <hr>
            <div>
                <a v-link="{path: '/singers/' + singer.ID + '/edit'}" class="btn btn-default">
                    <span class="glyphicon glyphicon-pencil"></span> 編輯
                </a>
                <a v-on:click="remove(singer.ID)" class="btn btn-danger">
                    <span class="glyphicon glyphicon-trash"></span> 刪除
                </a>
            </div>
        </li>
    </ul>
    </div>
</template>

<script>
import { getSingers } from '../vuex/getters.js';
import { deleteSinger } from '../vuex/actions.js';

module.exports = {
  vuex: {
    getters: {
      singers: getSingers
    },
    actions: {
      deleteSinger
    }
  },
  methods: {
    remove(id){
      if (!window.confirm('確定要刪除嗎？')) {
        return;
      }
      this.deleteSinger(id);
    }
  }
};
</script>

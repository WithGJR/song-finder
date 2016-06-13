<template>
  <div>
    <h2>{{title}}</h2>
    <form v-on:submit.prevent="submit">
      <div class="form-group">
        <label for="name" class="col-sm-2 control-label">專輯姓名</label>
        <input type="input" class="form-control" id="name" placeholder="專輯姓名" v-model="album.Name">
      </div>  

      <div class="form-group">
        <label for="country" class="col-sm-2 control-label">發行公司</label>
        <input type="input" class="form-control" id="country" placeholder="發行公司" v-model="album.Company">
      </div> 

      <div class="form-group">
        <label for="country" class="col-sm-2 control-label">發行日期</label>
        <input type="input" class="form-control" id="country" placeholder="發行日期" v-model="album.PublishedDate">
      </div>   

      <div class="form-group">
        <label for="country" class="col-sm-2 control-label">語言</label>
        <input type="input" class="form-control" id="country" placeholder="語言" v-model="album.Language">
      </div> 

      <div class="form-group">
        <label for="country" class="col-sm-2 control-label">得獎紀錄</label>
        <input type="input" class="form-control" id="country" placeholder="得獎紀錄" v-model="album.Reward">
      </div> 

      <input type="submit" class="btn btn-default" value="{{btnText}}" />
    </form>    
  </div>
</template>

<script>
import BaseForm from '../mixins/BaseForm.js';
import { addNewAlbum } from '../vuex/actions.js';
import { getAlbums } from '../vuex/getters.js';

const albumModel = {Name: '', Company: '', PublishedDate: '', Language: '', Reward: '', Photo: ''};

module.exports = {
  vuex: {
    actions: {
      addNewAlbum
    },
    getters: {
      albums: getAlbums
    }
  },
  mixins: [BaseForm],
  computed: {
    newFormTitle() {
      return '新增專輯';
    },
    editFormTitle() {
      return '修改專輯';
    },
    newFormBtnText() {
      return '確認新增';
    },
    editFormBtnText() {
      return '確認更新';
    },
    currentAlbum() {
      for (var i = 0; i < this.albums.length; i++) {
        if (this.albums[i].ID === this.currentId) {
          return this.albums[i];
        }
      }
    }
  },
  methods: {
    submit() {
      if (this.formType === 'edit') {
        
      }else{
        this.addNewAlbum(this.album);
        this.album = Object.assign({}, this.currentAlbum);
      }
    }
  },
  data(){
    return {
      album: Object.assign({}, albumModel)
    };
  },
  ready(){
    if (this.formType == 'edit') {
      this.album = Object.assign({}, this.currentAlbum);
    }
  }
};
</script>
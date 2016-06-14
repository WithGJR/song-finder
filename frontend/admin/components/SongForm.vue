<template>
  <div>
    <h2>{{title}}</h2>
    <form v-on:submit.prevent="submit">
      <div class="form-group">
        <label for="name" class="col-sm-2 control-label">曲名</label>
        <input type="input" class="form-control" id="name" placeholder="曲名" v-model="song.Name">
      </div>  

      <div class="form-group">
        <label for="lyricist" class="col-sm-2 control-label">作詞人</label>
        <input type="input" class="form-control" id="lyricist" placeholder="作詞人" v-model="song.Lyricist">
      </div>  

      <div class="form-group">
        <label for="composer" class="col-sm-2 control-label">作曲人</label>
        <input type="input" class="form-control" id="composer" placeholder="作曲人" v-model="song.Composer">
      </div> 

      <div class="form-group">
        <label for="length" class="col-sm-2 control-label">歌曲長度</label>
        <input type="input" class="form-control" id="length" placeholder="歌曲長度" v-model="song.Length">
      </div> 

      <div>
        <h3>歌手</h3>
        <label v-for="singer in singers" class="radio-inline">
          <div v-if="singer.ID === song.SingerId">
            <input v-model="song.SingerId" number type="radio" value="{{singer.ID}}" checked> {{singer.Name}}
          </div>
          <div v-else>
            <input v-model="song.SingerId" number type="radio" value="{{singer.ID}}"> {{singer.Name}}  
          </div>
        </label>
      </div>

      <div>
        <h3>專輯</h3>
        <label v-for="album in albums" class="radio-inline">
          <input v-model="song.AlbumId" number type="radio" value="{{album.ID}}"> {{album.Name}}
        </label>
      </div>

      <hr>

      <input type="submit" class="btn btn-default" value="{{btnText}}" />
    </form>
  </div>
</template>

<script>
import { getSongs, getSingers, getAlbums } from '../vuex/getters.js';
import { addNewSong, updateSong } from '../vuex/actions.js';
import BaseForm from '../mixins/BaseForm.js';

var songModel = {Name: '', Lyricist: '', Composer: '', Length: '', SingerId: null, AlbumId: null};

module.exports = {
  vuex: {
    getters: {
      songs: getSongs,
      singers: getSingers,
      albums: getAlbums
    },
    actions: {
      addNewSong,
      updateSong
    }
  },
  mixins: [BaseForm],
  data(){
    return {
      song: Object.assign({}, songModel)
    };
  },
  computed: {
    newFormTitle() {
      return '新增歌曲';
    },
    editFormTitle() {
      return '修改歌曲';
    },
    newFormBtnText() {
      return '確認新增';
    },
    editFormBtnText() {
      return '確認更新';
    },
    currentSong(){
      for (var i = 0; i < this.songs.length; i++) {
        if (this.songs[i].ID === this.currentId) {
          return this.songs[i];
        }
      }
    }
  },
  methods: {
    isValid(){
      return this.song.Name !== ''; 
    },
    submit(){
      if(!this.isValid()){
        return;
      }
      if (this.formType === 'edit') {
        this.updateSong(this.currentId, this.song);
      }else{
        this.addNewSong(this.song);
        this.song = Object.assign({}, songModel)
      }
    }
  },
  ready(){
    if (this.formType == 'edit') {
      this.song = Object.assign({}, this.currentSong);
    }
  }
};
</script>

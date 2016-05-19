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

      <input type="submit" class="btn btn-default" value="{{btnText}}" />
    </form>
  </div>
</template>

<script>
import { getSongs } from '../vuex/getters.js';
import { addNewSong, updateSong } from '../vuex/actions.js';

module.exports = {
  vuex: {
    getters: {
      songs: getSongs
    },
    actions: {
      addNewSong,
      updateSong
    }
  },
  data(){
    return {
      song: {Name: '', Lyricist: '', Composer: ''}
    };
  },
  computed: {
    formType(){
      return (this.$route.params.id === undefined) ? 'new' : 'edit';
    },
    title(){
      return (this.formType === 'new') ? '新增歌曲' : '修改歌曲';
    },
    btnText(){
      return (this.formType === 'new') ? '確認新增' : '確認更新';
    },
    currentId(){
      return parseInt(this.$route.params.id);
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
        this.song = {name: ''};
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

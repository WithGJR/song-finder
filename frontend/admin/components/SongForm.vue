<template>
  <div>
    <h2>{{title}}</h2>
    <form v-on:submit.prevent="submit">
      <div class="form-group">
        <label for="name" class="col-sm-2 control-label">曲名</label>
        <input type="input" class="form-control" id="name" placeholder="曲名" v-model="song.name">
      </div>  

      <input type="submit" class="btn btn-default" value="{{btnText}}" />
    </form>
  </div>
</template>

<script>
import { getSongs } from '../vuex/getters.js';
import { addNewSong, editSong } from '../vuex/actions.js';

module.exports = {
  vuex: {
    getters: {
      songs: getSongs
    },
    actions: {
      addNewSong,
      editSong
    }
  },
  data(){
    return {
      song: {name: ''}
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
      return this.songs[this.currentId];
    }
  },
  methods: {
    isValid(){
      return this.song.name !== ''; 
    },
    submit(){
      if(!this.isValid()){
        return;
      }
      if (this.formType === 'edit') {
        this.editSong(this.currentId, this.song);
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

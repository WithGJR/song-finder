<template>
  <h2>新增新歌曲</h2>
  <form v-on:submit.prevent="submitNewSong">
    <div class="form-group">
      <label for="name" class="col-sm-2 control-label">曲名</label>
      <input type="input" class="form-control" id="name" placeholder="Song Name" v-model="name">
    </div>  

    <input type="submit" class="btn btn-default" value="Add" />
  </form>
</template>

<script>
import {addNewSong} from '../vuex/actions.js';

module.exports = {
  vuex: {
    actions: {
      addNewSong
    }
  },
  data: () => {
    return {
      name: ''
    };
  },
  methods: {
    isValid(){
      return this.name !== ''; 
    },
    submitNewSong(){
      if(!this.isValid()){
        return;
      };
      this.addNewSong({name: this.name});
      this.name = '';
    }
  }
};
</script>

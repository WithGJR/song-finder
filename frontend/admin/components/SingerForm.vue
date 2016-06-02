<template>
  <div>
    <h2>{{title}}</h2>
    <form v-on:submit.prevent="submit">
      <div class="form-group">
        <label for="name" class="col-sm-2 control-label">歌手姓名</label>
        <input type="input" class="form-control" id="name" placeholder="歌手姓名" v-model="singer.Name">
      </div>  

      <div class="form-group">
        <label for="country" class="col-sm-2 control-label">國家</label>
        <input type="input" class="form-control" id="country" placeholder="國家" v-model="singer.Country">
      </div>  

     

      <input type="submit" class="btn btn-default" value="{{btnText}}" />
    </form>
  </div>
</template>

<script>
import BaseForm from '../mixins/BaseForm.js';
import { addNewSinger } from '../vuex/actions.js';

module.exports = {
  vuex: {
    actions: {
      addNewSinger
    }
  },
  mixins: [BaseForm],
  computed: {
    newFormTitle() {
      return '新增歌手';
    },
    editFormTitle() {
      return '修改歌手';
    },
    newFormBtnText() {
      return '確認新增';
    },
    editFormBtnText() {
      return '確認更新';
    }
  },
  methods: {
    submit() {
      this.addNewSinger(this.singer);
      this.singer = {Name: '', Country: ''};
    }
  },
  data(){
    return {
      singer: {Name: '', Country: ''}
    };
  },
  ready(){
    
  }
};
</script>
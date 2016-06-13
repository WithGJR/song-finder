<template>
  <div>
    <h2>上傳「{{album.Name}}」的封面</h2>
    <form v-on:submit.prevent="submit">
      <div class="form-group">
        <label for="file">封面照片</label>
        <input type="file" id="file" v-on:change="onFileChange">
        <p class="help-block">請靜待至上傳完畢</p>
      </div>

      <input type="submit" class="btn btn-default btn-lg" value="確認上傳" />
    </form>

    <hr>
    <h2>若上傳成功，照片會顯示於此</h2>
    <img v-bind:src="photoURL" class="img-responsive" v-show="this.album.Photo !== ''" />    
  </div>
</template>

<script>
import BaseForm from '../mixins/BaseForm.js';
import { uploadAlbumPhoto } from '../vuex/actions.js';
import { getAlbums } from '../vuex/getters.js';

module.exports = {
  vuex: {
    actions: {
      uploadAlbumPhoto
    },
    getters: {
      albums: getAlbums
    }
  },
  mixins: [BaseForm],
  computed: {
    album() {
      for (var i = 0; i < this.albums.length; i++) {
        if (this.albums[i].ID === this.currentId) {
          return this.albums[i];
        }
      }
    },
    photoURL() {
      return '/static/images/' + this.album.Photo;
    }
  },
  methods: {
    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      this.file = files[0];
    },
    submit() {
      this.uploadAlbumPhoto(this.album.ID, this.file);
    }
  },
  data(){
    return {
      file: ''
    };
  }
};
</script>
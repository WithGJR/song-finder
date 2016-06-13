import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const router = new VueRouter();

router.map({
  '/': {
    component: require('./components/SongList.vue')
  },
  '/songs/new': {
    component: require('./components/SongForm.vue')
  },
  '/songs/:id/edit': {
    component: require('./components/SongForm.vue')
  },
  '/singers': {
    component: require('./components/SingerList.vue')
  },
  '/singers/new': {
    component: require('./components/SingerForm.vue')
  },
  '/singers/:id/edit': {
    component: require('./components/SingerForm.vue')
  },
  '/albums': {
    component: require('./components/AlbumList.vue')
  },
  '/albums/new': {
    component: require('./components/AlbumForm.vue')
  },
  '/albums/:id/edit': {
    component: require('./components/AlbumForm.vue')
  },
  '/albums/:id/upload': {
    component: require('./components/UploadAlbumPhotoForm.vue')
  }
});

const App = Vue.extend(require('./components/App.vue'));
router.start(App, '#app');

import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const router = new VueRouter();

router.map({
    '/': {
        component: require('./components/AdminPanel.vue')
    },
    '/songs/new': {
        component: require('./components/NewSongForm.vue')
    }
});

const AdminApp = Vue.extend(require('./components/AdminApp.vue'));
router.start(AdminApp, '#app');

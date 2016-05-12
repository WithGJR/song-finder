<template>
  <div class="container">
  	<div class="page-header">
    	<h1 class="text-center">管理後台</h1>
    	<ul class="nav nav-pills">
  			<li role="presentation" class="active"><a v-link="{path: '/'}">歌曲清單</a></li>
  			<li role="presentation"><a v-link="{path: '/songs/new'}">新增歌曲</a></li>
		</ul>
  	</div>
    <router-view></router-view>
  </div>
</template>

<script>
import Vue from 'vue';
import Vuex from 'vuex';
import store from '../vuex/store.js';

Vue.use(Vuex);

module.exports = {
  store
};
</script>

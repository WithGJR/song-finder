<template>
	<ul class="nav nav-pills">
		<li role="presentation" class="active">
			<a v-link="{path: '/'}">歌曲清單</a>
		</li>
		<li role="presentation">
			<a v-link="{path: '/songs/new'}">新增歌曲</a>
		</li>
		<li role="presentation">
			<a v-link="{path: '/singers/new'}">新增歌手</a>
		</li>
		<li role="presentation">
			<a v-link="{path: '/singers'}">歌手清單</a>
		</li>
	</ul>
</template>

module.exports = {};
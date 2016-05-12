<template>
    <div v-if="isLoading" class="text-center">
        <img src="/static/loading.gif" />
    </div>
    <div v-if="haveNoSongs" class="alert alert-warning">
        找不到歌曲
    </div>
    
    <ul class="list-group">
        <li v-for="song in songs" class="list-group-item">
            {{song.name}}
        </li>
    </ul>
    
</template>

<script>
module.exports = {
    props: ['songs'],
    vuex: {
        getters: {
            isLoading: ({ songs }) => songs.isLoading
        }
    },
    computed: {
        haveNoSongs(){
            return !this.isLoading && this.songs.length === 0;
        }
    }
};
</script>

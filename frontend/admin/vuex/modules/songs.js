import types from '../mutation-types.js';

const state = {
	all: [],
	isLoading: false
};

const mutations = {
    [types.FETCH_SONGS_REQEUST](state){
        state.all = [];
        state.isLoading = true;
    },
    [types.FETCH_SONGS_SUCESS](state, songs){
        state.all = songs; 
        state.isLoading = false;
    },
    [types.FETCH_SONGS_FAILURE](state){
    	
    },
    [types.ADD_NEW_SONG](state, song){
    	state.all.push(song);
    },
    [types.EDIT_SONG](state, id, newSong){
        //TODO
        state.all[id] = newSong;
    },
    [types.DELETE_SONG](state, id){
        //TODO
        state.all.splice(id, 1);
    }
};

module.exports = {
	state,
	mutations
};
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
    	// TODO
    	state.all.push(song);
    }
};

module.exports = {
	state,
	mutations
};
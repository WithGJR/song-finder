import types from '../mutation-types.js';

const state = {
  all: []
};

const mutations = {
  [types.FETCH_SINGERS_SUCESS](state, singers){
    state.all = singers;
  },
  [types.ADD_NEW_SINGER](state, singer){
    state.all.push(singer);
  },
  [types.UPDATE_SINGER_SUCESS](state, id, newSinger){
    for (var i = 0; i < state.all.length; i++) {
      if (state.all[i].ID === id) {
        state.all.splice(i, 1, newSinger);
        return;
      }
    }
  },
  [types.DELETE_SINGER](state, id){
    for (var i = 0; i < state.all.length; i++) {
      if (state.all[i].ID === id) {
        state.all.splice(i, 1);
        return;
      }
    }
    
  }
};

module.exports = {
  state,
  mutations
};
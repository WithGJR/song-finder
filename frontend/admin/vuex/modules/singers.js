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
        state.all[i] = newSinger;
        return;
      }
    }
  },
  [types.DELETE_SINGER](state, id){
    state.all.splice(id, 1);
  }
};

module.exports = {
  state,
  mutations
};
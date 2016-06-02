import types from '../mutation-types.js';

const state = {
  all: []
};

const mutations = {
  [types.ADD_NEW_SINGER](state, singer){
    state.all.push(singer);
  },
};

module.exports = {
  state,
  mutations
};
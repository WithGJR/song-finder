import types from '../mutation-types.js';

const state = {
  all: []
};

const mutations = {
  [types.FETCH_ALBUMS_SUCESS](state, albums){
    state.all = albums;
  },
  [types.ADD_NEW_ALBUM](state, album){
    state.all.push(album);
  },
  [types.UPLOAD_ALBUM_PHOTO_SUCESS](state, upadatedAlbum){
    for (var i = 0; i < state.all.length; i++) {
      if (state.all[i].ID === upadatedAlbum.ID) {
        state.all.splice(i, 1, upadatedAlbum);
      }
    }
  },
  [types.DELETE_ALBUM](state, deletedAlbumId){
    for (var i = 0; i < state.all.length; i++) {
      if (state.all[i].ID === deletedAlbumId) {
        state.all.splice(i, 1);
      }
    }
  }
};

module.exports = {
  state,
  mutations
};
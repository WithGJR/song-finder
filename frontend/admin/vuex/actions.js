import api from '../api/';
import types from './mutation-types.js';

module.exports = {
  fetchSongs({dispatch}){
    dispatch(types.FETCH_SONGS_REQEUST);
    
    api.fetchAllSongs(data => {
      dispatch(types.FETCH_SONGS_SUCESS, data); 
    });
  },
  addNewSong({dispatch}, newSong){
    api.addNewSong(newSong, data => {
      dispatch(types.ADD_NEW_SONG, data); 
    });
  },
  updateSong({dispatch}, id, song){
    api.updateSong(id, song, updatedSong => {
      dispatch(types.UPDATE_SONG_SUCESS, updatedSong.ID, updatedSong);  
    });
  },
  deleteSong({dispatch}, id){
    api.deleteSong(id, data => {
      dispatch(types.DELETE_SONG, data.id);
    });
  }
};

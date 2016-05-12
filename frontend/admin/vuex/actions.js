import api from '../api/';
import type from './mutation-types.js';

module.exports = {
  fetchSongs({dispatch}){
    dispatch(type.FETCH_SONGS_REQEUST);
    
    api.fetchAllSongs(data => {
      dispatch(type.FETCH_SONGS_SUCESS, data); 
    });
  },
  addNewSong({dispatch}, newSong){
    api.addNewSong(newSong, data => {
      dispatch('ADD_NEW_SONG', data); 
    });
  }
};

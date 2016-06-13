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
  },

  fetchSingers({dispatch}){
    api.fetchAllSingers(data => {
      dispatch(types.FETCH_SINGERS_SUCESS, data); 
    });
  },
  addNewSinger({dispatch}, newSinger){
    api.addNewSinger(newSinger, data => {
      dispatch(types.ADD_NEW_SINGER, data);
    });
  },
  updateSinger({dispatch}, id, singer){
    api.updateSinger(id, singer, updatedSinger => {
      dispatch(types.UPDATE_SINGER_SUCESS, updatedSinger.ID, updatedSinger);  
    });
  },
  deleteSinger({dispatch}, id){
    api.deleteSinger(id, data => {
      dispatch(types.DELETE_SINGER, data.id);
    });
  }
};

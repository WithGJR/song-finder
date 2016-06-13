import api from '../api/';
import types from './mutation-types.js';

module.exports = {
  fetchSongs({dispatch}){
    dispatch(types.FETCH_SONGS_REQEUST);
    
    api.fetchAllSongs(songs => {
      dispatch(types.FETCH_SONGS_SUCESS, songs); 
    });
  },
  addNewSong({dispatch}, newSong){
    api.addNewSong(newSong, song => {
      dispatch(types.ADD_NEW_SONG, song); 
    });
  },
  updateSong({dispatch}, id, song){
    api.updateSong(id, song, updatedSong => {
      dispatch(types.UPDATE_SONG_SUCESS, updatedSong.ID, updatedSong);  
    });
  },
  deleteSong({dispatch}, id){
    api.deleteSong(id, deletedSongId => {
      dispatch(types.DELETE_SONG, deletedSongId);
    });
  },

  fetchSingers({dispatch}){
    api.fetchAllSingers(singers => {
      dispatch(types.FETCH_SINGERS_SUCESS, singers); 
    });
  },
  addNewSinger({dispatch}, newSinger){
    api.addNewSinger(newSinger, singer => {
      dispatch(types.ADD_NEW_SINGER, singer);
    });
  },
  updateSinger({dispatch}, id, singer){
    api.updateSinger(id, singer, updatedSinger => {
      dispatch(types.UPDATE_SINGER_SUCESS, updatedSinger.ID, updatedSinger);  
    });
  },
  deleteSinger({dispatch}, id){
    api.deleteSinger(id, deletedSingerId => {
      dispatch(types.DELETE_SINGER, deletedSingerId);
    });
  },

  fetchAlbums({dispatch}){
    api.fetchAllAlbums(albums => {
      dispatch(types.FETCH_ALBUMS_SUCESS, albums); 
    });
  },
  addNewAlbum({dispatch}, newAlbum){
    api.addNewAlbum(newAlbum, album => {
      dispatch(types.ADD_NEW_ALBUM, album); 
    });
  },
  uploadAlbumPhoto({dispatch}, albumId, photo){
    api.uploadAlbumPhoto(albumId, photo, updatedAlbum => {
      dispatch(types.UPLOAD_ALBUM_PHOTO_SUCESS, updatedAlbum); 
    });
  },
  deleteAlbum({dispatch}, id){
    api.deleteAlbum(id, deletedAlbumId => {
      dispatch(types.DELETE_ALBUM, deletedAlbumId);
    });
  }
};

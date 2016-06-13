module.exports = {
  getSongs({ songs }){
    return songs.all;
  },
  getSingers({ singers }){
    return singers.all;
  },
  getAlbums({ albums }){
    return albums.all;  
  }
};
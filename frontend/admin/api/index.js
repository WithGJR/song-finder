import 'whatwg-fetch';

module.exports = {
  fetchAllSongs(callback){
    fetch('http://localhost:7777/songs').then(data => {
      data.json().then(json => {
        callback(json); 
      });
    }).catch(error => {
    }); 
  },

  addNewSong(song, callback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch('http://localhost:7777/songs', {
      method: 'post',
      headers: headers,
      body: JSON.stringify(song)     
    }).then(data => {
      data.json().then(json => {
        callback(json); 
      });
    }).catch(error => {
    }); 
  }
};

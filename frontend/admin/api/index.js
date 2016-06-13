import 'whatwg-fetch';

module.exports = {
  fetchAllSongs(successCallback){
    fetch('http://localhost:7777/songs').then(data => {
      data.json().then(json => {
        successCallback(json); 
      });
    }).catch(error => {
    }); 
  },

  addNewSong(song, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch('http://localhost:7777/songs', {
      method: 'post',
      headers: headers,
      body: JSON.stringify(song)     
    }).then(data => {
      data.json().then(json => {
        successCallback(json); 
      });
    }).catch(error => {
    }); 
  },

  updateSong(id, song, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch('http://localhost:7777/songs/'+id, {
      method: 'put',
      headers: headers,
      body: JSON.stringify(song)     
    }).then(data => {
      data.json().then(updatedSong => {
        successCallback(updatedSong); 
      });
    }).catch(error => {
    }); 
  },

  deleteSong(id, successCallback){  
    fetch('http://localhost:7777/songs/'+id, {
      method: 'delete'
    }).then(data => {
      data.json().then(id => {
        successCallback(id); 
      });
    }).catch(error => {
    }); 
  },

  fetchAllSingers(successCallback){
    fetch('http://localhost:7777/singers').then(data => {
      data.json().then(json => {
        successCallback(json); 
      });
    }).catch(error => {
    }); 
  },
  addNewSinger(singer, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch('http://localhost:7777/singers', {
      method: 'post',
      headers: headers,
      body: JSON.stringify(singer)     
    }).then(data => {
      data.json().then(json => {
        successCallback(json); 
      });
    }).catch(error => {
    }); 
  },
  updateSinger(id, singer, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch('http://localhost:7777/singers/'+id, {
      method: 'put',
      headers: headers,
      body: JSON.stringify(singer)     
    }).then(data => {
      data.json().then(updatedSinger => {
        successCallback(updatedSinger); 
      });
    }).catch(error => {
    }); 
  },
  deleteSinger(id, successCallback){
    fetch('http://localhost:7777/singers/'+id, {
      method: 'delete'
    }).then(data => {
      data.json().then(id => {
        successCallback(id); 
      });
    }).catch(error => {
    }); 
  }
};

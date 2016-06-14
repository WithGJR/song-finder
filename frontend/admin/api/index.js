import 'whatwg-fetch';

const _domainName = 'http://localhost:7777';

module.exports = {
  fetchAllSongs(successCallback){
    fetch(_domainName + '/songs').then(data => {
      data.json().then(songs => {
        successCallback(songs); 
      });
    }).catch(error => {
    }); 
  },

  addNewSong(song, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    if(song.SingerId === null) {
      delete song.SingerId;
    }
    if(song.AlbumId === null) {
      delete song.AlbumId;
    }

    fetch(_domainName + '/songs', {
      method: 'post',
      headers: headers,
      body: JSON.stringify(song)     
    }).then(data => {
      data.json().then(createdSong => {
        successCallback(createdSong); 
      });
    }).catch(error => {
    }); 
  },

  updateSong(id, song, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch(_domainName + '/songs/' + id, {
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
    fetch(_domainName + '/songs/'+id, {
      method: 'delete'
    }).then(data => {
      data.json().then(id => {
        successCallback(id); 
      });
    }).catch(error => {
    }); 
  },

  fetchAllSingers(successCallback){
    fetch(_domainName + '/singers').then(data => {
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

    fetch(_domainName + '/singers', {
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

    fetch(_domainName + '/singers/' + id, {
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
    fetch(_domainName + '/singers/'+id, {
      method: 'delete'
    }).then(data => {
      data.json().then(id => {
        successCallback(id); 
      });
    }).catch(error => {
    }); 
  },

  fetchAllAlbums(successCallback){
    fetch(_domainName + '/albums').then(data => {
      data.json().then(albums => {
        successCallback(albums); 
      });
    }).catch(error => {
    }); 
  },
  uploadAlbumPhoto(albumId, photo, successCallback){
    var formData = new FormData();
    formData.append('file', photo);

    fetch(_domainName + '/albums/' + albumId + '/image-upload', {
       method: 'post',
       body: formData
    }).then(data => {
      data.json().then(updatedAlbum => {
        successCallback(updatedAlbum); 
      });
    }).catch(error => {
    }); 
  },
  addNewAlbum(album, successCallback){
    var headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');

    fetch(_domainName + '/albums', {
      method: 'post',
      headers: headers,
      body: JSON.stringify(album)     
    }).then(data => {
      data.json().then(json => {
        successCallback(json); 
      });
    }).catch(error => {
    }); 
  },
  deleteAlbum(id, successCallback){
    fetch(_domainName + '/albums/' + id, {
      method: 'delete'
    }).then(data => {
      data.json().then(id => {
        successCallback(id); 
      });
    }).catch(error => {
    }); 
  }
};

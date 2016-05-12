module.exports = {
  entry: { 
    admin: './admin/main.js'
  },
  output: {
    path: './build',
    filename: '[name].bundle.js' 
  },
  module: {
    loaders: [
      {
        test: /\.vue$/,
        loader: 'vue'
      },
      {
        test: /\.js$/,
        loader: 'babel',
        exclude: /node_modules/,
        query: {
          presets: ['es2015'] 
        }
      }   
    ]         
  }
};

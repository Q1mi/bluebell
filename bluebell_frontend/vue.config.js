module.exports = {
    devServer: {
        proxy: {
            '/api/v1': {
              target: 'http://127.0.0.1:8080',
              changeOrigin: true,
            }
        }
    }
  }

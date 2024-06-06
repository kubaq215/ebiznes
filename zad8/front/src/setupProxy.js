const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/auth',
    createProxyMiddleware({
      target: 'http://172-104-249-16.ip.linodeusercontent.com:5000',
      changeOrigin: true,
    })
  );
};

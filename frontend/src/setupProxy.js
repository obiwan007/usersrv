const { createProxyMiddleware } = require('http-proxy-middleware');
module.exports = function (app) {
    app.use(
        '/auth/login',
        createProxyMiddleware({
            target: 'http://localhost:8090',
            changeOrigin: true,
        })
    );
    app.use(
        '/auth/callback',
        createProxyMiddleware({
            target: 'http://localhost:8090',
            changeOrigin: true,
        })
    );
    app.use(
        '/auth/refresh',
        createProxyMiddleware({
            target: 'http://localhost:8090',
            changeOrigin: true,
        })
    );
    app.use(
        '/query',
        createProxyMiddleware({
            target: 'http://localhost:8090',
            changeOrigin: true,
        })
    );
    app.use(
        '/graphiql',
        createProxyMiddleware({
            target: 'http://localhost:8090/',
            changeOrigin: true,
        })
    );
};
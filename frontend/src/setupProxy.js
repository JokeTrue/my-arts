const {createProxyMiddleware} = require("http-proxy-middleware");

module.exports = function (app) {
    let target = "http://localhost:8080";
    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
        target = "https://my-arts.herokuapp.com";
    }
    app.use(
        "/api",
        createProxyMiddleware({
            target: target,
            changeOrigin: true,
        })
    );
};
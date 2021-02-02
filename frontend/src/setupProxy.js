const {createProxyMiddleware} = require("http-proxy-middleware");

module.exports = function (app) {
    let target = "https://my-arts.herokuapp.com";
    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
        target = "http://localhost:8080";
    }
    app.use(
        "/api",
        createProxyMiddleware({
            target: target,
            changeOrigin: true,
        })
    );
};
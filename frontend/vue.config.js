module.exports = {
    devServer: {
        port: 3333, // 启动端口
        open: true,  // 启动后是否自动打开网页
        proxy: {
            '/api/v1/*': {
                target: 'http://127.0.0.1:9090', //设置你调用的接口域名和端口
                changeOrigin: true, //true表示实现跨域
                pathRewrite: {
                    '^/api/v1': '/'
                }
            }
        }
    },

}

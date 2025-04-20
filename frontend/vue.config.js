module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:7001',
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    },
    // 添加以下配置解决开发服务器访问问题
    host: '0.0.0.0',
    port: 5173,
    open: true
  }
}
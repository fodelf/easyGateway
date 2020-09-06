/*
 * @Description: 描述
 * @Author: 吴文周
 * @Github: http://gitlab.yzf.net/wuwenzhou
 * @Date: 2019-11-19 08:46:03
 * @LastEditors: pym
 * @LastEditTime: 2020-09-06 18:40:52
 */
const path = require('path')
const ispro = process.env.NODE_ENV !== 'development'
function resolve(dir) {
  return path.join(__dirname, dir)
}
module.exports = {
  pages: {
    // 多页面时可以配置按需打包
    index: {
      // page 的入口
      entry: 'src/pages/index/main.js',
      // 模板来源
      template: 'src/pages/index/index.html',
      // 在 dist/index.html 的输出
      filename: 'index.html',
      // 当使用 title 选项时，
      // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
      title: 'EasyWork',
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    }
  },
  publicPath: ispro ? '' : '/',
  outputDir: '../gateway-server/web',
  assetsDir: 'static',
  lintOnSave: false,
  productionSourceMap: false,
  devServer: {
    port:9527,
    proxy: {
      '/uiApi': {
        target: 'http://192.168.0.105:9523'
      },
      '/socket.io': {
        target: 'http://localhost:9526',
        ws:true
      },
    },
    before:(app) =>{
      app.all('*', function (req, res, next) {
        res.header('Access-Control-Allow-Origin', '*')
        res.header('Access-Control-Allow-Headers', 'X-Requested-With')
        res.header('Access-Control-Allow-Methods', 'PUT,POST,GET,DELETE,OPTIONS')
        res.header('X-Powered-By', ' 3.2.1')
        res.header('Cache-Control', 'no-cache, no-store, must-revalidate')
        res.header('Pragma', 'no-cache');
        res.header('Expires', '0')
        next()
      })
    }
  },
  chainWebpack: config => {
    config.resolve.alias
      .set('@', resolve('src'))
      .set('assets', resolve('src/assets'))
      .set('components', resolve('src/components'))
      .set('base', resolve('baseConfig'))
      .set('public', resolve('public'))
  }
}

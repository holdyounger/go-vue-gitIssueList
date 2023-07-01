import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' // 需安装此模块
import AutoImport from "unplugin-auto-import/vite"
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // AutoImport ({
    //   imports: ["vue", "vue-router"], // 自动导入vue和vue-router相关函数
    //   dts: "src/auto-import.d.ts", // 生成 `auto-import.d.ts` 全局声明
    //   resolvers: [ElementPlusResolver()],
    // }),
    // Components({
    //   resolvers: [ElementPlusResolver()],
    // }),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],

  resolve: {
    alias: {  // 这里就是需要配置resolve里的别名
      "@": path.join(__dirname, "./src"), // path记得引入
      // 'vue': 'vue/dist/vue.esm-bundler.js' // 定义vue的别名，如果使用其他的插件，可能会用到别名
    },
  },
  server: {
    port: 3001, // 设置服务启动端口号
    open: false, // 设置服务启动时是否自动打开浏览器
    // 代理
    // proxy: {
    //   '/api': {
    //     target: 'http://API网关所在域名',
    //     changeOrigin: true,
    //     rewrite: (path) => path.replace(/^\/api/, '')
    //   },
    // }
  },
  esbuild: {
    jsxFactory: 'h',
    jsxFragment: 'Fragment',
    jsxInject: "import { h } from 'vue';"
  },
  configureWebpack: {
    configureWebpack :{
      devtools: 'source-map'
    }
  }

})

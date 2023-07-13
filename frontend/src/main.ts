import {createApp} from 'vue'

import App from './App.vue'
import { store, key } from './store'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'
import router from '@/router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import setupComponent from '@/components'
import '@/styles/style.scss' // global css
import 'virtual:windi.css'
import 'virtual:windi-devtools'

import 'element-plus/theme-chalk/src/dark/css-vars.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'element-plus/theme-chalk/index.css';

const app = createApp(App)

//注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)
app.use(store, key)
app.use(ElementPlus, { size: 'default',locale: zhCn });

//全局组件注册
setupComponent(app)

app.mount('#app')
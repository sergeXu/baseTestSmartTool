import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from 'axios'

const app = createApp(App)

// 创建axios实例并挂载到全局属性
const axiosInstance = axios.create({
  baseURL: process.env.NODE_ENV === 'development' ? '/api' : '/'
})

// 方式1：挂载到全局属性（Options API使用）
app.config.globalProperties.$axios = axiosInstance

// 方式2：使用provide/inject（Composition API推荐）
app.provide('axios', axiosInstance)

app.use(router)
app.use(ElementPlus)
app.mount('#app')
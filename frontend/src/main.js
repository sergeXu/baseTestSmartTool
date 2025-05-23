import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from 'axios'

const app = createApp(App)
app.use(router)
app.use(ElementPlus)

// 全局axios配置
app.config.globalProperties.$axios = axios.create({
    baseURL: process.env.NODE_ENV === 'development' ? '/api' : '/'
})

app.mount('#app')
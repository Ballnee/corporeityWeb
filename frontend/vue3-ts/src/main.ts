import { createApp } from 'vue'
import { createPinia } from 'pinia'

// import ElementUI from 'element-plus';
// import 'element-ui/lib/theme-chalk/index.css';

import App from './App.vue'
import router from './router'

import './assets/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
// app.use(ElementUI)  //全局引入
app.mount('#app')

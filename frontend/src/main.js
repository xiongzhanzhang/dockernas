import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import i18n from './locales/index'

const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.use(i18n)

app.mount('#app')

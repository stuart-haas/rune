import { createApp } from 'vue'
import './assets/index.css'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faPenToSquare, faTrash, faTableCells, faTableList } from '@fortawesome/free-solid-svg-icons'
import router from './router'

library.add(faPenToSquare, faTrash, faTableCells, faTableList)

const app = createApp(App)

app.use(VueQueryPlugin)
app.use(router)
app.mount('#app')
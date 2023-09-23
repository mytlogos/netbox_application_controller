import 'primevue/resources/themes/bootstrap4-dark-blue/theme.css'
import 'primeflex/primeflex.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Menu from 'primevue/menu'
import Divider from 'primevue/divider'
import Menubar from 'primevue/menubar'
import Card from 'primevue/card';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';


import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue)
app.component('InputText', InputText)
app.component('p-button', Button)
app.component('p-menu', Menu)
app.component('p-divider', Divider)
app.component('p-menubar', Menubar)
app.component('p-card', Card)
app.component('p-data-table', DataTable)
app.component('p-column', Column)

app.mount('#app')

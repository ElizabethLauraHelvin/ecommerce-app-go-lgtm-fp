import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import { initFaro } from './observability/faro'

// 1. Init Faro PERTAMA sebelum apapun
//    Semua event, trace, error sejak awal sudah ter-capture
initFaro()

// 2. Router
import HomeView from './views/HomeView.vue'
import ProductView from './views/ProductView.vue'
import CartView from './views/CartView.vue'
import OrderView from './views/OrderView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/',             name: 'home',    component: HomeView },
    { path: '/products/:id', name: 'product', component: ProductView },
    { path: '/cart',         name: 'cart',    component: CartView },
    { path: '/orders/:id',   name: 'order',   component: OrderView },
  ],
})

// 3. Mount
const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')

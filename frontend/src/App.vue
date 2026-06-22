<template>
  <div id="app">
    <header>
      <h1>🛍️ E-Commerce Platform</h1>
      <nav>
        <button @click="currentPage = 'products'">Products</button>
        <button @click="currentPage = 'orders'">Orders</button>
        <button @click="currentPage = 'users'">Users</button>
      </nav>
    </header>

    <main>
      <div v-if="currentPage === 'products'" class="section">
        <h2>Products</h2>
        <div class="loading" v-if="loading">Loading products...</div>
        <div v-else class="grid">
          <div v-for="product in products" :key="product.id" class="card">
            <h3>{{ product.name }}</h3>
            <p>${{ product.price }}</p>
            <p>Stock: {{ product.stock }}</p>
          </div>
        </div>
      </div>

      <div v-if="currentPage === 'orders'" class="section">
        <h2>Orders</h2>
        <div class="loading" v-if="loading">Loading orders...</div>
        <div v-else class="list">
          <div v-for="order in orders" :key="order.id" class="item">
            <p><strong>Order #{{ order.id }}</strong> - {{ order.status }}</p>
            <p>Total: ${{ order.total }}</p>
          </div>
        </div>
      </div>

      <div v-if="currentPage === 'users'" class="section">
        <h2>Users</h2>
        <div class="loading" v-if="loading">Loading users...</div>
        <div v-else class="list">
          <div v-for="user in users" :key="user.id" class="item">
            <p><strong>{{ user.name }}</strong></p>
            <p>{{ user.email }} ({{ user.role }})</p>
          </div>
        </div>
      </div>
    </main>

    <footer>
      <p>&copy; 2024 E-Commerce Platform. All rights reserved.</p>
    </footer>
  </div>
</template>

<script>
import axios from 'axios'

const API_URL = 'http://4.144.133.123:8080/api'

export default {
  name: 'App',
  data() {
    return {
      currentPage: 'products',
      products: [],
      orders: [],
      users: [],
      loading: false
    }
  },
  watch: {
    currentPage(newPage) {
      this.loadData(newPage)
    }
  },
  mounted() {
    console.log('App mounted, API_URL:', API_URL)
    this.loadData('products')
  },
  methods: {
    async loadData(page) {
      this.loading = true
      try {
        let url = `${API_URL}/${page}`
        console.log('Fetching:', url)
        const res = await axios.get(url)
        console.log('Response:', res.data)
        
        if (page === 'products') {
          this.products = res.data || []
        } else if (page === 'orders') {
          this.orders = res.data || []
        } else if (page === 'users') {
          this.users = res.data || []
        }
      } catch (error) {
        console.error(`Error loading ${page}:`, error.message)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body, html, #app {
  height: 100%;
}

#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #333;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

header {
  background: rgba(255, 255, 255, 0.95);
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

header h1 {
  margin-bottom: 15px;
  color: #667eea;
}

nav {
  display: flex;
  gap: 10px;
}

nav button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  background: #667eea;
  color: white;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s;
}

nav button:hover {
  background: #764ba2;
}

main {
  flex: 1;
  padding: 30px;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

.section {
  background: white;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.section h2 {
  margin-bottom: 20px;
  color: #667eea;
}

.loading {
  text-align: center;
  padding: 40px;
  color: #999;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
}

.card {
  background: #f9f9f9;
  padding: 15px;
  border-radius: 6px;
  border-left: 4px solid #667eea;
}

.card h3 {
  margin-bottom: 10px;
  color: #333;
}

.card p {
  margin: 5px 0;
  color: #666;
  font-size: 14px;
}

.list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.item {
  background: #f9f9f9;
  padding: 15px;
  border-radius: 6px;
  border-left: 4px solid #667eea;
}

.item p {
  margin: 5px 0;
  color: #666;
  font-size: 14px;
}

footer {
  background: rgba(255, 255, 255, 0.95);
  padding: 20px;
  text-align: center;
  border-top: 1px solid #eee;
  color: #666;
  font-size: 14px;
}
</style>

<template>
  <div id="app">
    <!-- Header -->
    <header class="header">
      <h1>🛍️ E-Commerce Platform</h1>
      <nav class="nav">
        <button
          v-for="page in pages"
          :key="page.key"
          :class="['nav-btn', { active: currentPage === page.key }]"
          @click="currentPage = page.key"
        >
          {{ page.label }}
        </button>
      </nav>
    </header>

    <!-- Main Content -->
    <main class="main">
      <div v-if="loading" class="loading">Loading...</div>

      <!-- Products Page -->
      <section v-if="currentPage === 'products'" class="section">
        <div class="card">
          <h2>Products</h2>
          <div class="grid">
            <div v-for="product in products" :key="product.id" class="item-card">
              <h3>{{ product.name }}</h3>
              <p class="price">${{ product.price }}</p>
              <p class="stock">Stock: {{ product.stock }}</p>
              <button class="btn btn-primary" @click="addToOrder(product)">Add to Order</button>
            </div>
          </div>
          <p v-if="products.length === 0 && !loading" class="empty">No products found</p>
        </div>
      </section>

      <!-- Orders Page -->
      <section v-if="currentPage === 'orders'" class="section">
        <div class="card">
          <h2>Orders</h2>
          <div class="form-section">
            <h3>Create New Order</h3>
            <form @submit.prevent="createOrder" class="form">
              <select v-model="newOrder.product_id" required>
                <option value="">Select Product</option>
                <option v-for="p in products" :key="p.id" :value="p.id">{{ p.name }} - ${{ p.price }}</option>
              </select>
              <input v-model.number="newOrder.quantity" type="number" min="1" placeholder="Quantity" required />
              <button type="submit" class="btn btn-primary">Create Order</button>
            </form>
          </div>
          <div class="table-container">
            <table v-if="orders.length > 0">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Product ID</th>
                  <th>Quantity</th>
                  <th>Total</th>
                  <th>Status</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="order in orders" :key="order.id">
                  <td>{{ order.id }}</td>
                  <td>{{ order.product_id }}</td>
                  <td>{{ order.quantity }}</td>
                  <td>${{ order.total?.toFixed(2) }}</td>
                  <td><span :class="['status', order.status]">{{ order.status }}</span></td>
                  <td>
                    <button v-if="order.status === 'pending'" class="btn btn-small btn-success" @click="payOrder(order)">Pay Now</button>
                  </td>
                </tr>
              </tbody>
            </table>
            <p v-else class="empty">No orders found</p>
          </div>
        </div>
      </section>

      <!-- Payments Page -->
      <section v-if="currentPage === 'payments'" class="section">
        <div class="card">
          <h2>Payments</h2>
          <div class="table-container">
            <table v-if="payments.length > 0">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Order ID</th>
                  <th>Amount</th>
                  <th>Method</th>
                  <th>Status</th>
                  <th>Transaction ID</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="payment in payments" :key="payment.id">
                  <td>{{ payment.id }}</td>
                  <td>{{ payment.order_id }}</td>
                  <td>${{ payment.amount?.toFixed(2) }}</td>
                  <td>{{ payment.method }}</td>
                  <td><span :class="['status', payment.status]">{{ payment.status }}</span></td>
                  <td>{{ payment.transaction_id || '-' }}</td>
                </tr>
              </tbody>
            </table>
            <p v-else class="empty">No payments found</p>
          </div>
        </div>
      </section>

      <!-- Users Page -->
      <section v-if="currentPage === 'users'" class="section">
        <div class="card">
          <h2>Users</h2>
          <div class="table-container">
            <table v-if="users.length > 0">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Name</th>
                  <th>Email</th>
                  <th>Role</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="user in users" :key="user.id">
                  <td>{{ user.id }}</td>
                  <td>{{ user.name }}</td>
                  <td>{{ user.email }}</td>
                  <td><span :class="['role', user.role]">{{ user.role }}</span></td>
                </tr>
              </tbody>
            </table>
            <p v-else class="empty">No users found</p>
          </div>
        </div>
      </section>
    </main>

    <!-- Notification -->
    <div v-if="notification" :class="['notification', notification.type]">
      {{ notification.message }}
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { initFaro } from './faro'

// Initialize Faro RUM
initFaro()

// Runtime config from window.__CONFIG__ (injected by entrypoint.sh) 
// Falls back to VITE env var, then to hardcoded default
const API_URL = (window.__CONFIG__ && window.__CONFIG__.API_URL) || import.meta.env.VITE_API_URL || 'http://4.144.133.123:8080/api'

export default {
  name: 'App',
  data() {
    return {
      currentPage: 'products',
      pages: [
        { key: 'products', label: 'Products' },
        { key: 'orders', label: 'Orders' },
        { key: 'payments', label: 'Payments' },
        { key: 'users', label: 'Users' }
      ],
      products: [],
      orders: [],
      payments: [],
      users: [],
      loading: false,
      notification: null,
      newOrder: {
        product_id: '',
        quantity: 1
      }
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
        if (page === 'products') {
          const res = await axios.get(`${API_URL}/products`)
          this.products = res.data || []
        } else if (page === 'orders') {
          if (this.products.length === 0) {
            const prodRes = await axios.get(`${API_URL}/products`)
            this.products = prodRes.data || []
          }
          const res = await axios.get(`${API_URL}/orders`)
          this.orders = res.data || []
        } else if (page === 'payments') {
          const res = await axios.get(`${API_URL}/payments`)
          this.payments = res.data || []
        } else if (page === 'users') {
          const res = await axios.get(`${API_URL}/users`)
          this.users = res.data || []
        }
      } catch (error) {
        console.error(`Error loading ${page}:`, error)
        this.showNotification(`Failed to load ${page}`, 'error')
      } finally {
        this.loading = false
      }
    },

    async createOrder() {
      try {
        const product = this.products.find(p => String(p.id) === String(this.newOrder.product_id))
        if (!product) {
          this.showNotification('Please select a product', 'error')
          return
        }

        const orderData = {
          user_id: 'user1',
          product_id: String(this.newOrder.product_id),
          quantity: this.newOrder.quantity,
          total: product.price * this.newOrder.quantity,
          status: 'pending'
        }

        await axios.post(`${API_URL}/orders`, orderData)
        this.showNotification('Order created successfully!', 'success')
        this.newOrder = { product_id: '', quantity: 1 }
        this.loadData('orders')
      } catch (error) {
        console.error('Error creating order:', error)
        this.showNotification('Failed to create order', 'error')
      }
    },

    async payOrder(order) {
      try {
        const paymentData = {
          order_id: parseInt(order.id),
          amount: order.total,
          method: 'credit_card'
        }

        await axios.post(`${API_URL}/payments`, paymentData)
        this.showNotification('Payment processed successfully!', 'success')
        this.loadData('orders')
      } catch (error) {
        console.error('Error processing payment:', error)
        this.showNotification('Failed to process payment', 'error')
      }
    },

    addToOrder(product) {
      this.newOrder.product_id = String(product.id)
      this.newOrder.quantity = 1
      this.currentPage = 'orders'
      this.showNotification(`Added ${product.name} to order form`, 'success')
    },

    showNotification(message, type) {
      this.notification = { message, type }
      setTimeout(() => {
        this.notification = null
      }, 3000)
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

#app {
  min-height: 100vh;
}

.header {
  background: rgba(255, 255, 255, 0.95);
  padding: 20px 40px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.header h1 {
  color: #667eea;
  margin-bottom: 15px;
  font-size: 2em;
}

.nav {
  display: flex;
  gap: 10px;
}

.nav-btn {
  padding: 10px 20px;
  border: none;
  border-radius: 25px;
  cursor: pointer;
  font-weight: 600;
  font-size: 14px;
  transition: all 0.3s;
  background: #e8eaf6;
  color: #333;
}

.nav-btn:hover {
  background: #667eea;
  color: white;
}

.nav-btn.active {
  background: #667eea;
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.main {
  padding: 40px;
  max-width: 1200px;
  margin: 0 auto;
}

.card {
  background: white;
  border-radius: 16px;
  padding: 30px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
}

.card h2 {
  color: #667eea;
  margin-bottom: 20px;
  font-size: 1.5em;
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.item-card {
  border: 2px solid #e8eaf6;
  border-left: 4px solid #667eea;
  border-radius: 12px;
  padding: 20px;
  transition: transform 0.2s, box-shadow 0.2s;
}

.item-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.item-card h3 {
  color: #333;
  margin-bottom: 8px;
}

.price {
  color: #667eea;
  font-weight: 600;
  font-size: 1.1em;
}

.stock {
  color: #666;
  font-size: 0.9em;
  margin-bottom: 12px;
}

.btn {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s;
}

.btn-primary {
  background: #667eea;
  color: white;
}

.btn-primary:hover {
  background: #5a6fd6;
}

.btn-success {
  background: #4caf50;
  color: white;
}

.btn-success:hover {
  background: #43a047;
}

.btn-small {
  padding: 6px 12px;
  font-size: 12px;
}

.form-section {
  background: #f8f9ff;
  padding: 20px;
  border-radius: 12px;
  margin-bottom: 20px;
}

.form-section h3 {
  margin-bottom: 12px;
  color: #333;
}

.form {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  align-items: center;
}

.form select,
.form input {
  padding: 10px 15px;
  border: 2px solid #e8eaf6;
  border-radius: 8px;
  font-size: 14px;
  outline: none;
  transition: border-color 0.3s;
}

.form select:focus,
.form input:focus {
  border-color: #667eea;
}

.table-container {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

th, td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #e8eaf6;
}

th {
  background: #f8f9ff;
  color: #333;
  font-weight: 600;
}

tr:hover {
  background: #f8f9ff;
}

.status {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  text-transform: capitalize;
}

.status.completed {
  background: #e8f5e9;
  color: #2e7d32;
}

.status.pending {
  background: #fff3e0;
  color: #ef6c00;
}

.role {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  text-transform: capitalize;
}

.role.admin {
  background: #fce4ec;
  color: #c62828;
}

.role.customer {
  background: #e3f2fd;
  color: #1565c0;
}

.loading {
  text-align: center;
  padding: 40px;
  color: white;
  font-size: 1.2em;
}

.empty {
  text-align: center;
  color: #999;
  padding: 20px;
}

.notification {
  position: fixed;
  bottom: 20px;
  right: 20px;
  padding: 15px 25px;
  border-radius: 10px;
  color: white;
  font-weight: 600;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  animation: slideIn 0.3s ease;
  z-index: 1000;
}

.notification.success {
  background: #4caf50;
}

.notification.error {
  background: #f44336;
}

@keyframes slideIn {
  from { transform: translateX(100%); opacity: 0; }
  to { transform: translateX(0); opacity: 1; }
}

@media (max-width: 768px) {
  .header { padding: 15px 20px; }
  .main { padding: 20px; }
  .nav { flex-wrap: wrap; }
  .form { flex-direction: column; }
}
</style>
<template>
  <div id="app">
    <!-- Login/Register Page -->
    <div v-if="!user" class="auth-page">
      <div class="auth-container">
        <div class="auth-left">
          <div class="auth-brand">
            <div class="brand-icon">
              <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 01-8 0"/></svg>
            </div>
            <h1>ShopHub</h1>
            <p>Discover amazing deals on premium products</p>
          </div>
          <div class="auth-features">
            <div class="feature-item">
              <span class="feature-icon">🚀</span>
              <span>Fast & Free Delivery</span>
            </div>
            <div class="feature-item">
              <span class="feature-icon">🔒</span>
              <span>Secure Payment</span>
            </div>
            <div class="feature-item">
              <span class="feature-icon">💯</span>
              <span>100% Authentic Products</span>
            </div>
          </div>
        </div>
        <div class="auth-right">
          <div class="auth-form">
            <h2>{{ authMode === 'login' ? 'Welcome Back!' : 'Create Account' }}</h2>
            <p class="auth-subtitle">{{ authMode === 'login' ? 'Sign in to continue shopping' : 'Join us and start exploring' }}</p>
            <form @submit.prevent="handleAuth">
              <div class="input-group" v-if="authMode === 'register'">
                <svg class="input-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
                <input v-model="authForm.name" type="text" placeholder="Full Name" required />
              </div>
              <div class="input-group">
                <svg class="input-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><polyline points="22,6 12,13 2,6"/></svg>
                <input v-model="authForm.email" type="email" placeholder="Email address" required />
              </div>
              <div class="input-group">
                <svg class="input-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0110 0v4"/></svg>
                <input v-model="authForm.password" type="password" placeholder="Password" required />
              </div>
              <button type="submit" class="btn-primary" :disabled="authLoading">
                <span v-if="authLoading" class="btn-spinner"></span>
                {{ authLoading ? '' : (authMode === 'login' ? 'Sign In' : 'Create Account') }}
              </button>
            </form>
            <p v-if="authError" class="auth-error">{{ authError }}</p>
            <div class="auth-divider"><span>or</span></div>
            <p class="auth-switch">
              {{ authMode === 'login' ? "Don't have an account?" : "Already have an account?" }}
              <a @click="authMode = authMode === 'login' ? 'register' : 'login'">
                {{ authMode === 'login' ? 'Sign Up' : 'Sign In' }}
              </a>
            </p>
            <p class="auth-hint">Demo: demo@shop.com / demo123</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Main App -->
    <div v-else class="shop">
      <!-- Top Bar -->
      <header class="topbar">
        <div class="topbar-inner">
          <div class="topbar-left" @click="currentPage = 'products'">
            <svg class="brand-logo" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 01-8 0"/></svg>
            <h1>ShopHub</h1>
          </div>
          <div class="topbar-center">
            <div class="search-box">
              <svg class="search-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>
              <input v-model="searchQuery" type="text" placeholder="Search in ShopHub..." class="search-input" />
            </div>
          </div>
          <div class="topbar-right">
            <button class="nav-btn" @click="currentPage = 'cart'" title="Cart">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/></svg>
              <span v-if="cart.length" class="badge">{{ cartItemCount }}</span>
            </button>
            <button class="nav-btn" @click="currentPage = 'orders'" title="Orders">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>
            </button>
            <button class="nav-btn error-trigger-btn" @click="triggerError" title="Trigger Test Error">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M10.29 3.86L1.82 18a2 2 0 001.71 3h16.94a2 2 0 001.71-3L13.71 3.86a2 2 0 00-3.42 0z"/><line x1="12" y1="9" x2="12" y2="13"/><line x1="12" y1="17" x2="12.01" y2="17"/></svg>
            </button>
            <div class="user-pill">
              <div class="user-avatar">{{ user.name.charAt(0).toUpperCase() }}</div>
              <span class="user-name">{{ user.name }}</span>
              <button @click="logout" class="logout-btn" title="Logout">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" y1="12" x2="9" y2="12"/></svg>
              </button>
            </div>
          </div>
        </div>
      </header>

      <!-- Products Page -->
      <main v-if="currentPage === 'products'" class="main-content">
        <!-- Hero Banner -->
        <div class="hero-banner">
          <div class="hero-content">
            <span class="hero-badge">🔥 Hot Deals</span>
            <h2>Up to 50% OFF</h2>
            <p>Premium electronics & lifestyle products</p>
          </div>
          <div class="hero-decorations">
            <div class="hero-circle c1"></div>
            <div class="hero-circle c2"></div>
          </div>
        </div>

        <!-- Categories -->
        <div class="categories-section">
          <div class="categories">
            <button :class="{ active: selectedCategory === 'All' }" @click="selectedCategory = 'All'">
              <span class="cat-icon">🏪</span><span class="cat-label">All</span>
            </button>
            <button v-for="cat in categories" :key="cat" :class="{ active: selectedCategory === cat }" @click="selectedCategory = cat">
              <span class="cat-icon">{{ getCategoryIcon(cat) }}</span><span class="cat-label">{{ cat }}</span>
            </button>
          </div>
        </div>

        <!-- Products -->
        <div class="section-header">
          <h3>{{ selectedCategory === 'All' ? '✨ All Products' : selectedCategory }}</h3>
          <span class="product-count">{{ filteredProducts.length }} items</span>
        </div>
        <div class="products-grid">
          <div v-for="product in filteredProducts" :key="product.id" class="product-card">
            <div class="product-image">
              <img :src="product.image" :alt="product.name" @error="handleImageError" loading="lazy" />
              <div class="product-badge" v-if="product.stock < 15">
                <span v-if="product.stock === 0">Sold Out</span>
                <span v-else>🔥 Hot</span>
              </div>
              <button class="quick-add" @click="addToCart(product)" :disabled="product.stock === 0" title="Add to cart">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/></svg>
              </button>
            </div>
            <div class="product-info">
              <span class="product-category">{{ product.category }}</span>
              <h3 class="product-name">{{ product.name }}</h3>
              <div class="product-rating">
                <span class="stars">★★★★★</span>
                <span class="rating-count">({{ Math.floor(Math.random() * 200 + 50) }})</span>
              </div>
              <div class="product-footer">
                <div class="price-block">
                  <span class="product-price">${{ product.price.toFixed(2) }}</span>
                  <span class="product-price-original" v-if="product.price > 500">${{ (product.price * 1.2).toFixed(2) }}</span>
                </div>
                <span class="product-sold">{{ product.stock }} left</span>
              </div>
            </div>
          </div>
        </div>
      </main>

      <!-- Cart Page -->
      <main v-if="currentPage === 'cart'" class="main-content">
        <div class="page-header">
          <button class="back-btn" @click="currentPage = 'products'">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="19" y1="12" x2="5" y2="12"/><polyline points="12 19 5 12 12 5"/></svg>
          </button>
          <h2>🛒 Shopping Cart <span class="count-label">({{ cartItemCount }} items)</span></h2>
        </div>
        <div v-if="cart.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="#ccc" stroke-width="1.5"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6"/></svg>
          </div>
          <h3>Your cart is empty</h3>
          <p>Looks like you haven't added anything yet</p>
          <button class="btn-primary" @click="currentPage = 'products'">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 002 2h14a2 2 0 002-2V6l-3-4z"/><line x1="3" y1="6" x2="21" y2="6"/><path d="M16 10a4 4 0 01-8 0"/></svg>
            Continue Shopping
          </button>
        </div>
        <div v-else class="cart-layout">
          <div class="cart-items">
            <div v-for="item in cart" :key="item.product.id" class="cart-item">
              <img :src="item.product.image" :alt="item.product.name" @error="handleImageError" />
              <div class="cart-item-info">
                <h4>{{ item.product.name }}</h4>
                <span class="cart-item-category">{{ item.product.category }}</span>
                <p class="cart-item-price">${{ item.product.price.toFixed(2) }}</p>
              </div>
              <div class="cart-item-qty">
                <button @click="updateQty(item, -1)">−</button>
                <span>{{ item.qty }}</span>
                <button @click="updateQty(item, 1)">+</button>
              </div>
              <span class="cart-item-total">${{ (item.product.price * item.qty).toFixed(2) }}</span>
              <button class="remove-btn" @click="removeFromCart(item)" title="Remove">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/></svg>
              </button>
            </div>
          </div>
          <div class="cart-summary">
            <h3>📋 Order Summary</h3>
            <div class="summary-row"><span>Subtotal ({{ cartItemCount }} items)</span><span>${{ cartTotal.toFixed(2) }}</span></div>
            <div class="summary-row"><span>Shipping</span><span class="free-shipping">FREE</span></div>
            <div class="summary-row discount"><span>🎁 Promo Discount</span><span>-$0.00</span></div>
            <div class="summary-row total"><span>Total Payment</span><span>${{ cartTotal.toFixed(2) }}</span></div>
            <div class="payment-method">
              <label>💳 Payment Method</label>
              <select v-model="paymentMethod">
                <option value="credit_card">💳 Credit Card</option>
                <option value="bank_transfer">🏦 Bank Transfer</option>
                <option value="e_wallet">📱 E-Wallet (GoPay/OVO)</option>
              </select>
            </div>
            <button class="btn-checkout" @click="checkout" :disabled="checkoutLoading">
              <span v-if="checkoutLoading" class="btn-spinner"></span>
              {{ checkoutLoading ? 'Processing...' : '🔒 Checkout Now' }}
            </button>
            <p class="secure-note">🛡️ Your payment is protected</p>
          </div>
        </div>
      </main>

      <!-- Orders Page -->
      <main v-if="currentPage === 'orders'" class="main-content">
        <div class="page-header">
          <button class="back-btn" @click="currentPage = 'products'">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="19" y1="12" x2="5" y2="12"/><polyline points="12 19 5 12 12 5"/></svg>
          </button>
          <h2>📦 My Orders</h2>
        </div>
        <div v-if="orders.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="#ccc" stroke-width="1.5"><path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
          </div>
          <h3>No orders yet</h3>
          <p>Start shopping to see your orders here</p>
          <button class="btn-primary" @click="currentPage = 'products'">Start Shopping</button>
        </div>
        <div v-else class="orders-list">
          <div v-for="order in orders" :key="order.id" class="order-card">
            <div class="order-header">
              <div class="order-id-row">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
                <span class="order-id">{{ order.id }}</span>
              </div>
              <span :class="['order-status', order.status]">
                <span class="status-dot"></span>
                {{ order.status }}
              </span>
            </div>
            <div class="order-items">
              <span v-for="item in order.items" :key="item.product_id" class="order-item-tag">
                {{ item.name }} × {{ item.quantity }}
              </span>
            </div>
            <div class="order-footer">
              <span class="order-total">${{ order.total.toFixed(2) }}</span>
              <span class="order-date">📅 {{ formatDate(order.created_at) }}</span>
            </div>
          </div>
        </div>
      </main>

      <!-- Notification -->
      <div v-if="notification" :class="['toast', notification.type]">
        <svg v-if="notification.type === 'success'" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 11-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>
        <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/></svg>
        <span>{{ notification.message }}</span>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import { initFaro, getFaro } from './faro'

initFaro()

const API_URL = (window.__CONFIG__ && window.__CONFIG__.API_URL) || import.meta.env.VITE_API_URL || 'http://4.144.133.123:8080/api'

export default {
  name: 'App',
  data() {
    return {
      currentPage: 'products',
      user: null,
      authMode: 'login',
      authForm: { name: '', email: '', password: '' },
      authLoading: false,
      authError: '',
      products: [],
      cart: [],
      orders: [],
      searchQuery: '',
      selectedCategory: 'All',
      paymentMethod: 'credit_card',
      checkoutLoading: false,
      notification: null,
    }
  },
  computed: {
    categories() {
      return [...new Set(this.products.map(p => p.category))]
    },
    filteredProducts() {
      return this.products.filter(p => {
        const matchCategory = this.selectedCategory === 'All' || p.category === this.selectedCategory
        const matchSearch = p.name.toLowerCase().includes(this.searchQuery.toLowerCase())
        return matchCategory && matchSearch
      })
    },
    cartItemCount() {
      return this.cart.reduce((sum, item) => sum + item.qty, 0)
    },
    cartTotal() {
      return this.cart.reduce((sum, item) => sum + item.product.price * item.qty, 0)
    },
  },
  mounted() {
    const saved = localStorage.getItem('shopUser')
    if (saved) {
      this.user = JSON.parse(saved)
      this.loadProducts()
      this.loadOrders()
    }
  },
  methods: {
    async handleAuth() {
      this.authLoading = true
      this.authError = ''
      try {
        const endpoint = this.authMode === 'login' ? '/auth/login' : '/auth/register'
        const res = await axios.post(`${API_URL}${endpoint}`, this.authForm)
        this.user = res.data.user
        localStorage.setItem('shopUser', JSON.stringify(this.user))
        this.loadProducts()
        this.loadOrders()
        getFaro()?.api.pushEvent('user_auth', { type: this.authMode, email: this.user.email })
      } catch (err) {
        this.authError = err.response?.data?.error || 'Authentication failed'
      } finally {
        this.authLoading = false
      }
    },
    logout() {
      this.user = null
      this.cart = []
      this.orders = []
      localStorage.removeItem('shopUser')
    },
    async loadProducts() {
      try {
        const res = await axios.get(`${API_URL}/products`)
        this.products = res.data || []
      } catch (err) {
        this.showNotification('Failed to load products', 'error')
      }
    },
    async loadOrders() {
      try {
        const res = await axios.get(`${API_URL}/orders?user_id=${this.user.id}`)
        this.orders = (res.data || []).sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
      } catch (err) {
        console.error('Failed to load orders', err)
      }
    },
    addToCart(product) {
      const existing = this.cart.find(i => i.product.id === product.id)
      if (existing) {
        existing.qty++
      } else {
        this.cart.push({ product, qty: 1 })
      }
      this.showNotification(`${product.name} added to cart`, 'success')
      getFaro()?.api.pushEvent('add_to_cart', { product_id: product.id, product_name: product.name })
    },
    updateQty(item, delta) {
      item.qty += delta
      if (item.qty <= 0) this.removeFromCart(item)
    },
    removeFromCart(item) {
      this.cart = this.cart.filter(i => i.product.id !== item.product.id)
    },
    async checkout() {
      this.checkoutLoading = true
      try {
        const orderData = {
          user_id: String(this.user.id),
          items: this.cart.map(i => ({
            product_id: i.product.id,
            name: i.product.name,
            price: i.product.price,
            quantity: i.qty,
          })),
        }
        const orderRes = await axios.post(`${API_URL}/orders`, orderData)
        const order = orderRes.data

        await axios.post(`${API_URL}/payments`, {
          order_id: order.id,
          amount: order.total,
          method: this.paymentMethod,
        })

        await axios.patch(`${API_URL}/orders/${order.id}`, { status: 'paid' })

        getFaro()?.api.pushEvent('checkout_completed', { order_id: order.id, total: String(order.total) })
        this.cart = []
        this.showNotification('Order placed successfully!', 'success')
        this.currentPage = 'orders'
        this.loadOrders()
      } catch (err) {
        this.showNotification('Checkout failed: ' + (err.response?.data?.error || err.message), 'error')
      } finally {
        this.checkoutLoading = false
      }
    },
    showNotification(message, type) {
      this.notification = { message, type }
      setTimeout(() => { this.notification = null }, 3000)
    },
    formatDate(d) {
      return new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })
    },
    handleImageError(e) {
      e.target.src = 'https://placehold.co/400x400/f8f9fa/adb5bd?text=Product'
    },
    getCategoryIcon(cat) {
      const icons = {
        Laptop: '💻', Smartphone: '📱', Audio: '🎧', Tablet: '📟',
        Accessories: '⌨️', Electronics: '📺', Fashion: '👟', Gaming: '🎮'
      }
      return icons[cat] || '📦'
    },
    triggerError() {
      try {
        const obj = null
        obj.nonExistentMethod()
      } catch (err) {
        getFaro()?.api.pushError(err)
        getFaro()?.api.pushLog(['[TEST ERROR] Manually triggered error for testing'], { level: 'error' })
        this.showNotification('Test error triggered! Check Grafana Faro dashboard.', 'error')
        throw err
      }
    },
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap');

* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif; background: #f0f2f5; color: #1a1a2e; }

/* Auth Page */
.auth-page { display: flex; justify-content: center; align-items: center; min-height: 100vh; background: linear-gradient(135deg, #0f0c29 0%, #302b63 50%, #24243e 100%); padding: 20px; }
.auth-container { display: flex; background: white; border-radius: 24px; overflow: hidden; width: 100%; max-width: 900px; box-shadow: 0 25px 80px rgba(0,0,0,0.4); min-height: 550px; }
.auth-left { flex: 1; background: linear-gradient(135deg, #ee5a24 0%, #f0932b 100%); padding: 60px 40px; display: flex; flex-direction: column; justify-content: center; color: white; }
.auth-brand { margin-bottom: 40px; }
.brand-icon { width: 60px; height: 60px; background: rgba(255,255,255,0.2); border-radius: 16px; display: flex; align-items: center; justify-content: center; margin-bottom: 20px; }
.auth-brand h1 { font-size: 2.2em; font-weight: 800; margin-bottom: 8px; }
.auth-brand p { font-size: 15px; opacity: 0.9; line-height: 1.5; }
.auth-features { display: flex; flex-direction: column; gap: 16px; }
.feature-item { display: flex; align-items: center; gap: 12px; font-size: 14px; opacity: 0.95; }
.feature-icon { font-size: 20px; }
.auth-right { flex: 1; padding: 50px 40px; display: flex; align-items: center; }
.auth-form { width: 100%; }
.auth-form h2 { font-size: 1.6em; font-weight: 700; margin-bottom: 6px; }
.auth-subtitle { color: #666; font-size: 14px; margin-bottom: 28px; }
.auth-form form { display: flex; flex-direction: column; gap: 14px; }
.input-group { position: relative; }
.input-icon { position: absolute; left: 14px; top: 50%; transform: translateY(-50%); color: #999; }
.input-group input { width: 100%; padding: 14px 14px 14px 44px; border: 2px solid #eee; border-radius: 12px; font-size: 14px; transition: all 0.2s; }
.input-group input:focus { outline: none; border-color: #ee5a24; box-shadow: 0 0 0 4px rgba(238,90,36,0.1); }
.btn-primary { padding: 14px; background: linear-gradient(135deg, #ee5a24, #f0932b); color: white; border: none; border-radius: 12px; font-size: 15px; font-weight: 700; cursor: pointer; transition: all 0.3s; display: flex; align-items: center; justify-content: center; gap: 8px; }
.btn-primary:hover { transform: translateY(-2px); box-shadow: 0 8px 25px rgba(238,90,36,0.3); }
.btn-primary:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }
.btn-spinner { width: 20px; height: 20px; border: 3px solid rgba(255,255,255,0.3); border-top-color: white; border-radius: 50%; animation: spin 0.8s linear infinite; }
@keyframes spin { to { transform: rotate(360deg); } }
.auth-error { color: #e74c3c; text-align: center; margin-top: 12px; font-size: 13px; background: #ffeaea; padding: 10px; border-radius: 8px; }
.auth-divider { display: flex; align-items: center; margin: 20px 0; color: #ccc; font-size: 13px; }
.auth-divider::before, .auth-divider::after { content: ''; flex: 1; height: 1px; background: #eee; }
.auth-divider span { padding: 0 12px; }
.auth-switch { text-align: center; font-size: 14px; color: #666; }
.auth-switch a { color: #ee5a24; font-weight: 600; cursor: pointer; }
.auth-switch a:hover { text-decoration: underline; }
.auth-hint { color: #bbb; text-align: center; margin-top: 16px; font-size: 12px; }

/* Topbar */
.topbar { background: white; box-shadow: 0 2px 20px rgba(0,0,0,0.06); position: sticky; top: 0; z-index: 100; border-bottom: 1px solid #f0f0f0; }
.topbar-inner { display: flex; align-items: center; padding: 14px 32px; max-width: 1400px; margin: 0 auto; }
.topbar-left { display: flex; align-items: center; gap: 10px; cursor: pointer; }
.brand-logo { color: #ee5a24; }
.topbar-left h1 { font-size: 1.4em; font-weight: 800; background: linear-gradient(135deg, #ee5a24, #f0932b); -webkit-background-clip: text; -webkit-text-fill-color: transparent; }
.topbar-center { flex: 1; max-width: 500px; margin: 0 32px; }
.search-box { position: relative; }
.search-icon { position: absolute; left: 16px; top: 50%; transform: translateY(-50%); color: #999; }
.search-input { width: 100%; padding: 12px 16px 12px 46px; border: 2px solid #f0f0f0; border-radius: 12px; font-size: 14px; background: #f8f9fa; transition: all 0.2s; }
.search-input:focus { outline: none; border-color: #ee5a24; background: white; box-shadow: 0 0 0 4px rgba(238,90,36,0.08); }
.topbar-right { display: flex; align-items: center; gap: 8px; }
.nav-btn { background: none; border: none; padding: 10px; border-radius: 12px; cursor: pointer; position: relative; color: #555; transition: all 0.2s; }
.nav-btn:hover { background: #fff3ed; color: #ee5a24; }
.error-trigger-btn { color: #e74c3c !important; }
.error-trigger-btn:hover { background: #fee !important; color: #c0392b !important; }
.badge { position: absolute; top: 2px; right: 2px; background: #ee5a24; color: white; border-radius: 50%; width: 18px; height: 18px; font-size: 10px; font-weight: 700; display: flex; align-items: center; justify-content: center; }
.user-pill { display: flex; align-items: center; gap: 8px; padding: 6px 12px; border-radius: 50px; background: #f8f9fa; margin-left: 8px; }
.user-avatar { width: 32px; height: 32px; border-radius: 50%; background: linear-gradient(135deg, #ee5a24, #f0932b); color: white; display: flex; align-items: center; justify-content: center; font-weight: 700; font-size: 13px; }
.user-name { font-size: 13px; font-weight: 600; color: #333; }
.logout-btn { background: none; border: none; padding: 6px; border-radius: 8px; cursor: pointer; color: #999; transition: all 0.2s; }
.logout-btn:hover { color: #e74c3c; background: #fee; }

/* Main */
.main-content { max-width: 1400px; margin: 0 auto; padding: 24px 32px; }

/* Hero Banner */
.hero-banner { background: linear-gradient(135deg, #ee5a24 0%, #f0932b 50%, #ffc048 100%); border-radius: 20px; padding: 40px 48px; margin-bottom: 32px; color: white; position: relative; overflow: hidden; }
.hero-content { position: relative; z-index: 2; }
.hero-badge { display: inline-block; background: rgba(255,255,255,0.2); backdrop-filter: blur(10px); padding: 6px 16px; border-radius: 20px; font-size: 13px; font-weight: 600; margin-bottom: 12px; }
.hero-banner h2 { font-size: 2.5em; font-weight: 800; margin-bottom: 8px; }
.hero-banner p { font-size: 16px; opacity: 0.9; }
.hero-decorations { position: absolute; top: 0; right: 0; bottom: 0; width: 50%; }
.hero-circle { position: absolute; border-radius: 50%; background: rgba(255,255,255,0.1); }
.hero-circle.c1 { width: 200px; height: 200px; top: -50px; right: -30px; }
.hero-circle.c2 { width: 150px; height: 150px; bottom: -40px; right: 100px; }

/* Categories */
.categories-section { margin-bottom: 28px; }
.categories { display: flex; gap: 10px; overflow-x: auto; padding-bottom: 8px; }
.categories::-webkit-scrollbar { height: 4px; }
.categories::-webkit-scrollbar-thumb { background: #ddd; border-radius: 4px; }
.categories button { display: flex; flex-direction: column; align-items: center; gap: 6px; padding: 14px 20px; border: 2px solid #eee; border-radius: 16px; background: white; cursor: pointer; font-size: 12px; font-weight: 600; transition: all 0.2s; white-space: nowrap; min-width: 80px; }
.categories button:hover { border-color: #ee5a24; background: #fff8f5; }
.categories button.active { background: linear-gradient(135deg, #ee5a24, #f0932b); color: white; border-color: #ee5a24; box-shadow: 0 4px 15px rgba(238,90,36,0.3); }
.cat-icon { font-size: 22px; }
.cat-label { font-size: 11px; }

/* Section Header */
.section-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.section-header h3 { font-size: 1.3em; font-weight: 700; }
.product-count { color: #999; font-size: 14px; }

/* Products Grid */
.products-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(240px, 1fr)); gap: 20px; }
.product-card { background: white; border-radius: 16px; overflow: hidden; box-shadow: 0 2px 12px rgba(0,0,0,0.04); transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1); border: 1px solid #f0f0f0; }
.product-card:hover { transform: translateY(-6px); box-shadow: 0 12px 40px rgba(0,0,0,0.1); border-color: #ee5a24; }
.product-image { height: 220px; background: #f8f9fa; display: flex; align-items: center; justify-content: center; overflow: hidden; position: relative; }
.product-image img { width: 100%; height: 100%; object-fit: cover; transition: transform 0.4s; }
.product-card:hover .product-image img { transform: scale(1.08); }
.product-badge { position: absolute; top: 12px; left: 12px; background: #e74c3c; color: white; padding: 4px 10px; border-radius: 6px; font-size: 11px; font-weight: 700; }
.quick-add { position: absolute; bottom: 12px; right: 12px; width: 40px; height: 40px; border-radius: 50%; background: white; border: none; cursor: pointer; display: flex; align-items: center; justify-content: center; box-shadow: 0 4px 15px rgba(0,0,0,0.15); opacity: 0; transform: translateY(10px); transition: all 0.3s; color: #ee5a24; }
.product-card:hover .quick-add { opacity: 1; transform: translateY(0); }
.quick-add:hover { background: #ee5a24; color: white; transform: scale(1.1); }
.quick-add:disabled { opacity: 0.4; cursor: not-allowed; }
.product-info { padding: 16px 18px 18px; }
.product-category { font-size: 11px; color: #ee5a24; font-weight: 700; text-transform: uppercase; letter-spacing: 0.5px; }
.product-name { font-size: 14px; margin: 6px 0 8px; color: #1a1a2e; line-height: 1.4; font-weight: 600; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }
.product-rating { display: flex; align-items: center; gap: 4px; margin-bottom: 10px; }
.stars { color: #f0932b; font-size: 12px; letter-spacing: -1px; }
.rating-count { color: #999; font-size: 11px; }
.product-footer { display: flex; justify-content: space-between; align-items: flex-end; }
.price-block { display: flex; flex-direction: column; }
.product-price { font-size: 20px; font-weight: 800; color: #ee5a24; }
.product-price-original { font-size: 12px; color: #999; text-decoration: line-through; }
.product-sold { font-size: 11px; color: #999; }

/* Page Header */
.page-header { display: flex; align-items: center; gap: 12px; margin-bottom: 24px; }
.page-header h2 { font-size: 1.5em; font-weight: 700; }
.count-label { font-size: 0.6em; color: #999; font-weight: 400; }
.back-btn { background: none; border: 2px solid #eee; padding: 8px; border-radius: 10px; cursor: pointer; color: #666; transition: all 0.2s; }
.back-btn:hover { border-color: #ee5a24; color: #ee5a24; }

/* Cart */
.cart-layout { display: grid; grid-template-columns: 1fr 380px; gap: 28px; }
.cart-items { display: flex; flex-direction: column; gap: 12px; }
.cart-item { display: flex; align-items: center; gap: 16px; background: white; padding: 18px; border-radius: 16px; border: 1px solid #f0f0f0; transition: all 0.2s; }
.cart-item:hover { border-color: #ee5a24; box-shadow: 0 4px 15px rgba(238,90,36,0.08); }
.cart-item img { width: 70px; height: 70px; object-fit: cover; border-radius: 12px; }
.cart-item-info { flex: 1; }
.cart-item-info h4 { font-size: 14px; font-weight: 600; margin-bottom: 2px; }
.cart-item-category { font-size: 11px; color: #999; }
.cart-item-price { color: #ee5a24; font-size: 14px; font-weight: 700; margin-top: 4px; }
.cart-item-qty { display: flex; align-items: center; gap: 0; border: 2px solid #eee; border-radius: 10px; overflow: hidden; }
.cart-item-qty button { width: 32px; height: 32px; border: none; background: #f8f9fa; cursor: pointer; font-size: 16px; font-weight: 600; transition: background 0.2s; }
.cart-item-qty button:hover { background: #fff3ed; color: #ee5a24; }
.cart-item-qty span { padding: 0 12px; font-weight: 600; font-size: 14px; }
.cart-item-total { font-weight: 700; min-width: 90px; text-align: right; font-size: 15px; }
.remove-btn { background: none; border: none; color: #ccc; cursor: pointer; padding: 8px; border-radius: 8px; transition: all 0.2s; }
.remove-btn:hover { color: #e74c3c; background: #fee; }
.cart-summary { background: white; padding: 28px; border-radius: 20px; height: fit-content; position: sticky; top: 90px; border: 1px solid #f0f0f0; }
.cart-summary h3 { margin-bottom: 20px; font-size: 1.1em; }
.summary-row { display: flex; justify-content: space-between; padding: 10px 0; font-size: 14px; color: #666; }
.summary-row.discount { color: #27ae60; }
.free-shipping { color: #27ae60; font-weight: 600; }
.summary-row.total { border-top: 2px solid #f0f0f0; margin-top: 12px; padding-top: 16px; font-size: 20px; font-weight: 800; color: #ee5a24; }
.payment-method { margin-top: 20px; }
.payment-method label { font-size: 13px; color: #666; display: block; margin-bottom: 8px; font-weight: 600; }
.payment-method select { width: 100%; padding: 12px 16px; border: 2px solid #eee; border-radius: 12px; font-size: 14px; appearance: none; background: white url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%23999' stroke-width='2'%3E%3Cpolyline points='6 9 12 15 18 9'/%3E%3C/svg%3E") no-repeat right 14px center; cursor: pointer; }
.payment-method select:focus { outline: none; border-color: #ee5a24; }
.btn-checkout { width: 100%; padding: 16px; background: linear-gradient(135deg, #ee5a24, #f0932b); color: white; border: none; border-radius: 14px; font-size: 16px; font-weight: 700; cursor: pointer; margin-top: 20px; transition: all 0.3s; display: flex; align-items: center; justify-content: center; gap: 8px; }
.btn-checkout:hover { transform: translateY(-2px); box-shadow: 0 8px 30px rgba(238,90,36,0.35); }
.btn-checkout:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }
.secure-note { text-align: center; font-size: 12px; color: #999; margin-top: 12px; }

/* Orders */
.orders-list { display: flex; flex-direction: column; gap: 14px; }
.order-card { background: white; padding: 22px 24px; border-radius: 16px; border: 1px solid #f0f0f0; transition: all 0.2s; }
.order-card:hover { border-color: #ee5a24; box-shadow: 0 4px 15px rgba(238,90,36,0.08); }
.order-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 14px; }
.order-id-row { display: flex; align-items: center; gap: 8px; color: #555; }
.order-id { font-weight: 700; font-size: 14px; color: #333; }
.order-status { padding: 5px 14px; border-radius: 20px; font-size: 12px; font-weight: 700; text-transform: capitalize; display: flex; align-items: center; gap: 6px; }
.status-dot { width: 6px; height: 6px; border-radius: 50%; }
.order-status.pending { background: #fff8e1; color: #f57c00; }
.order-status.pending .status-dot { background: #f57c00; }
.order-status.paid { background: #e8f5e9; color: #2e7d32; }
.order-status.paid .status-dot { background: #2e7d32; }
.order-status.completed { background: #e3f2fd; color: #1565c0; }
.order-status.completed .status-dot { background: #1565c0; }
.order-items { display: flex; flex-wrap: wrap; gap: 8px; margin-bottom: 14px; }
.order-item-tag { background: #f8f9fa; padding: 6px 14px; border-radius: 8px; font-size: 12px; font-weight: 500; border: 1px solid #eee; }
.order-footer { display: flex; justify-content: space-between; align-items: center; }
.order-total { font-weight: 800; font-size: 18px; color: #ee5a24; }
.order-date { color: #999; font-size: 13px; }

/* Empty State */
.empty-state { text-align: center; padding: 80px 20px; }
.empty-icon { margin-bottom: 20px; }
.empty-state h3 { font-size: 1.3em; margin-bottom: 8px; color: #333; }
.empty-state p { margin-bottom: 24px; font-size: 15px; color: #999; }

/* Toast */
.toast { position: fixed; bottom: 28px; right: 28px; padding: 16px 24px; border-radius: 14px; color: white; font-weight: 600; z-index: 999; animation: slideIn 0.4s cubic-bezier(0.4, 0, 0.2, 1); display: flex; align-items: center; gap: 10px; box-shadow: 0 10px 40px rgba(0,0,0,0.2); font-size: 14px; }
.toast.success { background: linear-gradient(135deg, #27ae60, #2ecc71); }
.toast.error { background: linear-gradient(135deg, #e74c3c, #c0392b); }
@keyframes slideIn { from { transform: translateX(100%) scale(0.8); opacity: 0; } to { transform: translateX(0) scale(1); opacity: 1; } }

/* Responsive */
@media (max-width: 1024px) {
  .cart-layout { grid-template-columns: 1fr; }
  .products-grid { grid-template-columns: repeat(3, 1fr); }
}
@media (max-width: 768px) {
  .topbar-inner { flex-wrap: wrap; gap: 10px; padding: 12px 16px; }
  .topbar-center { order: 3; max-width: 100%; margin: 0; flex-basis: 100%; }
  .products-grid { grid-template-columns: repeat(2, 1fr); gap: 12px; }
  .hero-banner { padding: 28px 24px; }
  .hero-banner h2 { font-size: 1.8em; }
  .auth-container { flex-direction: column; }
  .auth-left { padding: 30px; }
  .main-content { padding: 16px; }
  .user-name { display: none; }
}
@media (max-width: 480px) {
  .products-grid { grid-template-columns: repeat(2, 1fr); gap: 10px; }
  .product-image { height: 160px; }
  .product-info { padding: 12px; }
  .product-price { font-size: 16px; }
}
</style>

<template>
  <div class="home">
    <header class="hero">
      <h1>🛒 Ecommerce FP</h1>
      <p>Belanja produk terbaik dengan harga terjangkau</p>
    </header>

    <div class="container">
      <div class="categories">
        <button
          v-for="cat in categories"
          :key="cat.value"
          :class="['cat-btn', { active: selectedCategory === cat.value }]"
          @click="selectCategory(cat.value)"
        >{{ cat.label }}</button>
      </div>

      <div v-if="loading" class="loading">Memuat produk...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="product-grid">
        <div
          v-for="product in products"
          :key="product.id"
          class="product-card"
          @click="$router.push(`/products/${product.id}`)"
        >
          <div class="product-category">{{ product.category }}</div>
          <h3>{{ product.name }}</h3>
          <p class="product-desc">{{ product.description }}</p>
          <div class="product-footer">
            <span class="price">Rp {{ formatPrice(product.price) }}</span>
            <span :class="['stock', product.stock < 5 ? 'low' : '']">
              Stok: {{ product.stock }}
            </span>
          </div>
          <button class="btn-add" @click.stop="addToCart(product)">
            + Keranjang
          </button>
        </div>
      </div>
    </div>

    <button class="cart-fab" @click="$router.push('/cart')">
      🛒 <span class="badge">{{ cartStore.totalItems }}</span>
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useApi } from '@/composables/useApi'
import { useCartStore } from '@/stores/cart'
import { trackPageView } from '@/observability/faro'

const { getProducts, loading, error } = useApi()
const cartStore = useCartStore()

const products = ref([])
const selectedCategory = ref('')

const categories = [
  { label: 'Semua', value: '' },
  { label: 'Elektronik', value: 'electronics' },
  { label: 'Fashion', value: 'fashion' },
  { label: 'Buku', value: 'books' },
  { label: 'Makanan', value: 'food' },
  { label: 'Outdoor', value: 'outdoor' },
]

async function selectCategory(category) {
  selectedCategory.value = category
  try {
    products.value = await getProducts(category)
  } catch {}
}

function addToCart(product) {
  cartStore.addItem(product)
}

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID').format(price)
}

onMounted(() => {
  trackPageView('home')
  selectCategory('')
})
</script>

<style scoped>
.home { min-height: 100vh; }
.hero { text-align: center; padding: 40px 20px 20px; }
.hero h1 { font-size: 2rem; margin-bottom: 8px; }
.hero p { color: #6b7280; }
.container { max-width: 1200px; margin: 0 auto; padding: 0 20px 80px; }
.categories { display: flex; gap: 8px; flex-wrap: wrap; margin-bottom: 24px; }
.cat-btn { padding: 8px 16px; border: 2px solid #e0e0e0; border-radius: 20px; cursor: pointer; background: white; font-size: 14px; transition: all .2s; }
.cat-btn.active { background: #4f46e5; color: white; border-color: #4f46e5; }
.product-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 20px; }
.product-card { background: white; border: 1px solid #e5e7eb; border-radius: 12px; padding: 16px; cursor: pointer; transition: box-shadow .2s; }
.product-card:hover { box-shadow: 0 4px 20px rgba(0,0,0,.1); }
.product-category { font-size: 11px; color: #6b7280; text-transform: uppercase; margin-bottom: 6px; }
.product-desc { font-size: 13px; color: #6b7280; line-height: 1.4; margin: 8px 0; height: 40px; overflow: hidden; }
.product-footer { display: flex; justify-content: space-between; align-items: center; margin: 12px 0; }
.price { font-weight: 700; color: #4f46e5; }
.stock { font-size: 12px; color: #10b981; }
.stock.low { color: #ef4444; }
.btn-add { width: 100%; padding: 10px; background: #4f46e5; color: white; border: none; border-radius: 8px; cursor: pointer; font-size: 14px; }
.btn-add:hover { background: #4338ca; }
.cart-fab { position: fixed; bottom: 24px; right: 24px; background: #4f46e5; color: white; padding: 14px 20px; border-radius: 50px; cursor: pointer; font-size: 1rem; box-shadow: 0 4px 12px rgba(79,70,229,.4); border: none; }
.badge { background: #ef4444; border-radius: 50%; padding: 2px 7px; font-size: 12px; margin-left: 4px; }
.loading, .error { text-align: center; padding: 60px; color: #6b7280; }
.error { color: #ef4444; }
</style>

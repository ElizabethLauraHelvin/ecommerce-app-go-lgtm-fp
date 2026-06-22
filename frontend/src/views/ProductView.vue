<template>
  <div class="product-detail">
    <div class="container">
      <button class="btn-back" @click="$router.push('/')">← Kembali</button>

      <div v-if="loading" class="loading">Memuat produk...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else-if="product" class="detail-card">
        <div class="product-category">{{ product.category }}</div>
        <h2>{{ product.name }}</h2>
        <p class="desc">{{ product.description }}</p>
        <div class="price-row">
          <span class="price">Rp {{ formatPrice(product.price) }}</span>
          <span :class="['stock', product.stock < 5 ? 'low' : '']">
            Stok: {{ product.stock }}
          </span>
        </div>
        <div class="qty-row">
          <label>Jumlah:</label>
          <input v-model.number="quantity" type="number" min="1" :max="product.stock" />
        </div>
        <button
          class="btn-add"
          :disabled="product.stock === 0"
          @click="addToCart"
        >
          {{ product.stock === 0 ? 'Stok Habis' : '+ Tambah ke Keranjang' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useApi } from '@/composables/useApi'
import { useCartStore } from '@/stores/cart'
import { trackPageView } from '@/observability/faro'

const route = useRoute()
const router = useRouter()
const { getProduct, loading, error } = useApi()
const cartStore = useCartStore()

const product = ref(null)
const quantity = ref(1)

onMounted(async () => {
  try {
    product.value = await getProduct(route.params.id)
    trackPageView('product_detail', {
      product_id: route.params.id,
      product_name: product.value?.name,
    })
  } catch {}
})

function addToCart() {
  cartStore.addItem(product.value, quantity.value)
  router.push('/cart')
}

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID').format(price)
}
</script>

<style scoped>
.product-detail { min-height: 100vh; padding: 20px 0; }
.container { max-width: 600px; margin: 0 auto; padding: 0 20px; }
.btn-back { background: none; border: none; color: #4f46e5; cursor: pointer; font-size: 14px; margin-bottom: 20px; padding: 0; }
.detail-card { background: white; border: 1px solid #e5e7eb; border-radius: 16px; padding: 24px; }
.product-category { font-size: 12px; color: #6b7280; text-transform: uppercase; margin-bottom: 8px; }
h2 { font-size: 1.5rem; margin-bottom: 12px; }
.desc { color: #6b7280; line-height: 1.6; margin-bottom: 16px; }
.price-row { display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px; }
.price { font-size: 1.5rem; font-weight: 700; color: #4f46e5; }
.stock { font-size: 14px; color: #10b981; }
.stock.low { color: #ef4444; }
.qty-row { display: flex; align-items: center; gap: 12px; margin-bottom: 16px; }
.qty-row label { font-size: 14px; color: #6b7280; }
.qty-row input { width: 80px; padding: 8px; border: 1px solid #e5e7eb; border-radius: 8px; font-size: 16px; text-align: center; }
.btn-add { width: 100%; padding: 14px; background: #4f46e5; color: white; border: none; border-radius: 10px; font-size: 1rem; font-weight: 600; cursor: pointer; }
.btn-add:disabled { opacity: 0.5; cursor: not-allowed; }
.loading, .error { text-align: center; padding: 60px; color: #6b7280; }
</style>

<template>
  <div class="order-detail">
    <div class="container">
      <h2>✅ Order Berhasil!</h2>

      <div v-if="loading" class="loading">Memuat order...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else-if="order" class="order-card">
        <div class="order-header">
          <div>
            <div class="label">Order ID</div>
            <div class="value">#{{ order.id }}</div>
          </div>
          <span class="status-badge">{{ order.status }}</span>
        </div>

        <div class="order-items">
          <div v-for="item in order.items" :key="item.id" class="order-item">
            <span>Produk #{{ item.product_id }}</span>
            <span>× {{ item.quantity }}</span>
            <span>Rp {{ formatPrice(item.price * item.quantity) }}</span>
          </div>
        </div>

        <div class="order-total">
          <span>Total</span>
          <span class="total-price">Rp {{ formatPrice(order.total_price) }}</span>
        </div>

        <button class="btn-primary" @click="$router.push('/')">
          Lanjut Belanja
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useApi } from '@/composables/useApi'
import { trackPageView } from '@/observability/faro'

const route = useRoute()
const { getOrder, loading, error } = useApi()
const order = ref(null)

onMounted(async () => {
  try {
    order.value = await getOrder(route.params.id)
    trackPageView('order_success', {
      order_id: String(order.value?.id),
      total: String(order.value?.total_price),
    })
  } catch {}
})

function formatPrice(price) {
  return new Intl.NumberFormat('id-ID').format(price)
}
</script>

<style scoped>
.order-detail { min-height: 100vh; padding: 40px 0; }
.container { max-width: 600px; margin: 0 auto; padding: 0 20px; }
h2 { color: #10b981; margin-bottom: 24px; font-size: 1.8rem; }
.order-card { background: white; border: 1px solid #e5e7eb; border-radius: 16px; padding: 24px; }
.order-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px; }
.label { font-size: 12px; color: #6b7280; }
.value { font-size: 1.2rem; font-weight: 700; }
.status-badge { padding: 6px 14px; border-radius: 20px; font-size: 12px; font-weight: 600; background: #d1fae5; color: #065f46; text-transform: uppercase; }
.order-items { border-top: 1px solid #f3f4f6; border-bottom: 1px solid #f3f4f6; padding: 16px 0; margin-bottom: 16px; }
.order-item { display: flex; justify-content: space-between; padding: 6px 0; font-size: 14px; color: #374151; }
.order-total { display: flex; justify-content: space-between; font-size: 1.1rem; margin-bottom: 20px; font-weight: 600; }
.total-price { color: #4f46e5; }
.btn-primary { width: 100%; padding: 14px; background: #4f46e5; color: white; border: none; border-radius: 10px; font-size: 1rem; cursor: pointer; font-weight: 600; }
.loading, .error { text-align: center; padding: 60px; color: #6b7280; }
</style>

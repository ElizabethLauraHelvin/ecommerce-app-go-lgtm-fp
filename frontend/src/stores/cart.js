import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { trackAddToCart } from '@/observability/faro'

export const useCartStore = defineStore('cart', () => {
  const items = ref([])

  const totalItems = computed(() =>
    items.value.reduce((sum, i) => sum + i.quantity, 0)
  )

  const totalPrice = computed(() =>
    items.value.reduce((sum, i) => sum + i.product.price * i.quantity, 0)
  )

  function addItem(product, quantity = 1) {
    const existing = items.value.find((i) => i.product.id === product.id)
    if (existing) {
      existing.quantity += quantity
    } else {
      items.value.push({ product, quantity })
    }
    // Track ke Faro → Loki + terhubung ke trace
    trackAddToCart(product, quantity)
  }

  function removeItem(productId) {
    items.value = items.value.filter((i) => i.product.id !== productId)
  }

  function clearCart() {
    items.value = []
  }

  return { items, totalItems, totalPrice, addItem, removeItem, clearCart }
})

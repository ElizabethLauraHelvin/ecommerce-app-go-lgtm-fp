import { ref } from 'vue'
import { trackError } from '@/observability/faro'

// Microservices endpoints
const PRODUCT_SERVICE = window.PRODUCT_SERVICE_URL || 'http://localhost:8081'
const ORDER_SERVICE = window.ORDER_SERVICE_URL || 'http://localhost:8082'
const USER_SERVICE = window.USER_SERVICE_URL || 'http://localhost:8083'

export function useApi() {
  const loading = ref(false)
  const error = ref(null)

  async function request(path, options = {}, serviceURL = '') {
    loading.value = true
    error.value = null
    try {
      const url = serviceURL ? `${serviceURL}${path}` : path
      const response = await fetch(url, {
        headers: {
          'Content-Type': 'application/json',
          ...options.headers,
        },
        ...options,
      })
      if (!response.ok) {
        const body = await response.json().catch(() => ({}))
        throw new Error(body.error || `HTTP ${response.status}`)
      }
      return await response.json()
    } catch (err) {
      error.value = err.message
      // Track error ke Faro → Loki
      trackError(err.message, { path, method: options.method || 'GET' })
      throw err
    } finally {
      loading.value = false
    }
  }

  const getProducts = (category) =>
    request(`/api/products${category ? `?category=${category}` : ''}`, {}, PRODUCT_SERVICE)

  const getProduct = (id) =>
    request(`/api/products/${id}`, {}, PRODUCT_SERVICE)

  const createOrder = (payload) =>
    request('/api/orders', {
      method: 'POST',
      body: JSON.stringify(payload),
    }, ORDER_SERVICE)

  const getOrder = (id) =>
    request(`/api/orders/${id}`, {}, ORDER_SERVICE)

  const getUserOrders = (userId) =>
    request(`/api/users/${userId}/orders`, {}, USER_SERVICE)

  return {
    loading,
    error,
    getProducts,
    getProduct,
    createOrder,
    getOrder,
    getUserOrders,
  }
}

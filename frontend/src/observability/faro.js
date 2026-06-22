import {
  initializeFaro,
  getWebInstrumentations,
  LogLevel,
} from '@grafana/faro-web-sdk'
import { TracingInstrumentation } from '@grafana/faro-web-tracing'

let faroInstance = null

export function initFaro() {
  if (faroInstance || typeof window === 'undefined') return faroInstance

  // /faro/collect → nginx proxy → alloy.monitoring.svc.cluster.local:12347
  const collectorUrl =
    typeof window !== 'undefined' &&
    window.FARO_COLLECTOR_URL &&
    window.FARO_COLLECTOR_URL !== '%FARO_COLLECTOR_URL%'
      ? window.FARO_COLLECTOR_URL
      : '/faro/collect'

  faroInstance = initializeFaro({
    url: collectorUrl,
    app: {
      name: 'ecommerce-frontend',
      version: '1.0.0',
      environment: import.meta.env.MODE || 'production',
    },
    sessionTracking: {
      enabled: true,
      persistent: true,
    },
    instrumentations: [
      ...getWebInstrumentations({
        captureConsole: true,
        captureConsoleDisabledLevels: [],
        enablePerformanceInstrumentation: true,
      }),
      // Inject W3C traceparent header ke semua fetch ke BE
      // sehingga trace FE terhubung ke trace Go BE di Tempo
      new TracingInstrumentation({
        instrumentationOptions: {
          propagateTraceHeaderCorsUrls: [/.*/],
        },
      }),
    ],
  })

  return faroInstance
}

export function getFaro() {
  return faroInstance
}

// ── Custom event helpers ──────────────────────────────────────────────────

export function trackPageView(pageName, attributes = {}) {
  getFaro()?.api.pushEvent('page_view', { page: pageName, ...attributes })
}

export function trackAddToCart(product, quantity = 1) {
  getFaro()?.api.pushEvent('add_to_cart', {
    product_id: String(product.id),
    product_name: product.name,
    price: String(product.price),
    category: product.category,
    quantity: String(quantity),
  })
}

export function trackCheckout(cartItems, total) {
  getFaro()?.api.pushEvent('checkout_initiated', {
    item_count: String(cartItems.length),
    total_value: String(total),
    items: cartItems.map((i) => i.product.name).join(','),
  })
}

export function trackPurchase(orderId, total) {
  getFaro()?.api.pushEvent('purchase_completed', {
    order_id: String(orderId),
    total_value: String(total),
    timestamp: new Date().toISOString(),
  })
}

export function trackError(message, context = {}) {
  getFaro()?.api.pushError(new Error(message), { context })
}

export function setUser(userId, username) {
  getFaro()?.api.setUser({ id: String(userId), username, attributes: {} })
}

export function pushLog(message, level = LogLevel.INFO, context = {}) {
  getFaro()?.api.pushLog([message], { level, context })
}

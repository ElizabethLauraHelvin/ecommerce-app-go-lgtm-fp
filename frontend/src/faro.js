import {
  initializeFaro,
  getWebInstrumentations,
} from '@grafana/faro-web-sdk'
import { TracingInstrumentation } from '@grafana/faro-web-tracing'

let faroInstance = null

export function initFaro() {
  if (faroInstance) return faroInstance

  const config = window.__CONFIG__ || {}
  const collectorUrl = config.FARO_COLLECTOR_URL

  if (!collectorUrl) {
    console.warn('Faro: FARO_COLLECTOR_URL not set, skipping initialization')
    return null
  }

  faroInstance = initializeFaro({
    url: collectorUrl,
    app: {
      name: config.APP_NAME || 'ecommerce-frontend',
      version: config.APP_VERSION || '1.0.0',
      environment: config.APP_ENV || 'production',
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
      new TracingInstrumentation({
        instrumentationOptions: {
          propagateTraceHeaderCorsUrls: [/.*/],
        },
      }),
    ],
  })

  console.log('Faro initialized, collector:', collectorUrl)
  return faroInstance
}

export function getFaro() {
  return faroInstance
}

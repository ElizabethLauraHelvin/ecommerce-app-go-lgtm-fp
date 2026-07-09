#!/bin/sh

API_URL=${API_URL:-${VITE_API_URL:-http://localhost:8080/api}}

cat > /app/dist/config.js <<EOF
window.__CONFIG__ = {
  API_URL: "${API_URL}",
  FARO_COLLECTOR_URL: "${VITE_FARO_COLLECTOR_URL:-}",
  APP_NAME: "${VITE_APP_NAME:-ecommerce-frontend}",
  APP_VERSION: "${VITE_APP_VERSION:-1.0.0}",
  APP_ENV: "${VITE_APP_ENV:-production}"
};
EOF

echo "Runtime config generated with API_URL: ${API_URL}"

exec serve -s /app/dist -l 3000
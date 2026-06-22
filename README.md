# E-Commerce Microservices Application

Complete microservices e-commerce platform built with Go backends, Vue.js frontend, and Kubernetes deployment.

## 📂 Project Structure

```
ecommerce-app-go-lgtm-fp/
├── .github/
│   └── workflows/
│       └── build.yml                 # CI/CD Pipeline
├── backend/
│   ├── api-gateway/                  # API Gateway (routing)
│   │   ├── Dockerfile
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── go.sum
│   └── services/
│       ├── product/                  # Product Service
│       ├── order/                    # Order Service
│       ├── user/                     # User Service
│       └── payment/                  # Payment Service
├── frontend/                          # Vue.js Frontend
│   ├── Dockerfile
│   ├── package.json
│   ├── vite.config.js
│   ├── index.html
│   └── src/
│       ├── main.js
│       └── App.vue
├── k8s/                              # Kubernetes Manifests
│   ├── api-gateway/
│   ├── product-service/
│   ├── order-service/
│   ├── user-service/
│   ├── payment-service/
│   └── frontend/
└── README.md
```

## 🚀 Microservices

### Backend Services (Go)

#### 1. **API Gateway** (`backend/api-gateway`)
- Port: 8080
- Routes requests to microservices
- Reverse proxy implementation
- Health check endpoint: `/health`

#### 2. **Product Service** (`backend/services/product`)
- Port: 8080
- Endpoints:
  - `GET /products` - Get all products
  - `GET /products/:id` - Get product by ID
- Health check: `/health`

#### 3. **Order Service** (`backend/services/order`)
- Port: 8080
- Endpoints:
  - `GET /orders` - Get all orders
  - `GET /orders/:id` - Get order by ID
  - `POST /orders/create` - Create new order
- Health check: `/health`

#### 4. **User Service** (`backend/services/user`)
- Port: 8080
- Endpoints:
  - `GET /users` - Get all users
  - `GET /users/:id` - Get user by ID
  - `POST /users/create` - Create new user
- Health check: `/health`

#### 5. **Payment Service** (`backend/services/payment`)
- Port: 8080
- Endpoints:
  - `GET /payments` - Get all payments
  - `GET /payments/:id` - Get payment by ID
  - `POST /payments/process` - Process payment
- Health check: `/health`

### Frontend

#### Vue.js Frontend (`frontend`)
- Port: 3000
- Built with Vite
- Responsive UI with TailwindCSS
- Connects to API Gateway

## 🐳 Docker

### Build Images

```bash
# Build all services
docker build -t your-dockerhub-username/ecommerce-api-gateway ./backend/api-gateway
docker build -t your-dockerhub-username/ecommerce-product ./backend/services/product
docker build -t your-dockerhub-username/ecommerce-order ./backend/services/order
docker build -t your-dockerhub-username/ecommerce-user ./backend/services/user
docker build -t your-dockerhub-username/ecommerce-payment ./backend/services/payment
docker build -t your-dockerhub-username/ecommerce-frontend ./frontend
```

### Push to Docker Hub

```bash
docker push your-dockerhub-username/ecommerce-api-gateway:latest
docker push your-dockerhub-username/ecommerce-product:latest
docker push your-dockerhub-username/ecommerce-order:latest
docker push your-dockerhub-username/ecommerce-user:latest
docker push your-dockerhub-username/ecommerce-payment:latest
docker push your-dockerhub-username/ecommerce-frontend:latest
```

## ☸️ Kubernetes Deployment

### Prerequisites
- Kubernetes cluster (1.20+)
- kubectl configured
- Docker images pushed to registry

### Deploy

```bash
# Create namespace and deploy all services
kubectl apply -f k8s/api-gateway/api-gateway.yaml
kubectl apply -f k8s/product-service/product.yaml
kubectl apply -f k8s/order-service/order.yaml
kubectl apply -f k8s/user-service/user.yaml
kubectl apply -f k8s/payment-service/payment.yaml
kubectl apply -f k8s/frontend/frontend.yaml
```

### Verify Deployment

```bash
# Check deployments
kubectl get deployments -n ecommerce

# Check services
kubectl get svc -n ecommerce

# Check pods
kubectl get pods -n ecommerce

# Get API Gateway LoadBalancer IP
kubectl get svc api-gateway -n ecommerce
```

### Access Application

```bash
# Port forward to API Gateway
kubectl port-forward svc/api-gateway 8080:8080 -n ecommerce

# Port forward to Frontend
kubectl port-forward svc/frontend 3000:3000 -n ecommerce

# Visit frontend at http://localhost:3000
# API Gateway at http://localhost:8080
```

## 🔄 CI/CD Pipeline

GitHub Actions workflow automatically:
1. Builds Docker images for all services
2. Pushes to Docker Hub
3. Triggers on push to `main` branch
4. Only builds services with code changes

### Workflow File: `.github/workflows/build.yml`

**Triggered by:**
- Push to `main` branch with changes in `backend/**` or `frontend/**`
- Pull requests to `main` with same path changes

**Services Built:**
- API Gateway
- Product Service
- Order Service
- User Service
- Payment Service
- Frontend

## 🛠️ Development

### Local Setup

#### Backend Services
```bash
# Navigate to service directory
cd backend/api-gateway

# Install dependencies (if needed)
go mod download

# Run service
go run main.go
```

#### Frontend
```bash
cd frontend

# Install dependencies
npm install

# Development server
npm run dev

# Build for production
npm run build
```

## 📊 Architecture

```
┌─────────────────────────────────────────────────────┐
│                   Frontend (Vue.js)                 │
│                  http://localhost:3000              │
└──────────────────────┬──────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────┐
│              API Gateway (Go)                       │
│         Routes & Load Balancing                     │
│              http://localhost:8080                  │
└──────────────────────┬──────────────────────────────┘
        ┌───────────┬──────────┬───────────┬───────────┐
        │           │          │           │           │
        ▼           ▼          ▼           ▼           ▼
    ┌────────┐ ┌──────┐ ┌────────┐ ┌────────┐ ┌─────────┐
    │Product │ │Order │ │ User   │ │Payment │ │ Auth    │
    │Service │ │Service│ │Service │ │Service │ │Service  │
    └────────┘ └──────┘ └────────┘ └────────┘ └─────────┘
```

## 🔐 Environment Variables

### Frontend (.env)
```env
VITE_API_URL=http://api-gateway.ecommerce.svc.cluster.local:8080/api
VITE_FARO_COLLECTOR_URL=http://alloy.monitoring.svc.cluster.local:12347/collect
VITE_APP_NAME=ecommerce-frontend
VITE_APP_VERSION=1.0.0
VITE_APP_ENV=production
```

## 📝 Notes

- Each service has health check endpoints for Kubernetes probes
- Services use inter-pod communication via DNS
- LoadBalancer services for frontend and API gateway
- ClusterIP services for internal services
- Resource limits set for production workloads

## 🤝 Contributing

1. Create feature branch
2. Make changes
3. Push to GitHub
4. Create Pull Request
5. CI/CD pipeline runs automatically

## 📄 License

MIT License

## 👤 Author

Your Name


---

**Last Updated:** June 2024

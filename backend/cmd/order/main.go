package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/lgtm-fp/ecommerce-backend/internal/handlers"
	"github.com/lgtm-fp/ecommerce-backend/internal/middleware"
	"github.com/lgtm-fp/ecommerce-backend/internal/repository"
	"github.com/lgtm-fp/ecommerce-backend/observability"
)

func main() {
	// Set service name FIRST
	observability.SetServiceName("ecommerce-order-service")

	// 1. Init observability
	observability.Init()

	// 2. Koneksi DB
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		observability.GetEnv("DB_HOST", "postgres.ecommerce.svc.cluster.local"),
		observability.GetEnv("DB_PORT", "5432"),
		observability.GetEnv("DB_USER", "ecommerce"),
		observability.GetEnv("DB_PASSWORD", "secret"),
		observability.GetEnv("DB_NAME", "ecommerce"),
		observability.GetEnv("DB_SSLMODE", "disable"),
	)

	db, err := observability.OpenDBWithTracing(dsn)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak dapat dijangkau: %v", err)
	}
	log.Println("Database connected")

	// 3. Repositories
	orderRepo := repository.NewOrderRepository(db)

	// 4. Product Service URL (untuk inter-service call)
	productServiceURL := observability.GetEnv("PRODUCT_SERVICE_URL", "http://product-service:8081")

	// 5. Handlers
	orderHandler := handlers.NewOrderHandler(orderRepo, productServiceURL)

	// 5. Router
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"order"}`))
	})
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/api/orders", orderHandler.CreateOrder)
	mux.HandleFunc("/api/orders/", orderHandler.GetOrder)

	// 6. Middleware: tracing + metrics + CORS
	handler := middleware.Telemetry(middleware.CORS(mux))

	// 7. Log service URL
	log.Printf("Calling Product Service at: %s", productServiceURL)

	// 8. Start server
	port := observability.GetEnv("PORT", "8082")
	log.Printf("Order Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

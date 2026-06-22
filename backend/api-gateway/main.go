package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, `{"status":"healthy","service":"api-gateway"}`)
	})

	// Product Service
	productURL, _ := url.Parse("http://product-service.ecommerce.svc.cluster.local:8080")
	productProxy := httputil.NewSingleHostReverseProxy(productURL)
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.URL.Path = "/products"
		r.RequestURI = ""
		productProxy.ServeHTTP(w, r)
	})

	// Order Service
	orderURL, _ := url.Parse("http://order-service.ecommerce.svc.cluster.local:8080")
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)
	mux.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.URL.Path = "/orders"
		r.RequestURI = ""
		orderProxy.ServeHTTP(w, r)
	})

	// User Service
	userURL, _ := url.Parse("http://user-service.ecommerce.svc.cluster.local:8080")
	userProxy := httputil.NewSingleHostReverseProxy(userURL)
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.URL.Path = "/users"
		r.RequestURI = ""
		userProxy.ServeHTTP(w, r)
	})

	// Payment Service
	paymentURL, _ := url.Parse("http://payment-service.ecommerce.svc.cluster.local:8080")
	paymentProxy := httputil.NewSingleHostReverseProxy(paymentURL)
	mux.HandleFunc("/api/payments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.URL.Path = "/payments"
		r.RequestURI = ""
		paymentProxy.ServeHTTP(w, r)
	})

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
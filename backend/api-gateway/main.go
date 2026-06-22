package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, traceparent, tracestate")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func proxyHandler(targetURL *url.URL, pathPrefix string) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	return func(w http.ResponseWriter, r *http.Request) {
		// Strip /api prefix and forward the rest to the service
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
		r.URL.RawPath = strings.TrimPrefix(r.URL.RawPath, "/api")
		r.RequestURI = ""
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"api-gateway"}`)
	})

	// Product Service - handles /api/products and /api/products/*
	productURL, _ := url.Parse("http://product-service.ecommerce.svc.cluster.local:8080")
	mux.HandleFunc("/api/products/", proxyHandler(productURL, "/api"))
	mux.HandleFunc("/api/products", proxyHandler(productURL, "/api"))

	// Order Service - handles /api/orders and /api/orders/*
	orderURL, _ := url.Parse("http://order-service.ecommerce.svc.cluster.local:8080")
	mux.HandleFunc("/api/orders/", proxyHandler(orderURL, "/api"))
	mux.HandleFunc("/api/orders", proxyHandler(orderURL, "/api"))

	// User Service - handles /api/users and /api/users/*
	userURL, _ := url.Parse("http://user-service.ecommerce.svc.cluster.local:8080")
	mux.HandleFunc("/api/users/", proxyHandler(userURL, "/api"))
	mux.HandleFunc("/api/users", proxyHandler(userURL, "/api"))

	// Payment Service - handles /api/payments and /api/payments/*
	paymentURL, _ := url.Parse("http://payment-service.ecommerce.svc.cluster.local:8080")
	mux.HandleFunc("/api/payments/", proxyHandler(paymentURL, "/api"))
	mux.HandleFunc("/api/payments", proxyHandler(paymentURL, "/api"))

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
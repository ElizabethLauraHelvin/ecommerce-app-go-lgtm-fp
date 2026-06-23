package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ecommerce/observability"
)

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

var orders []Order
var orderCounter = 0

func getPaymentServiceURL() string {
	if url := os.Getenv("PAYMENT_SERVICE_URL"); url != "" {
		return url
	}
	return "http://payment-service.ecommerce.svc.cluster.local:8080"
}

func getUserServiceURL() string {
	if url := os.Getenv("USER_SERVICE_URL"); url != "" {
		return url
	}
	return "http://user-service.ecommerce.svc.cluster.local:8080"
}

func main() {
	mux := http.NewServeMux()
	httpClient := observability.TracedHTTPClient()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"order-service"}`)
	})

	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodPost {
			var order Order
			if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
				log.Printf("[ORDER] Create failed: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `{"error":"invalid request"}`)
				return
			}
			orderCounter++
			order.ID = fmt.Sprintf("ORD-%05d", orderCounter)
			order.CreatedAt = time.Now()
			order.Status = "pending"

			var total float64
			for _, item := range order.Items {
				total += item.Price * float64(item.Quantity)
			}
			order.Total = total

			// Validate user exists (call user-service)
			userReq, _ := http.NewRequestWithContext(r.Context(), "GET", getUserServiceURL()+"/users/"+order.UserID, nil)
			userResp, err := httpClient.Do(userReq)
			if err != nil {
				log.Printf("[ORDER] Failed to validate user=%s: %v", order.UserID, err)
			} else {
				userResp.Body.Close()
				log.Printf("[ORDER] Validated user=%s exists", order.UserID)
			}

			orders = append(orders, order)
			log.Printf("[ORDER] Created order=%s user=%s items=%d total=%.2f", order.ID, order.UserID, len(order.Items), order.Total)

			// Auto-process payment (call payment-service)
			paymentBody, _ := json.Marshal(map[string]interface{}{
				"order_id": order.ID,
				"amount":   order.Total,
				"method":   "credit_card",
			})
			payReq, _ := http.NewRequestWithContext(r.Context(), "POST", getPaymentServiceURL()+"/payments", bytes.NewBuffer(paymentBody))
			payReq.Header.Set("Content-Type", "application/json")
			payResp, err := httpClient.Do(payReq)
			if err != nil {
				log.Printf("[ORDER] Payment failed for order=%s: %v", order.ID, err)
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(order)
				return
			}
			payResp.Body.Close()

			if payResp.StatusCode == http.StatusCreated {
				order.Status = "paid"
				// Update in store
				for i, o := range orders {
					if o.ID == order.ID {
						orders[i].Status = "paid"
						break
					}
				}
				log.Printf("[ORDER] Payment successful, order=%s status=paid", order.ID)
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(order)
			return
		}

		// GET - return orders, optionally filter by user_id
		userID := r.URL.Query().Get("user_id")
		if userID != "" {
			var userOrders []Order
			for _, o := range orders {
				if o.UserID == userID {
					userOrders = append(userOrders, o)
				}
			}
			log.Printf("[ORDER] Listing orders for user=%s count=%d", userID, len(userOrders))
			json.NewEncoder(w).Encode(userOrders)
			return
		}

		log.Printf("[ORDER] Listing all orders count=%d", len(orders))
		json.NewEncoder(w).Encode(orders)
	})

	mux.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/orders/"):]

		// PATCH to update status
		if r.Method == http.MethodPatch {
			var update struct {
				Status string `json:"status"`
			}
			json.NewDecoder(r.Body).Decode(&update)
			for i, o := range orders {
				if o.ID == id {
					orders[i].Status = update.Status
					log.Printf("[ORDER] Updated order=%s status=%s", id, update.Status)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(orders[i])
					return
				}
			}
		}

		for _, o := range orders {
			if o.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(o)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error":"order not found"}`)
	})

	observability.Run("order-service", mux)
}

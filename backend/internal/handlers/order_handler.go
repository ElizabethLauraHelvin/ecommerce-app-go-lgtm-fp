package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/lgtm-fp/ecommerce-backend/internal/domain"
	"github.com/lgtm-fp/ecommerce-backend/internal/repository"
)

type OrderHandler struct {
	orderRepo          *repository.OrderRepository
	productServiceURL  string
	httpClient         *http.Client
}

func NewOrderHandler(orderRepo *repository.OrderRepository, productServiceURL string) *OrderHandler {
	return &OrderHandler{
		orderRepo:         orderRepo,
		productServiceURL: productServiceURL,
		httpClient:        &http.Client{},
	}
}

// getProductFromService calls Product Service to validate product
func (h *OrderHandler) getProductFromService(ctx context.Context, productID int64) (*domain.Product, error) {
	url := fmt.Sprintf("%s/api/products/%d", h.productServiceURL, productID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product not found")
	}
	
	var product domain.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req domain.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Request tidak valid")
		return
	}
	if len(req.Items) == 0 {
		writeError(w, http.StatusBadRequest, "Order harus memiliki minimal satu item")
		return
	}

	var totalPrice float64
	for _, item := range req.Items {
		// Call Product Service via HTTP untuk validasi
		product, err := h.getProductFromService(ctx, item.ProductID)
		if err != nil || product == nil {
			writeError(w, http.StatusBadRequest, "Produk tidak ditemukan")
			return
		}
		if product.Stock < item.Quantity {
			writeError(w, http.StatusBadRequest, "Stok tidak mencukupi")
			return
		}
		totalPrice += product.Price * float64(item.Quantity)
	}

	order, err := h.orderRepo.Create(ctx, req, totalPrice)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal membuat order")
		return
	}
	writeJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimPrefix(r.URL.Path, "/api/orders/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}
	order, err := h.orderRepo.GetByID(ctx, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil order")
		return
	}
	if order == nil {
		writeError(w, http.StatusNotFound, "Order tidak ditemukan")
		return
	}
	writeJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		writeError(w, http.StatusBadRequest, "User ID tidak valid")
		return
	}
	userID, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "User ID tidak valid")
		return
	}
	orders, err := h.orderRepo.GetByUserID(ctx, userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil order")
		return
	}
	writeJSON(w, http.StatusOK, orders)
}

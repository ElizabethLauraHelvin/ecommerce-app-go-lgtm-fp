package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/lgtm-fp/ecommerce-backend/internal/repository"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category := r.URL.Query().Get("category")
	var (
		result interface{}
		err    error
	)
	if category != "" {
		result, err = h.repo.GetByCategory(ctx, category)
	} else {
		result, err = h.repo.GetAll(ctx)
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil produk")
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}
	product, err := h.repo.GetByID(ctx, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Gagal mengambil produk")
		return
	}
	if product == nil {
		writeError(w, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}
	writeJSON(w, http.StatusOK, product)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

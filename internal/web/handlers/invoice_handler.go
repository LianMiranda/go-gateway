package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/domain"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/dto"
	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/service"
)

type invoiceHandler struct {
	service *service.InvoiceService
}

func NewInvoiceHandler(service *service.InvoiceService) *invoiceHandler {
	return &invoiceHandler{
		service: service,
	}
}

//Request autenticação via X-API-KEY
//EndPoint: /invoice
//Method: POST

func (h *invoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateInvoiceInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-KEY")

	output, err := h.service.Create(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

//EndPoint: /invoice/{id}
//Method: GET

func (h *invoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API KEY is required", http.StatusBadRequest)
		return
	}

	output, err := h.service.GetByID(id, apiKey)

	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

//EndPoint: /invoice
//Method: GET

func (h *invoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API KEY is required", http.StatusUnauthorized)
		return
	}

	output, err := h.service.ListByAccountAPIKey(apiKey)

	if err != nil {
			switch err {
			case domain.ErrAccountNotFound:
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

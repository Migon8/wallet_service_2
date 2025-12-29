package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"wallet_service_2/internal/models"
)

type WalleService interface {
	ProcessOperation(ctx context.Context, walletID uuid.UUID, opType string, amount float64) error
	GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error)
}

type WalletHandler struct {
	service WalletService
}

func NewWalletHadler(s WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}

func (hadler *WalletHandler) HandleOperation(w http.ResponseWriter, r *http.Request) {
	var req models.WalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("DEBUG: decode error: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := handler.service.ProcessOperation(r.Context(), req.ValletID, req.OperationType, req.Amount)

	if err != nil {
		log.Printf("ERROR IN PROCESS: %v", err)

		switch {
		case err == service.ErrValletNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case err == service.ErrLowBalance || err == service.ErrOperationType: // Ошибки при HTTP запросе
			http.Error(w, err.Error(), http.StatusBadRequest)

		default:
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (handler *WalletHandler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	walletID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid wallet uuid", http.StatusBadRequest)
		return
	}

	balance, err := handler.service.GetBalance(r.Context(), walletID)
	if err != nil {
		if err == service.ErrValletNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{ // САМ JSON Ответ
		"valletId": walletID,
		"balance":  balance,
	})
}

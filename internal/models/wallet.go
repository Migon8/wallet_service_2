package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	Deposit    = "DEPOSIT"
	Withdrawal = "WITHDRAW"
)

type Wallet struct {
	WalletID  uuid.UUID `json:"walletId"`
	Balance   float64   `json:"balance"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WalletRequest struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
}

package handler

import (
	"log/slog"
	"net/http"
	"time"
)

func InitRouter(handler *WalletHandler) http.Handler {

	mux := http.NewServeMux() // Может мне это не надо , ServeMux — стандартный роутер из net/http.

	// сопоставляет URL + метод → handler сам вызывает нужную функцию при запросе

	mux.HandleFunc("POST /api/v1/wallet", handler.HandleOperation) // Сами ручки , маршрут (2)
	mux.HandleFunc("GET /api/v1/wallets/{id}", handler.HandleGetBalance)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		mux.ServeHTTP(w, r)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)

		// Обёртка над mux , непонятно для чего пока что

	})

}

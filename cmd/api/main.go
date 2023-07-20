package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dyhalmeida/golang-order/internal/entity"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

func main() {
	addr := ":3333"
	err := http.ListenAndServe(addr, Router())
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", ShowOrder)
	fmt.Println("Server is running...")
	return r
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	order, err := entity.NewOrder(uuid.NewString(), 2000, 0.8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/go-event-driven/internal/application/dto"
	"github.com/AdiPP/go-event-driven/internal/application/usecase"
)

type OrderController struct {
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderController(createOrderUseCase *usecase.CreateOrderUseCase) *OrderController {
	return &OrderController{
		createOrderUseCase: createOrderUseCase,
	}
}

func (u *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var requestData dto.CreateOrderDTO
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = u.createOrderUseCase.Execute(r.Context(), requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("order created"))
}

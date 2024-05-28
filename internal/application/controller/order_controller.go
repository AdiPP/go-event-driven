package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/go-event-driven/internal/application/dto"
	"github.com/AdiPP/go-event-driven/internal/application/usecase"
	"github.com/AdiPP/go-event-driven/internal/domain/event"
)

type OrderController struct {
	createOrderUseCase *usecase.CreateOrderUseCase
	processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase
	stockMovementUseCase *usecase.StockMovementUseCase
	sendOrderEmailUseCase *usecase.SendOrderEmailUseCase
}

func NewOrderController(createOrderUseCase *usecase.CreateOrderUseCase, processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase, stockMovementUseCase *usecase.StockMovementUseCase, sendOrderEmailUseCase *usecase.SendOrderEmailUseCase) *OrderController {
	return &OrderController{
		createOrderUseCase: createOrderUseCase,
		processOrderPaymentUseCase: processOrderPaymentUseCase, 
		stockMovementUseCase: stockMovementUseCase,
		sendOrderEmailUseCase: sendOrderEmailUseCase,
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

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("order created"))
}

func(u *OrderController) ProcessOrderPayment(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent

	json.NewDecoder(r.Body).Decode(&body)

	err := u.processOrderPaymentUseCase.Execute(r.Context(), &body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("payment created"))
}

func(u *OrderController) StockMovement(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent

	json.NewDecoder(r.Body).Decode(&body)

	err := u.stockMovementUseCase.Execute(r.Context(), &body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("stock moved"))
}

func (u *OrderController) SendOrderEmail(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent

	json.NewDecoder(r.Body).Decode(&body)

	err := u.sendOrderEmailUseCase.Execute(r.Context(), &body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("stock moved"))
}
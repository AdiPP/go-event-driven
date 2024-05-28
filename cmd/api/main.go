package main

import (
	"fmt"
	"net/http"

	"github.com/AdiPP/go-event-driven/internal/application/controller"
	"github.com/AdiPP/go-event-driven/internal/application/usecase"
	"github.com/AdiPP/go-event-driven/internal/infra/queue"
)

func main() {
	queueAdapter := queue.NewMemoryQueueAdapter()
	
	createOrderUseCase := usecase.NewCreateOrderUseCase(queueAdapter)
	processOrderPaymentUseCase := usecase.NewProcessOrderPaymentUseCase(queueAdapter)
	stockMovementUseCase := usecase.NewStockMovementUseCase()
	sendOrderEmailUseCase := usecase.NewSendOrderEmailUseCase()
	
	orderController := controller.NewOrderController(
		createOrderUseCase,
		processOrderPaymentUseCase,
		stockMovementUseCase,
		sendOrderEmailUseCase,
	)

	http.HandleFunc("POST /create-order", orderController.CreateOrder)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"github.com/AdiPP/go-event-driven/internal/application/controller"
	"github.com/AdiPP/go-event-driven/internal/infra/queue"
	"github.com/AdiPP/go-event-driven/internal/usecase"
	"net/http"
)

func main() {
	queueAdapter := queue.NewMemoryQueueAdapter()
	createOrderUseCase := usecase.NewCreateOrderUseCase(queueAdapter)
	orderController := controller.NewOrderController(createOrderUseCase)

	http.HandleFunc("POST /create-order", orderController.CreateOrder)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

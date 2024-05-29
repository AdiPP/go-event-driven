package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/AdiPP/go-event-driven/internal/application/controller"
	"github.com/AdiPP/go-event-driven/internal/application/usecase"
	"github.com/AdiPP/go-event-driven/internal/domain/event"
	"github.com/AdiPP/go-event-driven/internal/infra/queue"
)

func main() {
	ctx := context.Background()

	// initialize queue
	queueAdapter := queue.NewRabbitMQAdapter("amqp://guest:guest@localhost:5672/")
	
	// use cases
	createOrderUseCase := usecase.NewCreateOrderUseCase(queueAdapter)
	processOrderPaymentUseCase := usecase.NewProcessOrderPaymentUseCase(queueAdapter)
	stockMovementUseCase := usecase.NewStockMovementUseCase()
	sendOrderEmailUseCase := usecase.NewSendOrderEmailUseCase()
	
	// controllers
	orderController := controller.NewOrderController(
		createOrderUseCase,
		processOrderPaymentUseCase,
		stockMovementUseCase,
		sendOrderEmailUseCase,
	)

	// register routes
	http.HandleFunc("POST /create-order", orderController.CreateOrder)

	// mapping listeners
	var list map[reflect.Type][]func (w http.ResponseWriter, r *http.Request) = map[reflect.Type][] func (w http.ResponseWriter, r *http.Request)  {
		reflect.TypeOf(event.OrderCreatedEvent{}): {
			orderController.ProcessOrderPayment,
			orderController.StockMovement,
			orderController.SendOrderEmail,
		},
	}

	// register listeners
	for eventType, handlers := range list {
		for _, handler := range handlers {
			queueAdapter.ListenerRegister(eventType, handler)
		}
	}

	// connect queue
	err := queueAdapter.Connect(ctx)

	if err != nil {
		log.Fatalf("Error connect queue %s", err)
	}

	defer queueAdapter.Dissconect(ctx)

	// start consuming queues
	orderCreatedEvent := reflect.TypeOf(event.OrderCreatedEvent{}).Name()

	go func (ctx context.Context, queueName string)  {
		queueAdapter.StartConsuming(ctx, queueName)

		if err != nil {
			log.Fatalf("Error running consumer %s: %s", queueName, err)
		}
	}(ctx, orderCreatedEvent)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

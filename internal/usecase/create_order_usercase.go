package usecase

import (
	"context"
	"fmt"

	"github.com/AdiPP/go-event-driven/internal/application/dto"
	"github.com/AdiPP/go-event-driven/internal/domain/entity"
	"github.com/AdiPP/go-event-driven/internal/domain/event"
	"github.com/AdiPP/go-event-driven/internal/domain/queue"
)

type CreateOrderUseCase struct {
	publisher queue.Publisher
}

func NewCreateOrderUseCase(publisher queue.Publisher) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		publisher: publisher,
	}
}

func (u *CreateOrderUseCase) Execute(ctx context.Context, input dto.CreateOrderDTO) error {
	fmt.Println("--- CreateOrderUseCase ---")

	order, err := entity.NewOrderEntity()
	if err != nil {
		return err
	}

	for _, item := range input.Items {
		fakeProductName := "Product" + item.ProductId
		fakeProductPrice := 10.50

		i := entity.NewOrderItemEntity(fakeProductName, fakeProductPrice, item.Qtd)

		order.AddItem(i)
	}

	var eventItems []event.OrderItem

	for _, item := range order.GetItems() {
		eventItems = append(eventItems, event.OrderItem{
			ProductName: item.GetProductName(),
			TotalPrice:  item.GetProductPrice(),
			Quantity:    item.GetQuantity(),
		})
	}

	err = u.publisher.Publish(ctx, event.OrderCreatedEvent{
		Id:         order.GetID(),
		TotalPrice: order.GetTotalPrice(),
		Status:     order.GetStatus(),
		Items:      eventItems,
	})
	if err != nil {
		return err
	}

	return nil
}

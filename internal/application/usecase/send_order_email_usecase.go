package usecase

import (
	"context"
	"fmt"

	"github.com/AdiPP/go-event-driven/internal/domain/event"
)

type SendOrderEmailUseCase struct {
}

func NewSendOrderEmailUseCase() *SendOrderEmailUseCase {
	return &SendOrderEmailUseCase{}
}

func (hh *SendOrderEmailUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	fmt.Println("--- SendOrderEmailUseCase ---")
	fmt.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
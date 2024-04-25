package queue

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
)

type MemoryQueueAdapter struct {
}

func NewMemoryQueueAdapter() *MemoryQueueAdapter {
	return &MemoryQueueAdapter{}
}

func (m MemoryQueueAdapter) Publish(ctx context.Context, eventPayload interface{}) error {
	eventType := reflect.TypeOf(eventPayload)

	payloadJson, err := json.Marshal(eventPayload)
	if err != nil {
		return err
	}

	log.Printf("** [Publish] %s: %v ---", eventType, string(payloadJson))
	return nil
}

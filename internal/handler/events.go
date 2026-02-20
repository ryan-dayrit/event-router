package handler

import (
	"sync"

	"event-router/internal/sender"
)

type eventsHandler struct{}

func NewEventsHandler() *eventsHandler {
	return &eventsHandler{}
}

func (h *eventsHandler) ProcessEvents(events []map[string]interface{}) error {
	customerEventsCh := make(chan map[string]interface{})
	productEventsCh := make(chan map[string]interface{})

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go sender.SendCustomerEvents(customerEventsCh, &waitGroup)
	go sender.SendProductEvents(productEventsCh, &waitGroup)

	for _, event := range events {
		switch event["event_type"].(string) {
		case "product":
			productEventsCh <- event
		case "customer":
			customerEventsCh <- event
		}
	}

	close(customerEventsCh)
	close(productEventsCh)
	waitGroup.Wait()

	return nil
}

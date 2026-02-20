package handler

import (
	"event-router/internal/sender"
)

type eventsHandler struct{}

func NewEventsHandler() *eventsHandler {
	return &eventsHandler{}
}

func (h *eventsHandler) ProcessEvents(events []map[string]interface{}) error {
	customerEventsCh := make(chan map[string]interface{})
	productEventsCh := make(chan map[string]interface{})

	go sender.SendCustomerEvents(customerEventsCh)
	go sender.SendProductEvents(productEventsCh)

	for _, event := range events {
		switch event["event_type"].(string) {
		case "customer":
			customerEventsCh <- event
		case "product":
			productEventsCh <- event
		}
	}
	close(customerEventsCh)
	close(productEventsCh)

	return nil
}

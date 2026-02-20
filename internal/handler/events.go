package handler

import (
	"fmt"
)

type eventsHandler struct{}

func NewEventsHandler() *eventsHandler {
	return &eventsHandler{}
}

func processCustomerEvents(customerEventsCh <-chan map[string]interface{}) {
	for event := range customerEventsCh {
		fmt.Printf("[customer] id: %s \n", event["id"].(string))

		metaData := event["meta_data"].(map[string]interface{})
		fmt.Printf("[customer] name: %s \n", metaData["name"].(string))
		fmt.Printf("[customer] phone: %s \n", metaData["phone"].(string))
		fmt.Printf("[customer] email: %s \n", metaData["email"].(string))

		address := metaData["address"].(map[string]interface{})
		fmt.Printf("[customer] state: %s \n", address["state"].(string))
		fmt.Printf("[customer] country: %s \n", address["country"].(string))
	}
}

func processProductEvents(productEventsCh <-chan map[string]interface{}) {
	for event := range productEventsCh {
		fmt.Printf("[product] id: %s \n", event["id"].(string))
		metaData := event["meta_data"].(map[string]interface{})
		fmt.Printf("[product] brand: %s \n", metaData["brand"].(string))
		fmt.Printf("[product] name: %s \n", metaData["name"].(string))
		fmt.Printf("[product] price: %s \n", metaData["price"].(string))
	}
}

func (h *eventsHandler) ProcessEvents(events []map[string]interface{}) error {
	customerEventsCh := make(chan map[string]interface{})
	productEventsCh := make(chan map[string]interface{})

	go processCustomerEvents(customerEventsCh)
	go processProductEvents(productEventsCh)

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

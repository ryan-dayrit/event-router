package main

import (
	"fmt"
	"log"

	"event-router/internal/handler"
	json_repository "event-router/internal/repository/json"
)

func main() {
	repostory := json_repository.NewRepository()
	events, err := repostory.GetEvents("testdata/events.json")
	if err != nil {
		log.Panic(fmt.Errorf("failed to fetch events from json file: %s", err))
	}

	handler := handler.NewEventsHandler()
	eventsMap := handler.ProcessEvents(events)

	fmt.Printf("# customer events processed: %d\n", len(eventsMap["customer"]))
	fmt.Printf("# product events processed: %d\n", len(eventsMap["product"]))
}

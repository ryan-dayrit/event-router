package main

import (
	_ "embed"
	"encoding/json"
	"log"

	"event-router/internal/handler"
)

//go:embed testdata/events.json
var eventsJson []byte

func main() {
	var events []map[string]interface{}
	if err := json.Unmarshal(eventsJson, &events); err != nil {
		log.Fatal(err)
	}

	handler := handler.NewEventsHandler()
	handler.ProcessEvents(events)
}

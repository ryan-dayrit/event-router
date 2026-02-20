package sender

import (
	"fmt"
)

func SendProductEvents(productEventsCh <-chan map[string]interface{}) {
	for event := range productEventsCh {
		fmt.Printf("[product] id: %s \n", event["id"].(string))
		metaData := event["meta_data"].(map[string]interface{})
		fmt.Printf("[product] brand: %s \n", metaData["brand"].(string))
		fmt.Printf("[product] name: %s \n", metaData["name"].(string))
		fmt.Printf("[product] price: %s \n", metaData["price"].(string))
	}
}

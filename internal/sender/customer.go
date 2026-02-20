package sender

import (
	"fmt"
)

func SendCustomerEvents(customerEventsCh <-chan map[string]interface{}) {
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

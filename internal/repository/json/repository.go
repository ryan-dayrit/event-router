package json

import (
	"encoding/json"
	"os"
)

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetEvents(path string) ([]map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var events []map[string]interface{}
	if err := json.Unmarshal(content, &events); err != nil {
		return nil, err
	}
	return events, nil
}

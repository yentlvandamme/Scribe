package parse

import (
	"encoding/json"
	"io"
	"time"
)

type Snippet struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"ModifiedOn"`
}

type Snippets struct {
	Version  string             `json:"version"`
	Snippets map[string]Snippet `json:"Snippets"`
}

func ParseJson(reader io.Reader) (Snippets, error) {
	var snippets Snippets

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&snippets)
	if err != nil {
		return snippets, err
	}

	return snippets, nil
}

func ParseToBytes(data interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	return jsonData, err
}
